//go:build live

// Package live runs the plugin against a real AWS account. It is the only suite that sees real Cloud Control handlers,
// eventual consistency and IAM, and the only one that costs anything if it leaks, so it refuses to run unless told
// which account it may use and the credentials really are that account.
//
//	INFRENA_AWS_LIVE_PROFILE=infrena-live INFRENA_AWS_LIVE_ACCOUNT=111111111111 go test -tags live -count=1 -v -timeout 30m ./live/
package live

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/infrena/infrena-provider-aws/internal/awsprov"
	"github.com/infrena/infrena-provider-aws/internal/awstest"
	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena/pkg/address"
	"github.com/infrena/infrena/pkg/plugintest"
	"github.com/infrena/infrena/pkg/provider"
	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/value"
)

const runTag = "infrena-live-run"

func s(v string) value.Value { return value.String(v, value.SourceExplicit) }
func n(v int64) value.Value  { return value.Int(v, value.SourceExplicit) }
func m(kv ...any) value.Value {
	items := map[string]value.Value{}
	for i := 0; i+1 < len(kv); i += 2 {
		items[kv[i].(string)] = kv[i+1].(value.Value)
	}
	return value.Map(items, value.SourceExplicit)
}
func l(items ...value.Value) value.Value { return value.List(items, value.SourceExplicit) }

// generatePassword returns a fresh RDS master password for this run: random, long enough for RDS's minimum, and
// drawn from a set that avoids the three characters MasterUserPassword's description forbids ("/", "\"", "@").
func generatePassword(t *testing.T) string {
	t.Helper()
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%^&*()-_=+"
	buf := make([]byte, 24)
	for i := range buf {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			t.Fatal(err)
		}
		buf[i] = alphabet[n.Int64()]
	}
	return string(buf)
}

// guard skips without the variables, and refuses when the profile is not the named account.
func guard(t *testing.T) (profile, region string) {
	t.Helper()
	profile, account := os.Getenv("INFRENA_AWS_LIVE_PROFILE"), os.Getenv("INFRENA_AWS_LIVE_ACCOUNT")
	if profile == "" || account == "" {
		t.Skip("set INFRENA_AWS_LIVE_PROFILE and INFRENA_AWS_LIVE_ACCOUNT to run against real AWS")
	}
	region = os.Getenv("INFRENA_AWS_LIVE_REGION")
	if region == "" {
		region = "us-east-1"
	}
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(profile), config.WithRegion(region))
	if err != nil {
		t.Fatal(err)
	}
	id, err := sts.NewFromConfig(cfg).GetCallerIdentity(context.Background(), &sts.GetCallerIdentityInput{})
	if err != nil {
		t.Fatal(err)
	}
	if got := aws.ToString(id.Account); got != account {
		t.Fatalf("profile %q is account %s, not INFRENA_AWS_LIVE_ACCOUNT=%s: refusing to create anything", profile, got, account)
	}
	if strings.HasSuffix(aws.ToString(id.Arn), ":root") {
		t.Fatalf("profile %q holds root credentials: use a least-privilege IAM identity (live/README.md)", profile)
	}
	return profile, region
}

func configure(t *testing.T, profile, region string) (provider.Provider, *catalog.Catalog) {
	t.Helper()
	host, err := plugintest.Open(context.Background(), awsprov.NewPlugin(), t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = host.Close() })
	cat, err := catalog.Embedded()
	if err != nil {
		t.Fatal(err)
	}
	role := awstest.TypeFor(t, cat, "AWS::IAM::Role").Name
	bucket := awstest.TypeFor(t, cat, "AWS::S3::Bucket").Name
	repo := awstest.TypeFor(t, cat, "AWS::ECR::Repository").Name
	cluster := awstest.TypeFor(t, cat, "AWS::ECS::Cluster").Name
	dbInstance := awstest.TypeFor(t, cat, "AWS::RDS::DBInstance").Name
	dbSubnetGroup := awstest.TypeFor(t, cat, "AWS::RDS::DBSubnetGroup").Name
	dbParameterGroup := awstest.TypeFor(t, cat, "AWS::RDS::DBParameterGroup").Name
	prov, err := host.Configure(provider.Config{Instance: "live", Values: map[string]value.Value{
		"profile":          s(profile),
		"discover_regions": l(s(region)),
		"discover_types": l(s("aws.vpc"), s("aws.subnet"), s("aws.securitygroup"), s(role), s(bucket), s(repo), s(cluster),
			s(dbInstance), s(dbSubnetGroup), s(dbParameterGroup)),
	}})
	if err != nil {
		t.Fatal(err)
	}
	return prov, cat
}

