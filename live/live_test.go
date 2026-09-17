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
	taskDefinition := awstest.TypeFor(t, cat, "AWS::ECS::TaskDefinition").Name
	ecsService := awstest.TypeFor(t, cat, "AWS::ECS::Service").Name
	logGroup := awstest.TypeFor(t, cat, "AWS::Logs::LogGroup").Name
	dbInstance := awstest.TypeFor(t, cat, "AWS::RDS::DBInstance").Name
	dbSubnetGroup := awstest.TypeFor(t, cat, "AWS::RDS::DBSubnetGroup").Name
	dbParameterGroup := awstest.TypeFor(t, cat, "AWS::RDS::DBParameterGroup").Name
	loadBalancer := awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::LoadBalancer").Name
	targetGroup := awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::TargetGroup").Name
	listener := awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::Listener").Name
	hostedZone := awstest.TypeFor(t, cat, "AWS::Route53::HostedZone").Name
	recordSet := awstest.TypeFor(t, cat, "AWS::Route53::RecordSet").Name
	function := awstest.TypeFor(t, cat, "AWS::Lambda::Function").Name
	prov, err := host.Configure(provider.Config{Instance: "live", Values: map[string]value.Value{
		"profile":          s(profile),
		"discover_regions": l(s(region)),
		"discover_types": l(s("aws.vpc"), s("aws.subnet"), s("aws.securitygroup"), s(role), s(bucket), s(repo), s(cluster),
			s(taskDefinition), s(ecsService), s(logGroup),
			s(dbInstance), s(dbSubnetGroup), s(dbParameterGroup), s(loadBalancer), s(targetGroup), s(listener),
			s(hostedZone), s(recordSet), s(function)),
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

// elbName fits a load balancer or target group name into what ELBv2 accepts: at most 32 characters, alphanumeric
// characters and hyphens only, no leading or trailing hyphen, and (for a load balancer) not beginning "internal-".
// "infrena-live-" plus a 10-digit unix time is 23 characters and "infrena-live-tg-" plus one is 26, so nothing is
// truncated today; the truncation is here so a longer prefix is cut to a valid name rather than rejected by AWS.
func elbName(prefix, run string) string {
	name := prefix + run
	if len(name) > 32 {
		name = name[:32]
	}
	return strings.TrimRight(name, "-")
}

// lingeringDependency reports whether AWS refused a delete because something still points at the resource. Deleting
// an ALB leaves its elastic network interfaces in the subnets for a while after Cloud Control reports the load
// balancer gone, so the security group and the subnets can be refused for a few minutes afterwards; a target group
// is refused while a listener still forwards to it. Only these are worth waiting out — anything else is a real
// failure, and matching on the message is deliberate: Cloud Control reports the downstream service's error text.
func lingeringDependency(err error) bool {
	text := err.Error()
	for _, marker := range []string{"DependencyViolation", "dependent object", "has dependencies", "currently in use", "ResourceInUse"} {
		if strings.Contains(text, marker) {
			return true
		}
	}
	return false
}

// deleteAndConfirm deletes st and fails unless the read-back afterwards finds nothing, waiting out a lingering
// dependency (see lingeringDependency) for up to six minutes. It is the ELB test's teardown step: the other tests
// delete resources nothing else holds a network interface in, so they call prov.Delete directly.
func deleteAndConfirm(t *testing.T, prov provider.Provider, st *resource.ResourceState) {
	t.Helper()
	ctx := context.Background()
	start := time.Now()
	deadline := start.Add(6 * time.Minute)
	for {
		err := prov.Delete(ctx, st)
		if err == nil {
			break
		}
		if !lingeringDependency(err) || time.Now().After(deadline) {
			t.Fatalf("delete %s: %v", st.ProviderID, err)
		}
		t.Logf("delete %s: %v; waiting for AWS to release it", st.ProviderID, err)
		time.Sleep(15 * time.Second)
	}
	t.Logf("deleted %s in %s", st.ProviderID, time.Since(start).Round(time.Millisecond))
	if got, err := prov.Read(ctx, st); err != nil || got != nil {
		t.Errorf("read %s after delete = %v, %v; want gone", st.ProviderID, got, err)
	}
}

// TestLoadBalancersAgainstRealAWS covers Elastic Load Balancing v2: an Application Load Balancer, a target group
// and a listener forwarding to it, in a dedicated VPC with two subnets in different Availability Zones (an ALB
// must span at least two) and its own security group, so nothing depends on the account's default VPC.
//
// NOT free tier: an ALB costs roughly $0.0225 an hour plus LCU-hours in us-east-1, billed per hour started, so one
// run costs well under a cent. See live/README.md.
//
// How long AWS takes: creating an ALB typically takes 2 to 4 minutes and deleting one 1 to 3, and the network
// interfaces it leaves behind can hold the security group and the subnets for a few minutes more (deleteAndConfirm
// waits that out). With the VPC, subnets, target group and listener either side of that, this test needs `go
// test`'s own -timeout at 30m; the other tests' usual few minutes is not enough. The catalog registers this type's
// create and update handlers with Cloud Control's own 2160-minute ceiling, which is not a prediction of either.
//
// Scheme is "internal", not "internet-facing", on purpose. AWS requires an internet gateway attached to the VPC
// before it will create an internet-facing load balancer, which would mean an AWS::EC2::InternetGateway plus an
// AWS::EC2::VPCGatewayAttachment here — and the attachment has to be deleted before the VPC, while the ALB's
// network interfaces are still lingering in it, so the one part of teardown most likely to need a retry would get
// another resource in the middle of it. An internal ALB exercises the same three ELBv2 types with no gateway at
// all, and is not reachable from the internet, which a live test creating a public endpoint otherwise would be.
//
// Kept separate from the other tests so a run can target it alone with -run TestLoadBalancersAgainstRealAWS.
func TestLoadBalancersAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))

	vpc := create(t, prov, "aws.vpc", map[string]value.Value{"region": s(region), "CidrBlock": s("10.97.0.0/16"), "Tags": tags})
	vpcID := vpc.Attributes["VpcId"]
	subnetA := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.97.1.0/24"), "AvailabilityZone": s(region + "a"), "Tags": tags})
	subnetB := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.97.2.0/24"), "AvailabilityZone": s(region + "b"), "Tags": tags})

	// Ingress on the listener's port, from the VPC's own range: an internal load balancer is only reachable from
	// inside the VPC, so a 0.0.0.0/0 rule would claim more than this test needs.
	ingress := l(m("ip_protocol", s("tcp"), "from_port", n(80), "to_port", n(80), "cidr_ip", s("10.97.0.0/16")))
	sg := create(t, prov, "aws.securitygroup", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"GroupDescription": s("infrena live " + run), "SecurityGroupIngress": ingress, "Tags": tags})

	// The load balancer's subnets and the target group's VpcId have to agree — an ALB can only forward to a target
	// group in the VPC its subnets are in — and no schema says so, so both are set explicitly from the VPC created
	// above rather than left to AWS. Subnets is insertionOrder: false in the schema, so AWS is free to return the
	// two in the other order; reconciliation reorders an unordered list back to the reference's order, which is why
	// create's converged check can assert it at all. IpAddressType is set because an internal ALB must be ipv4:
	// spelling out the value AWS would have chosen keeps it out of every later plan.
	lbType := awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::LoadBalancer").Name
	lb := create(t, prov, lbType, map[string]value.Value{
		"region":         s(region),
		"Name":           s(elbName("infrena-live-", run)),
		"Type":           s("application"),
		"Scheme":         s("internal"),
		"IpAddressType":  s("ipv4"),
		"Subnets":        l(subnetA.Attributes["SubnetId"], subnetB.Attributes["SubnetId"]),
		"SecurityGroups": l(sg.Attributes["GroupId"]),
		"Tags":           tags,
	})

	// Every health check value is configured, not defaulted. AWS fills in its own defaults for the ones a create
	// leaves out (path "/", port "traffic-port", interval 30, timeout 5, thresholds, matcher 200) and reports them
	// on every read, so a value left unset here would read back as something this test never asserted and a value
	// set to something AWS rewrites would plan a change forever. HealthCheckTimeoutSeconds must stay below
	// HealthCheckIntervalSeconds, including after the update below (5 < 10).
	tgType := awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::TargetGroup").Name
	tg := create(t, prov, tgType, map[string]value.Value{
		"region":                     s(region),
		"Name":                       s(elbName("infrena-live-tg-", run)),
		"Protocol":                   s("HTTP"),
		"Port":                       n(80),
		"VpcId":                      vpcID,
		"TargetType":                 s("instance"),
		"HealthCheckEnabled":         value.Bool(true, value.SourceExplicit),
		"HealthCheckProtocol":        s("HTTP"),
		"HealthCheckPath":            s("/"),
		"HealthCheckPort":            s("traffic-port"),
		"HealthCheckIntervalSeconds": n(30),
		"HealthCheckTimeoutSeconds":  n(5),
		"HealthyThresholdCount":      n(2),
		"UnhealthyThresholdCount":    n(3),
		"Matcher":                    m("http_code", s("200")),
		"Tags":                       tags,
	})
	// No targets are registered: an empty target group is legal, reports its targets as unhealthy to nobody, and
	// keeps this test to three ELBv2 resources with no EC2 instance to pay for or wait on.

	// The listener's Protocol and Port must suit the target group's — HTTP:80 to an HTTP:80 target group — and
	// again no schema couples them, so both ends are written out. AWS answers a read of a single-target-group
	// forward action with more than was sent (a ForwardConfig naming the same target group with weight 1, and an
	// Order): reconciliation drops keys AWS added that the reference does not have, so this reads back as the two
	// keys configured here.
	listenerType := awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::Listener").Name
	listener := create(t, prov, listenerType, map[string]value.Value{
		"region":          s(region),
		"LoadBalancerArn": lb.Attributes["LoadBalancerArn"],
		"Port":            n(80),
		"Protocol":        s("HTTP"),
		"DefaultActions":  l(m("type", s("forward"), "target_group_arn", tg.Attributes["TargetGroupArn"])),
		"Tags":            tags,
	})

	// Cheap, in-place updates only: ModifyTargetGroup changes a health check with no replacement and no traffic
	// impact, and a load balancer tag is an AddTags call. Deliberately not the load balancer's Subnets or Scheme —
	// Scheme is create-only, and changing Subnets moves the network interfaces AWS is still holding — and not the
	// listener's Protocol or Port, which are coupled to the target group's.
	tg = update(t, prov, tg, map[string]value.Value{"HealthCheckPath": s("/healthz"), "HealthCheckIntervalSeconds": n(10)})
	lb = update(t, prov, lb, map[string]value.Value{"Tags": m(runTag, s(run), "Purpose", s("live-test"))})

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
	for _, st := range []*resource.ResourceState{lb, tg, sg, subnetB, subnetA, vpc} {
		if !seen[st.ProviderID] {
			t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", st.ProviderID)
		}
	}
	// The listener is reported, not asserted. Its list handler's schema is `oneOf` [LoadBalancerArn | ListenerArns]
	// with no top-level `required`, and ListNeedsModel (internal/cfn/schema.go) only looks at a top-level
	// `required` — so the catalog marks listeners plainly listable and discovery calls ListResources for them with
	// no resource model. Whether AWS accepts that has never been checked against real Cloud Control. If this line
	// prints, the generator's parent-resource test is what needs changing, not this test: record it in the plan's
	// Verification log and tell James.
	if !seen[listener.ProviderID] {
		t.Logf("discovery did not find %s: check stderr for a ListResources failure on %s, and see the comment above",
			listener.ProviderID, listenerType)
	}

	// Dependency order: the listener forwards to the target group and sits on the load balancer; the load balancer
	// holds network interfaces in the subnets and uses the security group; the target group belongs to the VPC.
	// This is also the order create's t.Cleanup deletes would run in (LIFO from creation order), so a failure
	// partway through still tears down cleanly — though without deleteAndConfirm's wait, so a cleanup delete of the
	// security group or the subnets can still lose a race with the load balancer's network interfaces and leave
	// them for TestSweepLeftovers.
	for _, st := range []*resource.ResourceState{listener, lb, tg, sg, subnetB, subnetA, vpc} {
		deleteAndConfirm(t, prov, st)
	}
}

