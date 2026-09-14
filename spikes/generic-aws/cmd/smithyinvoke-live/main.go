// Command smithyinvoke-live is THROWAWAY SPIKE CODE: the model-driven invoker against REAL EC2, read-only.
// The fake proves the mechanics; only AWS proves the signing and wire format are accepted.
//
//	go run ./cmd/smithyinvoke-live -model smithyinvoke/testdata/ec2.json -profile infrata -region us-east-1
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/infrata/infrata-provider-aws/spikes/genericaws/smithyinvoke"
)

func main() {
	model := flag.String("model", "smithyinvoke/testdata/ec2.json", "EC2 Smithy model")
	profile := flag.String("profile", "infrata", "AWS profile")
	region := flag.String("region", "us-east-1", "AWS region")
	flag.Parse()

	ctx := context.Background()
	m, err := smithyinvoke.Load(*model)
	fail(err)
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(*profile), config.WithRegion(*region))
	fail(err)
	t := smithyinvoke.Target{
		Endpoint: fmt.Sprintf("https://ec2.%s.amazonaws.com", *region), Region: *region,
		SigningName: "ec2", Credentials: cfg.Credentials,
	}

	for _, call := range []struct {
		op string
		in map[string]any
	}{
		{"DescribeVpcs", map[string]any{}},
		{"DescribeAvailabilityZones", map[string]any{"Filters": []any{map[string]any{"Name": "state", "Values": []any{"available"}}}}},
		{"DescribeVpcs", map[string]any{"VpcIds": []any{"vpc-0000000000000dead"}}},
	} {
		out, err := smithyinvoke.Invoke(ctx, m, t, call.op, call.in)
		if err != nil {
			fmt.Printf("%s -> error: %v\n", call.op, err)
			continue
		}
		fmt.Printf("%s -> %v\n", call.op, out)
	}
}

func fail(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