// converged fails unless every configured value reads back equal: otherwise the next plan is not clean.
func converged(t *testing.T, attrs map[string]value.Value, st *resource.ResourceState) {
	t.Helper()
	for name, v := range attrs {
		if got := st.Attributes[name]; !got.Equal(v) {
			t.Errorf("%s %s: %s reads back as %v, configured %v", st.Type, st.ProviderID, name, got, v)
		}
	}
}

func create(t *testing.T, prov provider.Provider, typ string, attrs map[string]value.Value) *resource.ResourceState {
	t.Helper()
	start := time.Now()
	st, err := prov.Create(context.Background(), &resource.DesiredResource{Address: address.Address{Name: typ}, Type: typ, Attrs: attrs})
	if err != nil {
		t.Fatalf("create %s: %v", typ, err)
	}
	t.Logf("created %s %s in %s", typ, st.ProviderID, time.Since(start).Round(time.Millisecond))
	t.Cleanup(func() { _ = prov.Delete(context.Background(), st) })
	converged(t, attrs, st)
	read, err := prov.Read(context.Background(), st)
	if err != nil || read == nil {
		t.Fatalf("read %s straight after create = %v, %v", st.ProviderID, read, err)
	}
	converged(t, attrs, read)
	return read
}

func update(t *testing.T, prov provider.Provider, st *resource.ResourceState, changes map[string]value.Value) *resource.ResourceState {
	t.Helper()
	attrs := map[string]value.Value{}
	for k, v := range st.Attributes {
		attrs[k] = v
	}
	for k, v := range changes {
		attrs[k] = v
	}
	got, err := prov.Update(context.Background(), st, &resource.DesiredResource{Address: address.Address{Name: st.Type}, Type: st.Type, Attrs: attrs})
	if err != nil {
		t.Fatalf("update %s: %v", st.ProviderID, err)
	}
	converged(t, changes, got)
	return got
}

func TestTheLifecycleAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))

	vpc := create(t, prov, "aws.vpc", map[string]value.Value{"region": s(region), "CidrBlock": s("10.99.0.0/16"), "Tags": tags})
	vpcID := vpc.Attributes["VpcId"]
	subnet := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.99.1.0/24"), "AvailabilityZone": s(region + "a"), "Tags": tags})
	ingress := l(m("ip_protocol", s("tcp"), "from_port", n(443), "to_port", n(443), "cidr_ip", s("0.0.0.0/0")))
	sg := create(t, prov, "aws.securitygroup", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"GroupDescription": s("infrena live " + run), "SecurityGroupIngress": ingress, "Tags": tags})
	roleType := awstest.TypeFor(t, cat, "AWS::IAM::Role").Name
	// JSON text, as infrena's compiler requires for this object-or-string property, spaced unlike AWS's answer.
	policy := s(`{
  "Version": "2012-10-17",
  "Statement": [ { "Effect": "Allow", "Principal": { "Service": "ec2.amazonaws.com" }, "Action": "sts:AssumeRole" } ]
}`)
	role := create(t, prov, roleType, map[string]value.Value{"RoleName": s("infrena-live-" + run), "AssumeRolePolicyDocument": policy, "Tags": tags})

	vpc = update(t, prov, vpc, map[string]value.Value{"EnableDnsHostnames": value.Bool(true, value.SourceExplicit)})
	sg = update(t, prov, sg, map[string]value.Value{"SecurityGroupIngress": l(
		m("ip_protocol", s("tcp"), "from_port", n(443), "to_port", n(443), "cidr_ip", s("0.0.0.0/0")),
		m("ip_protocol", s("tcp"), "from_port", n(80), "to_port", n(80), "cidr_ip", s("0.0.0.0/0")))})
	role = update(t, prov, role, map[string]value.Value{"MaxSessionDuration": n(7200)})

	var everything []string
	for _, typ := range cat.Types {
		everything = append(everything, typ.Name)
	}
	found, err := prov.Discover(ctx, provider.DiscoverRequest{Types: everything})
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, r := range found {
		seen[r.ProviderID] = true
	}
	for _, st := range []*resource.ResourceState{vpc, subnet, sg, role} {
		if !seen[st.ProviderID] {
			t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", st.ProviderID)
		}
	}
	if imported, err := prov.Import(ctx, "aws.vpc", vpc.ProviderID); err != nil || !imported.Attributes["CidrBlock"].Equal(s("10.99.0.0/16")) {
		t.Errorf("import %s = %v, %v", vpc.ProviderID, imported, err)
	}

	for _, st := range []*resource.ResourceState{role, sg, subnet, vpc} {
		start := time.Now()
		if err := prov.Delete(ctx, st); err != nil {
			t.Fatalf("delete %s: %v", st.ProviderID, err)
		}
		t.Logf("deleted %s in %s", st.ProviderID, time.Since(start).Round(time.Millisecond))
		if got, err := prov.Read(ctx, st); err != nil || got != nil {
			t.Errorf("read %s after delete = %v, %v; want gone", st.ProviderID, got, err)
		}
	}
}