// TestDNSAgainstRealAWS covers Route 53: a private hosted zone this test creates, associated with a VPC it also
// creates, and two record sets inside that zone. Route 53 is a GLOBAL type in this plugin (gen/overlay.yaml's
// `global:` list has "AWS::Route53::*"): aws.hostedzone and aws.recordset take no region attribute and their
// provider IDs start "global/", so region is never set on either below.
//
// The zone is PRIVATE, not public: AWS::Route53::HostedZone's VPC definition in
// schemas/CloudformationSchema.zip's aws-route53-hostedzone.json says "For public hosted zones, omit VPCs,
// VPCId, and VPCRegion" — associating a VPC is what makes a zone private. Nothing is published to the internet
// and no registrar is involved. The domain itself cannot collide with anything real:
// "infrena-live-<unix time>.internal" is a fresh name each run in the .internal TLD, which is reserved (RFC
// 8375) and cannot be registered on the public internet even by accident.
//
// Both AWS::Route53::HostedZone and AWS::Route53::RecordSet declare a propertyTransform on Name in their
// schemas — CloudFormation's own drift-detection normalisation, which this plugin's reconciliation
// (internal/ccprov/reconcile.go) does not know about — so Name is configured below already in the form AWS
// actually stores, rather than quietly skipping the convergence assertion:
//   - HostedZone: `Name $OR $join([Name, "."]) $OR $lowercase(Name) $OR $lowercase($join([Name, "."]))` — AWS
//     lowercases and appends a trailing dot. zoneName is already lowercase with a trailing dot, so the
//     transform is a no-op and what's configured reads back unchanged.
//   - RecordSet: `$lowercase($replace(Name, /(.*)\.$/, "$1"))` — the opposite: AWS lowercases and STRIPS a
//     trailing dot. Record names below are already lowercase with no trailing dot, for the same reason.
//
// A hosted zone create/read/update/delete typically takes AWS a few seconds, and so do record sets — nothing
// like the RDS instance or the ALB above. -timeout 10m is generous headroom alongside the VPC this test also
// creates and destroys.
//
// Kept separate from the other tests so a run can target it alone with -run TestDNSAgainstRealAWS.
func TestDNSAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))
	zoneName := "infrena-live-" + run + ".internal." // lowercase, trailing dot: see the propertyTransform comment above

	vpc := create(t, prov, "aws.vpc", map[string]value.Value{"region": s(region), "CidrBlock": s("10.96.0.0/16"), "Tags": tags})
	vpcID := vpc.Attributes["VpcId"]

	// VPCs is a list of {VPCId, VPCRegion} (schema names); reconciliation matches a configured key against a
	// property's snake_case form (internal/ccprov/reconcile.go's matchProp), so "vpc_id"/"vpc_region" here
	// match VPCId/VPCRegion the same way the ELB tests above write "http_code" for Matcher's HttpCode.
	zoneType := awstest.TypeFor(t, cat, "AWS::Route53::HostedZone").Name
	zone := create(t, prov, zoneType, map[string]value.Value{
		"Name":             s(zoneName),
		"HostedZoneConfig": m("comment", s("infrena live "+run)),
		"VPCs":             l(m("vpc_id", vpcID, "vpc_region", s(region))),
		"HostedZoneTags":   tags,
	})
	zoneID := zone.Attributes["Id"]

	// HostedZoneId is set explicitly on every record from the zone just created, rather than left to
	// HostedZoneName, and each record's Name and Type are written out in full — none of this is coupled by the
	// schema, the same reasoning as the load balancer/target group/listener coupling above.
	recordType := awstest.TypeFor(t, cat, "AWS::Route53::RecordSet").Name
	aName := "www." + strings.TrimSuffix(zoneName, ".") // no trailing dot: see the propertyTransform comment above
	aRecord := create(t, prov, recordType, map[string]value.Value{
		"Name":            s(aName),
		"Type":            s("A"),
		"HostedZoneId":    zoneID,
		"ResourceRecords": l(s("192.0.2.1")),
		"TTL":             s("300"),
	})

	txtName := "txt." + strings.TrimSuffix(zoneName, ".")
	txtRecord := create(t, prov, recordType, map[string]value.Value{
		"Name":         s(txtName),
		"Type":         s("TXT"),
		"HostedZoneId": zoneID,
		// A TXT value must itself be wrapped in quotes: that is Route 53's own format for this record type, not
		// something this plugin or AWS adds.
		"ResourceRecords": l(s(`"infrena live test"`)),
		"TTL":             s("300"),
	})

	// Update something cheap and in place on each resource type: the A record's TTL, and the zone's comment.
	aRecord = update(t, prov, aRecord, map[string]value.Value{"TTL": s("600")})
	zone = update(t, prov, zone, map[string]value.Value{"HostedZoneConfig": m("comment", s("infrena live "+run+" updated"))})

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
	if !seen[zone.ProviderID] {
		t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", zone.ProviderID)
	}
	// Record sets are never discovered, and that is correct rather than a gap. AWS::Route53::RecordSet's list
	// handler needs a parent hosted zone (its handlerSchema is `oneOf` [HostedZoneId | HostedZoneName], not a
	// top-level `required`), so the catalog marks it as needing a parent since `3d0f88f` and discovery skips it,
	// naming it once on stderr. A run on 2026-09-17 printed exactly that: "not discovering
	// aws.elasticloadbalancingv2.listener, aws.recordset: listing them needs a parent resource, which discovery
	// does not have". Left as a log rather than an assertion on absence: what matters is that discovery does not
	// FAIL on these, which the run above proves, and asserting a resource stays unfound would pass for the wrong
	// reason the day someone teaches discovery to walk parents.
	for _, st := range []*resource.ResourceState{aRecord, txtRecord} {
		t.Logf("discovery of %s %s: seen=%v (see the comment above; not asserted)", st.Type, st.ProviderID, seen[st.ProviderID])
	}

	// Dependency order: the record sets before the hosted zone (DeleteHostedZone refuses a zone that still
	// holds anything but its default NS/SOA records), the hosted zone before the VPC it is associated with.
	// deleteAndConfirm wraps only the VPC delete: disassociating a private zone happens as part of deleting the
	// zone itself, so a lingering association afterward is unlikely, but the wait costs nothing if a run
	// disagrees.
	for _, st := range []*resource.ResourceState{aRecord, txtRecord, zone} {
		start := time.Now()
		if err := prov.Delete(ctx, st); err != nil {
			t.Fatalf("delete %s: %v", st.ProviderID, err)
		}
		t.Logf("deleted %s in %s", st.ProviderID, time.Since(start).Round(time.Millisecond))
		if got, err := prov.Read(ctx, st); err != nil || got != nil {
			t.Errorf("read %s after delete = %v, %v; want gone", st.ProviderID, got, err)
		}
	}
	deleteAndConfirm(t, prov, vpc)
}

