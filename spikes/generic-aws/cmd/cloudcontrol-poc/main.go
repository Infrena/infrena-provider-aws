// Command cloudcontrol-poc is THROWAWAY SPIKE CODE. It manages an AWS::EC2::VPC end to end through the
// AWS Cloud Control API with no VPC-specific code: the resource type is a string, the desired state is a
// JSON document, and the attribute flags come from the resource's published schema.
//
//	go run ./cmd/cloudcontrol-poc -profile infrata -region us-east-1          # schema only, read-only
//	go run ./cmd/cloudcontrol-poc -profile infrata -region us-east-1 -live    # creates, updates and deletes ONE VPC
package main

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudcontrol"
	cctypes "github.com/aws/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/infrata/infrata-provider-aws/spikes/genericaws/cfnschema"
)

func main() {
	profile := flag.String("profile", "infrata", "AWS profile")
	region := flag.String("region", "us-east-1", "AWS region")
	typeName := flag.String("type", "AWS::EC2::VPC", "CloudFormation resource type")
	live := flag.Bool("live", false, "create, update and delete one resource")
	flag.Parse()
	if err := run(*profile, *region, *typeName, *live); err != nil {
		fmt.Fprintln(os.Stderr, "FAILED:", err)
		os.Exit(1)
	}
}

func run(profile, region, typeName string, live bool) error {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile), config.WithRegion(region))
	if err != nil {
		return err
	}

	// 1. The schema, from the registry. Nothing about VPCs is compiled in.
	dt, err := cloudformation.NewFromConfig(cfg).DescribeType(ctx, &cloudformation.DescribeTypeInput{
		Type: cftypes.RegistryTypeResource, TypeName: aws.String(typeName),
	})
	if err != nil {
		return fmt.Errorf("DescribeType: %w", err)
	}
	doc, err := cfnschema.Parse([]byte(aws.ToString(dt.Schema)))
	if err != nil {
		return err
	}
	def, rep := cfnschema.Definition(doc)
	if err := def.Validate(); err != nil {
		return fmt.Errorf("generated definition invalid: %w", err)
	}
	fmt.Printf("%s -> %s (provisioning %s), %d attributes, identifier %v, update handler %v\n",
		typeName, def.Type, dt.ProvisioningType, len(def.Attributes), rep.Identifier, rep.HasUpdate)
	for _, name := range []string{"CidrBlock", "VpcId", "EnableDnsHostnames", "Tags"} {
		if a, ok := def.Attributes[name]; ok {
			fmt.Printf("  %-20s kind=%-7s required=%-5v computed=%-5v forcenew=%v\n", name, a.Kind, a.Required, a.Computed, a.ForceNew)
		}
	}
	fmt.Printf("  write-only: %v\n  conditional create-only: %v\n  nested flags not expressible: %d\n",
		rep.WriteOnly, rep.Conditional, len(rep.NestedFlags))
	if !live {
		return nil
	}

	// 2. Create / read / update / delete, generically.
	cc := cloudcontrol.NewFromConfig(cfg)
	run := fmt.Sprintf("%d", time.Now().Unix())
	desired := map[string]any{
		"CidrBlock": "10.99.0.0/16",
		"Tags":      []map[string]string{{"Key": "infrata-spike", "Value": run}},
	}
	body, _ := json.Marshal(desired)

	start := time.Now()
	created, err := cc.CreateResource(ctx, &cloudcontrol.CreateResourceInput{
		TypeName: aws.String(typeName), DesiredState: aws.String(string(body)), ClientToken: aws.String(token()),
	})
	if err != nil {
		return fmt.Errorf("CreateResource: %w", err)
	}
	ev, err := await(ctx, cc, created.ProgressEvent)
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	id := aws.ToString(ev.Identifier)
	fmt.Printf("created %s in %s (request %s)\n", id, time.Since(start).Round(time.Millisecond), aws.ToString(created.ProgressEvent.RequestToken))
	defer func() {
		// Best-effort cleanup if anything below fails before the explicit delete.
		if id != "" {
			_, _ = cc.DeleteResource(context.Background(), &cloudcontrol.DeleteResourceInput{TypeName: aws.String(typeName), Identifier: aws.String(id), ClientToken: aws.String(token())})
		}
	}()

	if err := show(ctx, cc, typeName, id, "after create"); err != nil {
		return err
	}

	start = time.Now()
	patch := `[{"op":"add","path":"/EnableDnsHostnames","value":true},` +
		`{"op":"replace","path":"/Tags","value":[{"Key":"infrata-spike","Value":"` + run + `"},{"Key":"updated","Value":"yes"}]}]`
	updated, err := cc.UpdateResource(ctx, &cloudcontrol.UpdateResourceInput{
		TypeName: aws.String(typeName), Identifier: aws.String(id), PatchDocument: aws.String(patch), ClientToken: aws.String(token()),
	})
	if err != nil {
		return fmt.Errorf("UpdateResource: %w", err)
	}
	if _, err := await(ctx, cc, updated.ProgressEvent); err != nil {
		return fmt.Errorf("update: %w", err)
	}
	fmt.Printf("updated in %s\n", time.Since(start).Round(time.Millisecond))
	if err := show(ctx, cc, typeName, id, "after update"); err != nil {
		return err
	}

	start = time.Now()
	deleted, err := cc.DeleteResource(ctx, &cloudcontrol.DeleteResourceInput{
		TypeName: aws.String(typeName), Identifier: aws.String(id), ClientToken: aws.String(token()),
	})
	if err != nil {
		return fmt.Errorf("DeleteResource: %w", err)
	}
	if _, err := await(ctx, cc, deleted.ProgressEvent); err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	fmt.Printf("deleted in %s\n", time.Since(start).Round(time.Millisecond))
	gone := id
	id = ""
	_, err = cc.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String(typeName), Identifier: aws.String(gone)})
	var nf *cctypes.ResourceNotFoundException
	fmt.Printf("GetResource after delete: not-found=%v (%v)\n", errors.As(err, &nf), err)
	return nil
}