// TestStorageAndContainersAgainstRealAWS covers an S3 bucket, an ECR repository and an ECS cluster: three more
// types, each with nested configuration that exercises reconciliation, kept in a separate test so a run can target
// it alone with -run. Everything here stays inside the AWS free tier: no objects go into the bucket or images into
// the repository (an empty bucket/repository is required to delete it), and the ECS cluster setting is left
// "disabled" the whole time, since enabling Container Insights is not free.
func TestStorageAndContainersAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))
	name := "infrena-live-" + run

	bucketType := awstest.TypeFor(t, cat, "AWS::S3::Bucket").Name
	versioning := m("status", s("Enabled"))
	lifecycle := m("rules", l(m("id", s("expire-noncurrent"), "status", s("Enabled"), "expiration_in_days", n(1))))
	bucket := create(t, prov, bucketType, map[string]value.Value{"region": s(region), "BucketName": s(name),
		"VersioningConfiguration": versioning, "LifecycleConfiguration": lifecycle, "Tags": tags})

	repoType := awstest.TypeFor(t, cat, "AWS::ECR::Repository").Name
	// JSON text for a nested string property, spaced unlike a minified document, the same reasoning as the role's
	// AssumeRolePolicyDocument above (a real user would not hand-write compact JSON).
	lifecyclePolicyText := s(`{
  "rules": [
    {
      "rulePriority": 1,
      "description": "Expire untagged images",
      "selection": { "tagStatus": "untagged", "countType": "imageCountMoreThan", "countNumber": 1 },
      "action": { "type": "expire" }
    }
  ]
}`)
	repo := create(t, prov, repoType, map[string]value.Value{"region": s(region), "RepositoryName": s(name),
		"ImageTagMutability":         s("MUTABLE"),
		"ImageScanningConfiguration": m("scan_on_push", value.Bool(true, value.SourceExplicit)),
		"LifecyclePolicy":            m("lifecycle_policy_text", lifecyclePolicyText),
		"Tags":                       tags})

	clusterType := awstest.TypeFor(t, cat, "AWS::ECS::Cluster").Name
	settings := l(m("name", s("containerInsights"), "value", s("disabled")))
	cluster := create(t, prov, clusterType, map[string]value.Value{"region": s(region), "ClusterName": s(name),
		"ClusterSettings": settings, "Tags": tags})

	newLifecycle := m("rules", l(m("id", s("expire-noncurrent"), "status", s("Enabled"), "expiration_in_days", n(3))))
	bucket = update(t, prov, bucket, map[string]value.Value{"LifecycleConfiguration": newLifecycle})
	repo = update(t, prov, repo, map[string]value.Value{"ImageTagMutability": s("IMMUTABLE")})
	cluster = update(t, prov, cluster, map[string]value.Value{"Tags": m(runTag, s(run), "Purpose", s("live-test"))})

	var everything []string
	for _, typ := range cat.Types {
		everything = append(everything, typ.Name)
	}
	found, err := prov.Discover(ctx, provider.DiscoverRequest{Types: everything})
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, r := range found {
		seen[r.ProviderID] = true
	}
	for _, st := range []*resource.ResourceState{bucket, repo, cluster} {
		if !seen[st.ProviderID] {
			t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", st.ProviderID)
		}
	}

	for _, st := range []*resource.ResourceState{cluster, repo, bucket} {
		start := time.Now()
		if err := prov.Delete(ctx, st); err != nil {
			t.Fatalf("delete %s: %v", st.ProviderID, err)
		}
		t.Logf("deleted %s in %s", st.ProviderID, time.Since(start).Round(time.Millisecond))
		if got, err := prov.Read(ctx, st); err != nil || got != nil {
			t.Errorf("read %s after delete = %v, %v; want gone", st.ProviderID, got, err)
		}
	}
}