// TestContainerServicesAgainstRealAWS covers ECS's application layer above the bare cluster
// TestStorageAndContainersAgainstRealAWS already exercises: a CloudWatch Logs log group, a Fargate task
// definition and a service running on it, in a dedicated VPC with two subnets in different Availability
// Zones and its own security group, so nothing depends on the account's default VPC or its default
// cluster.
//
// AWS::ECS::TaskDefinition is create-only in effect, not just in name. Its schema
// (schemas/CloudformationSchema.zip's aws-ecs-taskdefinition.json) lists every substantive property —
// Family, ContainerDefinitions, Cpu, Memory, NetworkMode, RequiresCompatibilities, ExecutionRoleArn,
// TaskRoleArn, Volumes and the rest — in createOnlyProperties; only Tags is left out, and the update
// handler's own permissions are exactly the three tag actions (ecs:TagResource, ecs:UntagResource,
// ecs:ListTagsForResource). The catalog's HasUpdate is true for this type and Cloud Control's
// UpdateResource is reachable, but calling it with anything but a tag change is, in AWS's own model,
// "register a new revision" — a Create, not an Update. So this test never calls update() on the task
// definition; the cheap in-place update below runs on the service instead, the same shape of update
// every other test in this file makes.
//
// Cost: the log group and the task definition are free. The service is created with DesiredCount: 0.
// ECS::Service's schema puts no minimum on DesiredCount, and CreateService is documented to accept 0 for
// a brand-new service — a service with no running tasks costs nothing and starts nothing. Confirmed
// against real AWS on 2026-09-17: AWS accepted DesiredCount: 0 on create and scheduled no task, so this
// test bills nothing. If that ever changes and 0 is rejected, change DesiredCount to 1 and update
// live/README.md's cost section to match — a running Fargate task on this shape bills a few cents an
// hour, not "free".
//
// What's expected to come back rewritten, and how each is handled: TaskDefinitionArn gains the revision
// suffix AWS assigns (":1" on a first register) — it's a computed attribute, never part of what's
// configured, so converged() never asserts it. ContainerDefinitions is asserted in full: reconciliation
// (CLAUDE.md, internal/ccprov/reconcile.go) drops keys AWS added that the reference doesn't have, so a
// create that configures only the fields this test cares about (Name, Image, Essential, Command,
// LogConfiguration) should read back unchanged, the same way the load balancer test's Matcher does. The
// container's LogConfiguration.Options map is opaque — schemas/CloudformationSchema.zip declares it
// patternProperties with no properties, oneOf/anyOf/allOf — so it is copied exactly and its keys are
// written in AWS's own kebab-case spelling ("awslogs-group", not "awslogs_group"), never translated.
//
// A log group is created explicitly (LogGroupName, RetentionInDays) rather than relying on the
// container's awslogs-create-group option, so this test owns the group's lifecycle and deletes it itself
// instead of leaving a group AWS auto-created behind.
//
// The execution role is a plain IAM role trusted by ecs-tasks.amazonaws.com with the AWS managed policy
// AmazonECSTaskExecutionRolePolicy attached via ManagedPolicyArns — IAM is global, so it takes no region
// attribute, the same as the role TestTheLifecycleAgainstRealAWS creates.
//
// No inbound rule is configured on the security group: DesiredCount: 0 means no task, and therefore no
// network interface, is ever created by this test, so there is nothing for an ingress rule to reach.
// GroupDescription is the only required attribute on aws.securitygroup.
//
// How long AWS takes: a log group, a task definition register/deregister and an ECS cluster are each a
// few seconds. A service create/delete at DesiredCount: 0 involves no task placement, so it is fast too,
// but ECS can still take a little while to report a service fully drained even with nothing running.
// -timeout 20m is generous headroom alongside the VPC this test also creates and destroys; raise it if a
// run shows the service delete waiting out deleteAndConfirm's retries.
//
// Kept separate from the other tests so a run can target it alone with -run TestContainerServicesAgainstRealAWS.
func TestContainerServicesAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))
	name := "infrena-live-" + run

	vpc := create(t, prov, "aws.vpc", map[string]value.Value{"region": s(region), "CidrBlock": s("10.95.0.0/16"), "Tags": tags})
	vpcID := vpc.Attributes["VpcId"]
	subnetA := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.95.1.0/24"), "AvailabilityZone": s(region + "a"), "Tags": tags})
	subnetB := create(t, prov, "aws.subnet", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"CidrBlock": s("10.95.2.0/24"), "AvailabilityZone": s(region + "b"), "Tags": tags})
	sg := create(t, prov, "aws.securitygroup", map[string]value.Value{"region": s(region), "VpcId": vpcID,
		"GroupDescription": s("infrena live " + run), "Tags": tags})

	logGroupType := awstest.TypeFor(t, cat, "AWS::Logs::LogGroup").Name
	logGroup := create(t, prov, logGroupType, map[string]value.Value{"region": s(region),
		"LogGroupName": s("/ecs/" + name), "RetentionInDays": n(1), "Tags": tags})

	roleType := awstest.TypeFor(t, cat, "AWS::IAM::Role").Name
	// JSON text, spaced unlike AWS's answer, the same reasoning as the role TestTheLifecycleAgainstRealAWS
	// creates: infrena's compiler requires this object-or-string property as text, and a real user would
	// not hand-write compact JSON.
	trustPolicy := s(`{
  "Version": "2012-10-17",
  "Statement": [ { "Effect": "Allow", "Principal": { "Service": "ecs-tasks.amazonaws.com" }, "Action": "sts:AssumeRole" } ]
}`)
	executionRole := create(t, prov, roleType, map[string]value.Value{
		"RoleName":                 s("infrena-live-ecsexec-" + run),
		"AssumeRolePolicyDocument": trustPolicy,
		"ManagedPolicyArns":        l(s("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy")),
		"Tags":                     tags,
	})

	clusterType := awstest.TypeFor(t, cat, "AWS::ECS::Cluster").Name
	cluster := create(t, prov, clusterType, map[string]value.Value{"region": s(region), "ClusterName": s(name),
		"ClusterSettings": l(m("name", s("containerInsights"), "value", s("disabled"))), "Tags": tags})

	container := m(
		"name", s("app"),
		"image", s("public.ecr.aws/docker/library/busybox:latest"),
		"essential", value.Bool(true, value.SourceExplicit),
		"command", l(s("sh"), s("-c"), s("sleep 3600")),
		// Options is opaque (see the comment above the function): its keys are AWS's own kebab-case, not
		// translated to snake_case the way LogDriver and the map's own key ("options") are.
		"log_configuration", m("log_driver", s("awslogs"), "options", m(
			"awslogs-group", logGroup.Attributes["LogGroupName"],
			"awslogs-region", s(region),
			"awslogs-stream-prefix", s("ecs"),
		)),
	)
	taskDefType := awstest.TypeFor(t, cat, "AWS::ECS::TaskDefinition").Name
	taskDef := create(t, prov, taskDefType, map[string]value.Value{
		"region":                  s(region),
		"Family":                  s(name),
		"Cpu":                     s("256"),
		"Memory":                  s("512"),
		"NetworkMode":             s("awsvpc"),
		"RequiresCompatibilities": l(s("FARGATE")),
		"ExecutionRoleArn":        executionRole.Attributes["Arn"],
		"ContainerDefinitions":    l(container),
		"Tags":                    tags,
	})

	serviceType := awstest.TypeFor(t, cat, "AWS::ECS::Service").Name
	service := create(t, prov, serviceType, map[string]value.Value{
		"region":         s(region),
		"Cluster":        cluster.Attributes["ClusterName"],
		"ServiceName":    s(name),
		"TaskDefinition": taskDef.Attributes["TaskDefinitionArn"],
		"LaunchType":     s("FARGATE"),
		// See the comment above the function: 0 is expected to be legal and to start nothing.
		"DesiredCount": n(0),
		"NetworkConfiguration": m("awsvpc_configuration", m(
			"subnets", l(subnetA.Attributes["SubnetId"], subnetB.Attributes["SubnetId"]),
			"security_groups", l(sg.Attributes["GroupId"]),
			"assign_public_ip", s("DISABLED"),
		)),
		"Tags": tags,
	})

	// Cheap, in-place update: a tag on the service, the same shape as every other test in this file — see
	// the comment above the function for why the task definition itself is never updated in place.
	service = update(t, prov, service, map[string]value.Value{"Tags": m(runTag, s(run), "Purpose", s("live-test"))})

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
	// All four new types list without a parent: none of aws.ecs.service, aws.ecs.taskdefinition,
	// aws.ecs.cluster or aws.loggroup carries list_needs_model in the catalog, unlike the ELB listener
	// and Route 53 record set elsewhere in this file, which both do and are skipped by discovery. So
	// every one of them is asserted here rather than only logged. Confirmed against real AWS on
	// 2026-09-17: all four were found, and the only skip reported was the listener/record set pair.
	for _, st := range []*resource.ResourceState{service, taskDef, cluster, logGroup, executionRole, sg, subnetB, subnetA, vpc} {
		if !seen[st.ProviderID] {
			t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", st.ProviderID)
		}
	}

	// Dependency order: the service depends on the cluster and the task definition (and would hold a
	// network interface in the subnets and the security group once it schedules a task, though
	// DesiredCount: 0 means it never does here); the task definition and the cluster are otherwise
	// independent of each other; the log group and the execution role are referenced from the task
	// definition's container and ExecutionRoleArn but not the reverse, so both outlive it here regardless.
	// deleteAndConfirm wraps the service (it can take a little while to report itself drained even with
	// nothing running) and the subnets/VPC (an ENI can hold a subnet); the rest have no lingering
	// dependency to wait out.
	deleteAndConfirm(t, prov, service)

	start := time.Now()
	if err := prov.Delete(ctx, taskDef); err != nil {
		t.Fatalf("delete %s: %v", taskDef.ProviderID, err)
	}
	t.Logf("deleted %s in %s", taskDef.ProviderID, time.Since(start).Round(time.Millisecond))
	// AWS::ECS::TaskDefinition's delete is DeregisterTaskDefinition, which marks the revision INACTIVE
	// rather than physically removing it — AWS keeps deregistered revisions visible to
	// DescribeTaskDefinition indefinitely. Cloud Control's GetResource for this type treats an INACTIVE
	// revision as not-found, the same as every other type's delete in this file — confirmed against real
	// AWS on 2026-09-17, where this read reported the resource gone while DescribeTaskDefinition still
	// returned revision :1 with status INACTIVE. That deregistered revision is permanent and free, and it
	// is the one thing this test cannot clean up after itself: Cloud Control's delete only deregisters,
	// and nothing in this suite calls AWS's own DeleteTaskDefinitions. See live/README.md.
	if got, err := prov.Read(ctx, taskDef); err != nil || got != nil {
		t.Errorf("read %s after delete = %v, %v; want gone", taskDef.ProviderID, got, err)
	}

	for _, st := range []*resource.ResourceState{cluster, logGroup, executionRole} {
		start := time.Now()
		if err := prov.Delete(ctx, st); err != nil {
			t.Fatalf("delete %s: %v", st.ProviderID, err)
		}
		t.Logf("deleted %s in %s", st.ProviderID, time.Since(start).Round(time.Millisecond))
		if got, err := prov.Read(ctx, st); err != nil || got != nil {
			t.Errorf("read %s after delete = %v, %v; want gone", st.ProviderID, got, err)
		}
	}
	for _, st := range []*resource.ResourceState{sg, subnetB, subnetA, vpc} {
		deleteAndConfirm(t, prov, st)
	}
}