// await polls the request until it settles. Cloud Control is asynchronous: every mutation returns a
// request token, not a result.
func await(ctx context.Context, cc *cloudcontrol.Client, ev *cctypes.ProgressEvent) (*cctypes.ProgressEvent, error) {
	for {
		switch ev.OperationStatus {
		case cctypes.OperationStatusSuccess:
			return ev, nil
		case cctypes.OperationStatusFailed, cctypes.OperationStatusCancelComplete:
			return ev, fmt.Errorf("%s %s: %s", ev.OperationStatus, ev.ErrorCode, aws.ToString(ev.StatusMessage))
		}
		time.Sleep(2 * time.Second)
		out, err := cc.GetResourceRequestStatus(ctx, &cloudcontrol.GetResourceRequestStatusInput{RequestToken: ev.RequestToken})
		if err != nil {
			return nil, err
		}
		ev = out.ProgressEvent
	}
}

func show(ctx context.Context, cc *cloudcontrol.Client, typeName, id, when string) error {
	got, err := cc.GetResource(ctx, &cloudcontrol.GetResourceInput{TypeName: aws.String(typeName), Identifier: aws.String(id)})
	if err != nil {
		return fmt.Errorf("GetResource: %w", err)
	}
	var props map[string]any
	_ = json.Unmarshal([]byte(aws.ToString(got.ResourceDescription.Properties)), &props)
	keys := make([]string, 0, len(props))
	for k := range props {
		keys = append(keys, k)
	}
	fmt.Printf("  %s: %d properties returned: %s\n", when, len(props), strings.Join(sorted(keys), ", "))
	fmt.Printf("    CidrBlock=%v EnableDnsHostnames=%v Tags=%v\n", props["CidrBlock"], props["EnableDnsHostnames"], props["Tags"])
	return nil
}

func sorted(s []string) []string {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
	return s
}

func token() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