// TestDatabasesAgainstRealAWS covers RDS: a DB subnet group, a DB parameter group and a DB instance, in a
// dedicated VPC with two subnets in different Availability Zones (a DB subnet group must span at least two AZs).
// The instance is the smallest free-tier-eligible shape (db.t3.micro, 20 GiB gp3, single-AZ, no Performance
// Insights, no public access) but this is NOT free if the account's free tier is already used up elsewhere —
// see live/README.md.
//
// AWS::RDS::DBInstance's create, update and delete handlers are each registered in the catalog with a
// 2160-minute (36-hour) timeoutInMinutes (internal/ccprov/catalog_test.go, taken from
// schemas/CloudformationSchema.zip's aws-rds-dbinstance.json) — that is Cloud Control's own outer ceiling on the
// request, not a prediction of how long a real create takes. A real DB instance create typically takes AWS 5 to
// 10 minutes and a delete several minutes more, but a run needs `go test`'s own -timeout raised well past this
// suite's usual few minutes to have room for that (see live/README.md's "Running it" section) — the ceiling above
// is not itself what to set -timeout to.
//
// AWS::RDS::DBInstance has no SkipFinalSnapshot (or similar) property in its schema (checked in
// docs/reference/rds/rds-dbinstance.md and schemas/CloudformationSchema.zip's aws-rds-dbinstance.json
// writeOnlyProperties/createOnlyProperties: neither lists it), so there is nothing to set here to avoid a final
// snapshot on delete — Cloud Control's DeleteResource for this type has no way to be told to take one either.
//
// Kept separate from the other tests so a run can target it alone with -run TestDatabasesAgainstRealAWS.
func TestDatabasesAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))
	name := "infrena-live-" + run

	vpc := create(t, prov, "aws.vpc", map[string]value.Value{"region": s(region), "CidrBlock": s("10.98.0.0/16"), "Tags": tags})
	vpcID := vpc.Attributes["VpcId"]
	subnetA := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.98.1.0/24"), "AvailabilityZone": s(region + "a"), "Tags": tags})
	subnetB := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.98.2.0/24"), "AvailabilityZone": s(region + "b"), "Tags": tags})

	subnetGroupType := awstest.TypeFor(t, cat, "AWS::RDS::DBSubnetGroup").Name
	subnetGroup := create(t, prov, subnetGroupType, map[string]value.Value{"region": s(region),
		"DBSubnetGroupName": s(name), "DBSubnetGroupDescription": s("infrena live " + run),
		"SubnetIds": l(subnetA.Attributes["SubnetId"], subnetB.Attributes["SubnetId"]), "Tags": tags})

	parameterGroupType := awstest.TypeFor(t, cat, "AWS::RDS::DBParameterGroup").Name
	// The parameter group's family and the instance's EngineVersion must name the same major version, and
	// both are pinned here to an EXACT minor version. Two runs on 2026-09-16 showed why neither looser form
	// works: configuring "17" read back as "17.9", which plans a change on every apply, and leaving
	// EngineVersion unset made AWS choose its current default major version, which then refused this
	// parameter group ("can't be used for this instance. Use a parameter group with DBParameterGroupFamily
	// postgres18"). When AWS retires 17.9 this test fails with that same clear message: pick a current pair
	// from `aws rds describe-db-engine-versions --engine postgres` and bump both lines together.
	parameterGroup := create(t, prov, parameterGroupType, map[string]value.Value{"region": s(region),
		"DBParameterGroupName": s(name), "Description": s("infrena live " + run), "Family": s("postgres17"), "Tags": tags})

	instanceType := awstest.TypeFor(t, cat, "AWS::RDS::DBInstance").Name
	instance := create(t, prov, instanceType, map[string]value.Value{
		"region":                s(region),
		"DBInstanceIdentifier":  s(name),
		"Engine":                s("postgres"),
		"EngineVersion":         s("17.9"),
		"DBInstanceClass":       s("db.t3.micro"),
		"AllocatedStorage":      s("20"),
		"StorageType":           s("gp3"),
		"MasterUsername":        s("infrenaadmin"),
		"MasterUserPassword":    s(generatePassword(t)),
		"DBSubnetGroupName":     subnetGroup.Attributes["DBSubnetGroupName"],
		"DBParameterGroupName":  parameterGroup.Attributes["DBParameterGroupName"],
		"PubliclyAccessible":    value.Bool(false, value.SourceExplicit),
		"MultiAZ":               value.Bool(false, value.SourceExplicit),
		"BackupRetentionPeriod": n(0),
		"Tags":                  tags,
	})

	// Update something cheap and in place: RDS applies a tag change immediately through AddTagsToResource, with
	// no reboot and no storage or instance-class change (unlike updating DBInstanceClass or AllocatedStorage,
	// which this test deliberately leaves alone).
	instance = update(t, prov, instance, map[string]value.Value{"Tags": m(runTag, s(run), "Purpose", s("live-test"))})

	var everything []string
	for _, typ := range cat.Types {
		everything = append(everything, typ.Name)
	}
	found, err := prov.Discover(ctx, provider.DiscoverRequest{Types: everything})
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]bool{}
	for _, r := range found {
		seen[r.ProviderID] = true
	}
	for _, st := range []*resource.ResourceState{instance, subnetGroup, parameterGroup, subnetB, subnetA, vpc} {
		if !seen[st.ProviderID] {
			t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", st.ProviderID)
		}
	}

	// Dependency order: the instance depends on the subnet group, the parameter group and both subnets; the
	// subnet group depends on both subnets; the subnets and the VPC come last. This is also the order create's
	// t.Cleanup deletes would run in (LIFO from creation order), so a failure partway through still tears down
	// cleanly.
	for _, st := range []*resource.ResourceState{instance, subnetGroup, parameterGroup, subnetB, subnetA, vpc} {
		start := time.Now()
		if err := prov.Delete(ctx, st); err != nil {
			t.Fatalf("delete %s: %v", st.ProviderID, err)
		}
		t.Logf("deleted %s in %s", st.ProviderID, time.Since(start).Round(time.Millisecond))
		if got, err := prov.Read(ctx, st); err != nil || got != nil {
			t.Errorf("read %s after delete = %v, %v; want gone", st.ProviderID, got, err)
		}
	}
}