// TestFunctionsAgainstRealAWS covers AWS Lambda: an execution role and a Lambda function whose deployment package
// is defined inline (Code.ZipFile), so nothing is uploaded to S3 and no container image is built or pushed. No
// VPC is created — the function sets no VpcConfig — so this is the smallest test in this file: two types, no
// subnets, no security group, and nothing for deleteAndConfirm to wait out.
//
// Cost: FREE, and deliberately so. Nothing here ever invokes the function — this test calls only the plugin's
// Create, Read, Update, Discover and Delete, never anything that runs code — and Lambda bills per invocation and
// per GB-second of execution, not for a function that merely exists. The inline deployment package is a few
// hundred bytes, far under Lambda's free code-storage allowance. Because the function is never invoked, Lambda
// never auto-creates its CloudWatch log group (normally named /aws/lambda/<FunctionName> and created lazily on
// first invocation) — none should exist after this test runs, and that absence is expected, not evidence of a
// leftover this test failed to clean up.
//
// The runtime: schemas/CloudformationSchema.zip's aws-lambda-function.json says inline code "works only for
// Node.js and Python functions", and types Runtime as a bare string with no enum in this schema bundle — so
// nothing about which runtimes AWS currently supports is checked mechanically here, unlike, say, an enum
// property would be. "python3.13" is used below on outside knowledge (Lambda added it in December 2024), not
// because the schema says so; if this test is run long after 2026, confirm python3.13 has not been deprecated
// (`aws lambda list-runtimes` or the console) before assuming a failure here is this test's fault. Handler is
// "index.handler" for the same reason as Node.js: CFN packages inline code into a file literally named "index"
// (with the extension matching the runtime, ".py" here), regardless of language.
//
// IAM PROPAGATION HAZARD — the most likely first-run failure, flagged rather than silently worked around:
// creating a Lambda function immediately after creating its execution role can fail with
// "InvalidParameterValueException: The role defined for the function cannot be assumed by Lambda", because the
// role has not yet propagated through IAM. Cloud Control's own Lambda handler already retries CreateFunction
// against this internally for a while before giving up, which is why this usually just works — but unlike
// aws.vpc/aws.subnet's read-side propagation retries (internal/ccprov/patience.go's patience, used by Read),
// there is no equivalent retry anywhere in this plugin, or in this file's create() helper, for a Create call AWS
// itself keeps rejecting synchronously. If the very first run fails on function creation with that message,
// that is the cause, not a bug: rerun it (the role will have propagated by then), or, if it recurs, add a short
// sleep between creating the role and creating the function.
//
// REWRITTEN VALUES — this test found a real bug on its first run, and Code is why it is worth keeping.
// Every leaf of Code (ImageUri, S3Bucket, S3Key, S3ObjectVersion, ZipFile, SourceKMSKeyArn, S3ObjectStorageMode)
// is writeOnlyProperties in aws-lambda-function.json, the same shape as RDS's MasterUserPassword (CLAUDE.md,
// internal/ccprov/values.go's stateFrom): AWS never returns a write-only value, so it is carried forward from
// what was last configured rather than compared against what Read gets back. That carry-forward only fires for
// an attribute the catalog's Type.WriteOnly lists, and the generator used to collect only pointers of the exact
// form "/properties/Name" — every one of Code's is one level deeper ("/properties/Code/ZipFile"), so none of
// them counted and Code was generated as an ordinary attribute.
//
// What the first live run showed (2026-09-17): Cloud Control's GetResource returns Code as an EMPTY OBJECT, so
// stateFrom recorded Code as an empty map with provider provenance against a configured zip_file, and
// converged() failed on Code specifically while create, the Description update, discovery and teardown all
// passed. That is a plan that never converges — every plan would want Code back, every apply would re-send it,
// and the next read would empty it again — which is the one failure class this provider cannot ship with.
//
// Fixed in the generator, not worked around here (commit 41cf7e8): a property whose every leaf is write-only is
// now write-only as a whole (cfn.Schema.WhollyNested), so Code is carried forward. A property with only SOME
// write-only leaves deliberately stays ordinary, because flagging it whole would hide real drift in the leaves
// AWS does return; those now warn instead of vanishing silently. Eight types gained an attribute and 87 partial
// cases became visible. If this assertion on Code ever fails again, suspect that rule before suspecting the test.
//
// The cheap, in-place update below is the function's Description, checked against the schema rather than
// assumed: createOnlyProperties for this type is only FunctionName, PackageType and TenancyConfig; Description
// is in neither readOnlyProperties nor writeOnlyProperties and is no part of Code, so changing it is a genuine
// UpdateFunctionConfiguration call, not a replacement and not a silent no-op. A tag would also have worked, the
// same way every other test in this file updates one, but Description exercises a plain scalar
// UpdateFunctionConfiguration path nothing else here does.
//
// Kept separate from the other tests so a run can target it alone with -run TestFunctionsAgainstRealAWS.
func TestFunctionsAgainstRealAWS(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	ctx := context.Background()
	run := strconv.FormatInt(time.Now().Unix(), 10)
	tags := m(runTag, s(run))
	name := "infrena-live-" + run

	roleType := awstest.TypeFor(t, cat, "AWS::IAM::Role").Name
	// JSON text, spaced unlike AWS's answer, the same reasoning as every other role this file creates: infrena's
	// compiler requires this object-or-string property as text, and a real user would not hand-write compact
	// JSON.
	trustPolicy := s(`{
  "Version": "2012-10-17",
  "Statement": [ { "Effect": "Allow", "Principal": { "Service": "lambda.amazonaws.com" }, "Action": "sts:AssumeRole" } ]
}`)
	role := create(t, prov, roleType, map[string]value.Value{
		"RoleName":                 s("infrena-live-lambdaexec-" + run),
		"AssumeRolePolicyDocument": trustPolicy,
		"Tags":                     tags,
	})
	// No managed policy is attached: this test never invokes the function, so nothing it does needs the
	// CloudWatch Logs permissions AWSLambdaBasicExecutionRole would grant. A trust policy naming
	// lambda.amazonaws.com is all CreateFunction checks.

	functionType := awstest.TypeFor(t, cat, "AWS::Lambda::Function").Name
	code := m("zip_file", s("def handler(event, context):\n    return {\"statusCode\": 200, \"body\": \"ok\"}\n"))
	fn := create(t, prov, functionType, map[string]value.Value{
		"region":       s(region),
		"FunctionName": s(name),
		"PackageType":  s("Zip"), // spelled out so create-only never differs from what AWS would infer; see the ALB test's IpAddressType for the same reasoning
		"Runtime":      s("python3.13"),
		"Handler":      s("index.handler"),
		"Code":         code,
		"Role":         role.Attributes["Arn"],
		"Description":  s("infrena live " + run),
		"Tags":         tags,
	})

	// Cheap, in-place update: see the comment above the function for why Description, not Tags, was chosen here.
	fn = update(t, prov, fn, map[string]value.Value{"Description": s("infrena live " + run + " updated")})

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
	// Both types list without a parent: neither aws.lambda.function nor aws.role carries list_needs_model in the
	// catalog (checked the same way as the ECS/Logs types above:
	// `zcat internal/catalog/catalog.json.gz | python3 -c '...'`), unlike aws.lambda.alias and
	// aws.lambda.permission, which both do and are skipped by discovery entirely — neither is used by this test.
	// So both are asserted here rather than only logged.
	for _, st := range []*resource.ResourceState{fn, role} {
		if !seen[st.ProviderID] {
			t.Errorf("discovery did not find %s (allow for propagation before calling it a bug)", st.ProviderID)
		}
	}

	// Dependency order: the function's Role names the role's Arn, so the function is deleted first.
	start := time.Now()
	if err := prov.Delete(ctx, fn); err != nil {
		t.Fatalf("delete %s: %v", fn.ProviderID, err)
	}
	t.Logf("deleted %s in %s", fn.ProviderID, time.Since(start).Round(time.Millisecond))
	if got, err := prov.Read(ctx, fn); err != nil || got != nil {
		t.Errorf("read %s after delete = %v, %v; want gone", fn.ProviderID, got, err)
	}

	start = time.Now()
	if err := prov.Delete(ctx, role); err != nil {
		t.Fatalf("delete %s: %v", role.ProviderID, err)
	}
	t.Logf("deleted %s in %s", role.ProviderID, time.Since(start).Round(time.Millisecond))
	if got, err := prov.Read(ctx, role); err != nil || got != nil {
		t.Errorf("read %s after delete = %v, %v; want gone", role.ProviderID, got, err)
	}
}

// TestSweepLeftovers deletes what a crashed run left: anything tagged by this suite more than an hour ago.
func TestSweepLeftovers(t *testing.T) {
	profile, region := guard(t)
	prov, cat := configure(t, profile, region)
	cutoff := time.Now().Add(-time.Hour).Unix()
	// Buckets and repositories sweep safely without checking for emptiness: this suite never puts objects or
	// images in them, so anything it tagged is always empty. Listeners, load balancers and target groups come
	// first, in that order: a load balancer's listeners go with it but a target group cannot be deleted while a
	// listener still forwards to it, and a load balancer holds network interfaces in the subnets and the security
	// group below. DB instances, subnet groups and parameter groups follow: an instance depends on the other two
	// plus subnets, and a subnet group depends on the subnets. The ECS service comes before the task definition
	// and the cluster it runs on (a service holds both), the task definition and the cluster are independent of
	// each other, and the log group they may still reference by name comes right after. The hosted zone comes
	// right before the VPC it may be associated with, for the same reason the security group and subnets do.
	//
	// All four ECS/Logs types added by TestContainerServicesAgainstRealAWS — the cluster, the task definition,
	// the service and the log group — declare tagging: {taggable: true} in their schemas, so unlike Route 53's
	// record set below, none of them needed leaving out of this sweep.
	//
	// AWS::Lambda::Function, added by TestFunctionsAgainstRealAWS, also declares tagging: {taggable: true}, so it
	// sweeps the same way. It comes right before the IAM role: the function's Role names the role's Arn, so the
	// function must go first.
	//
	// AWS::Route53::RecordSet is deliberately not in this list: its schema declares tagging: {taggable: false},
	// so unlike every other type here it cannot be tagged and therefore cannot be found by run tag at all. A
	// crash between TestDNSAgainstRealAWS creating a record set and its own teardown running leaves that record
	// behind untagged; DeleteHostedZone then refuses to remove the zone below until it is gone, which surfaces
	// here as a failed delete (t.Error, not t.Fatal, so it does not block the rest of this sweep) rather than a
	// silent leak. Narrow and cheap: a record set create is a few seconds of API calls, and nothing about an
	// orphaned record itself is billed.
	order := []string{
		awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::Listener").Name,
		awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::LoadBalancer").Name,
		awstest.TypeFor(t, cat, "AWS::ElasticLoadBalancingV2::TargetGroup").Name,
		awstest.TypeFor(t, cat, "AWS::RDS::DBInstance").Name,
		awstest.TypeFor(t, cat, "AWS::RDS::DBSubnetGroup").Name,
		awstest.TypeFor(t, cat, "AWS::RDS::DBParameterGroup").Name,
		awstest.TypeFor(t, cat, "AWS::ECS::Service").Name,
		awstest.TypeFor(t, cat, "AWS::ECS::TaskDefinition").Name,
		awstest.TypeFor(t, cat, "AWS::ECS::Cluster").Name,
		awstest.TypeFor(t, cat, "AWS::Logs::LogGroup").Name,
		awstest.TypeFor(t, cat, "AWS::ECR::Repository").Name,
		awstest.TypeFor(t, cat, "AWS::S3::Bucket").Name,
		awstest.TypeFor(t, cat, "AWS::Lambda::Function").Name,
		awstest.TypeFor(t, cat, "AWS::IAM::Role").Name,
		awstest.TypeFor(t, cat, "AWS::Route53::HostedZone").Name,
		"aws.securitygroup", "aws.subnet", "aws.vpc",
	}
	found, err := prov.Discover(context.Background(), provider.DiscoverRequest{Types: order})
	if err != nil {
		t.Fatal(err)
	}
	for _, typ := range order {
		// The tag property's own name, not always "Tags": AWS::Route53::HostedZone's is "HostedZoneTags"
		// (catalog.Type.TagsAsMap holds whatever the schema calls it; every other type swept here happens to
		// call it "Tags").
		tagAttr := "Tags"
		if ct, ok := cat.Lookup(typ); ok && ct.TagsAsMap != "" {
			tagAttr = ct.TagsAsMap
		}
		for _, r := range found {
			if r.Type != typ {
				continue
			}
			tags, _ := r.Attributes[tagAttr].Raw.(map[string]value.Value)
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