// TestSweepLeftovers deletes what a crashed run left: anything tagged by this suite more than an hour ago.
func TestSweepLeftovers(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	cutoff := time.Now().Add(-time.Hour).Unix()
	// Buckets and repositories sweep safely without checking for emptiness: this suite never puts objects or
	// images in them, so anything it tagged is always empty. DB instances, subnet groups and parameter groups
	// come first: an instance depends on the other two plus subnets, and a subnet group depends on the subnets.
	order := []string{
		awstest.TypeFor(t, cat, "AWS::RDS::DBInstance").Name,
		awstest.TypeFor(t, cat, "AWS::RDS::DBSubnetGroup").Name,
		awstest.TypeFor(t, cat, "AWS::RDS::DBParameterGroup").Name,
		awstest.TypeFor(t, cat, "AWS::ECS::Cluster").Name,
		awstest.TypeFor(t, cat, "AWS::ECR::Repository").Name,
		awstest.TypeFor(t, cat, "AWS::S3::Bucket").Name,
		awstest.TypeFor(t, cat, "AWS::IAM::Role").Name,
		"aws.securitygroup", "aws.subnet", "aws.vpc",
	}
	found, err := prov.Discover(context.Background(), provider.DiscoverRequest{Types: order})
	if err != nil {
		t.Fatal(err)
	}
	for _, typ := range order {
		for _, r := range found {
			if r.Type != typ {
				continue
			}
			tags, _ := r.Attributes["Tags"].Raw.(map[string]value.Value)
			started, err := strconv.ParseInt(fmt.Sprint(tags[runTag].Raw), 10, 64)
			if err != nil || started >= cutoff {
				continue
			}
			t.Logf("deleting %s %s", r.Type, r.ProviderID)
			if err := prov.Delete(context.Background(), &resource.ResourceState{Type: r.Type, ProviderID: r.ProviderID, Attributes: r.Attributes}); err != nil {
				t.Error(err)
			}
		}
	}
}
