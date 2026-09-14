package ccprov

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/infrena/infrena/pkg/provider"
	"github.com/infrena/infrena/pkg/resource"
	"github.com/infrena/infrena/pkg/value"
)

var ctx = context.Background()

func attr(t *testing.T, st *resource.ResourceState, name string) any {
	t.Helper()
	v, ok := st.Attributes[name]
	if !ok {
		t.Fatalf("state has no %s: %v", name, st.Attributes)
	}
	return v.Raw
}

func vpcAttrs(region string) map[string]value.Value {
	return map[string]value.Value{"region": sv(region), "CidrBlock": sv("10.0.0.0/16"), "Ipv4NetmaskLength": iv(16)}
}

func roleAttrs(name string) map[string]value.Value {
	return map[string]value.Value{
		"RoleName":                 sv(name),
		"AssumeRolePolicyDocument": value.Map(map[string]value.Value{"Version": sv("2012-10-17")}, value.SourceExplicit),
	}
}

func TestCreateReturnsWhatAWSReports(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil {
		t.Fatal(err)
	}
	id, ok := strings.CutPrefix(st.ProviderID, "us-east-1/")
	if !ok || !strings.HasPrefix(id, "vpc-") {
		t.Fatalf("provider ID = %q", st.ProviderID)
	}
	for name, want := range map[string]any{
		"region": "us-east-1", "CidrBlock": "10.0.0.0/16", "VpcId": id,
		"EnableDnsSupport":     true,                         // chosen by AWS: Optional+Computed
		"DefaultSecurityGroup": "defaultsecuritygroup-" + id, // read-only
		"Ipv4NetmaskLength":    int64(16),                    // write-only, carried from what was asked
	} {
		if got := attr(t, st, name); got != want {
			t.Errorf("%s = %#v, want %#v", name, got, want)
		}
	}
	stored, _ := fake.Resource("us-east-1", "AWS::EC2::VPC", id)
	if _, leaked := stored["region"]; leaked {
		t.Error("the plugin's region attribute was sent to AWS as a property")
	}
	if tokens := fake.Tokens("CreateResource"); len(tokens) != 1 || tokens[0] == "" {
		t.Errorf("client tokens = %v, want one", tokens)
	}
}

func TestRegionalTypesUseTheirRegionAndGlobalTypesUSEast1(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	if _, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("eu-west-1"))); err != nil {
		t.Fatal(err)
	}
	if n := len(fake.Resources("eu-west-1", "AWS::EC2::VPC")); n != 1 {
		t.Errorf("eu-west-1 holds %d VPCs, want 1", n)
	}
	role, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy")))
	if err != nil {
		t.Fatal(err)
	}
	if role.ProviderID != "global/deploy" {
		t.Errorf("role provider ID = %q", role.ProviderID)
	}
	if _, ok := fake.Resource("us-east-1", "AWS::IAM::Role", "deploy"); !ok {
		t.Error("a global type was not created through us-east-1")
	}
	if _, has := role.Attributes["region"]; has {
		t.Error("a global type reported a region attribute")
	}
}

// TestATypeWithItsOwnRegionPropertyUsesAwsRegion (J9).
func TestATypeWithItsOwnRegionPropertyUsesAwsRegion(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.test.regioned", map[string]value.Value{
		"aws_region": sv("eu-west-1"), "Region": sv("ap-south-1"), "Name": sv("x"),
	}))
	if err != nil {
		t.Fatal(err)
	}
	if attr(t, st, "aws_region") != "eu-west-1" || attr(t, st, "Region") != "ap-south-1" {
		t.Errorf("state = %v", st.Attributes)
	}
	if n := len(fake.Resources("eu-west-1", "AWS::Test::Regioned")); n != 1 {
		t.Errorf("created in the wrong region: eu-west-1 holds %d", n)
	}
}

func TestARegionalTypeWithoutARegionIsRefusedBeforeAnyCall(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	_, err := p.Create(ctx, desired("aws.vpc", map[string]value.Value{"CidrBlock": sv("10.0.0.0/16")}))
	if err == nil || !strings.Contains(err.Error(), "defaults: {region") {
		t.Fatalf("err = %v", err)
	}
	if fake.Calls("CreateResource") != 0 {
		t.Error("CreateResource was called")
	}
}

// TestAFailedCreateThatLeftAResourceIsRecordedNotOrphaned. The host drops an errored create's result, so returning the
// error would leave a real VPC nothing records.
func TestAFailedCreateThatLeftAResourceIsRecordedNotOrphaned(t *testing.T) {
	p, fake, log := fakeProvider(t)
	fake.FailNext("CREATE", "NotStabilized", "attachment did not settle", true)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil || st == nil {
		t.Fatalf("Create = %v, %v; want the state of what exists", st, err)
	}
	for _, want := range []string{`"test"`, "NotStabilized", "attachment did not settle", st.ProviderID} {
		if !strings.Contains(log.String(), want) {
			t.Errorf("stderr lacks %q:\n%s", want, log)
		}
	}
}

func TestAFailedCreateThatLeftNothingIsAnError(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.FailNext("CREATE", "ServiceLimitExceeded", "too many VPCs", false)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if st != nil || err == nil || !strings.Contains(err.Error(), "ServiceLimitExceeded") || !strings.Contains(err.Error(), "too many VPCs") {
		t.Fatalf("Create = %v, %v", st, err)
	}
	if p.ClassifyError(err) != provider.NotSafeToRetry {
		t.Errorf("classify = %v", p.ClassifyError(err))
	}
}

func TestANameThatExistsSuggestsImporting(t *testing.T) {
	p, _, _ := fakeProvider(t)
	if _, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy"))); err != nil {
		t.Fatal(err)
	}
	if _, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy"))); err == nil || !strings.Contains(err.Error(), "infrena import") {
		t.Fatalf("second create: err = %v", err)
	}
}

func TestCancellationAfterTheCreateIsSentStillRecordsTheResource(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.PollsToComplete = 2
	cctx, cancel := context.WithCancel(ctx)
	defer cancel()
	p.pacing.sleep = func(context.Context, time.Duration) error { cancel(); return nil }
	st, err := p.Create(cctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil || st == nil {
		t.Fatalf("Create after cancellation = %v, %v", st, err)
	}
}

// TestCreateNeverReportsFailureOnceAWSCreated. AWS said SUCCESS but the resource is not readable yet: the configured
// values are recorded, and the next refresh reads the rest.
func TestCreateNeverReportsFailureOnceAWSCreated(t *testing.T) {
	p, fake, log := fakeProvider(t)
	fake.HideFromGet("deploy", 100)
	st, err := p.Create(ctx, desired("aws.role", roleAttrs("deploy")))
	if err != nil || st == nil || st.ProviderID != "global/deploy" || attr(t, st, "RoleName") != "deploy" {
		t.Fatalf("Create = %+v, %v", st, err)
	}
	if !strings.Contains(log.String(), "could not be read back") {
		t.Errorf("stderr does not say so:\n%s", log)
	}
}

func TestReadWaitsOutPropagation(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1", "CidrBlock": "10.0.0.0/16"})
	fake.HideFromGet("vpc-1", 2)
	st, err := p.Read(ctx, &resource.ResourceState{Type: "aws.vpc", ProviderID: "us-east-1/vpc-1"})
	if err != nil || st == nil {
		t.Fatalf("Read = %v, %v", st, err)
	}
	if n := fake.Calls("GetResource"); n != 3 {
		t.Errorf("GetResource calls = %d, want 3", n)
	}
}

func TestReadReportsGoneAfterPatience(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Read(ctx, &resource.ResourceState{Type: "aws.vpc", ProviderID: "us-east-1/vpc-missing"})
	if err != nil || st != nil {
		t.Fatalf("Read = %v, %v; want (nil, nil)", st, err)
	}
	if n := fake.Calls("GetResource"); n != notFoundPatience.attempts {
		t.Errorf("GetResource calls = %d, want %d", n, notFoundPatience.attempts)
	}
}

func TestReadCarriesWriteOnlyValuesForward(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::RDS::DBInstance", "db-1", map[string]any{"DBInstanceIdentifier": "db-1"})
	st, err := p.Read(ctx, &resource.ResourceState{Type: "aws.dbinstance", ProviderID: "us-east-1/db-1",
		Attributes: map[string]value.Value{"MasterUserPassword": sv("hunter2"), "DBInstanceClass": sv("db.t3.micro")}})
	if err != nil {
		t.Fatal(err)
	}
	if attr(t, st, "MasterUserPassword") != "hunter2" {
		t.Error("a write-only value AWS never returns was dropped")
	}
	if _, has := st.Attributes["DBInstanceClass"]; has {
		t.Error("a readable property AWS did not return was carried forward: drift would be hidden")
	}
}

// TestAnEmptyCollectionAWSOmitsStaysEmpty. `policies: []` in configuration, and AWS leaving the property out, would
// otherwise plan a change forever.
func TestAnEmptyCollectionAWSOmitsStaysEmpty(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::IAM::Role", "deploy", map[string]any{"RoleName": "deploy"})
	empty := value.List(nil, value.SourceExplicit)
	st, err := p.Read(ctx, &resource.ResourceState{Type: "aws.role", ProviderID: "global/deploy",
		Attributes: map[string]value.Value{"Policies": empty}})
	if err != nil {
		t.Fatal(err)
	}
	if got, ok := st.Attributes["Policies"]; !ok || !got.Equal(empty) {
		t.Errorf("Policies = %v, want []", got)
	}
	notEmpty := value.List([]value.Value{sv("x")}, value.SourceExplicit)
	st, _ = p.Read(ctx, &resource.ResourceState{Type: "aws.role", ProviderID: "global/deploy",
		Attributes: map[string]value.Value{"Policies": notEmpty}})
	if _, has := st.Attributes["Policies"]; has {
		t.Error("a non-empty list AWS no longer has was carried forward")
	}
}

func TestMalformedIDsNeverReachAWS(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	if _, err := p.Read(ctx, &resource.ResourceState{Type: "aws.vpc", ProviderID: "vpc-1"}); err == nil {
		t.Error("an ID without a region was read")
	}
	if err := p.Delete(ctx, &resource.ResourceState{Type: "aws.role", ProviderID: "us-east-1/deploy"}); err == nil {
		t.Error("a global type was deleted by a regional ID")
	}
	if n := fake.Calls("GetResource") + fake.Calls("DeleteResource"); n != 0 {
		t.Errorf("%d calls were made", n)
	}
}

func TestDeleteRemovesAndTreatsGoneAsDone(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Delete(ctx, st); err != nil {
		t.Fatal(err)
	}
	if n := len(fake.Resources("us-east-1", "AWS::EC2::VPC")); n != 0 {
		t.Fatalf("%d VPCs remain", n)
	}
	if err := p.Delete(ctx, st); err != nil {
		t.Errorf("deleting what is already gone = %v, want success", err)
	}
}

func TestADeleteThatFailsSaysWhy(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	st, err := p.Create(ctx, desired("aws.vpc", vpcAttrs("us-east-1")))
	if err != nil {
		t.Fatal(err)
	}
	fake.FailNext("DELETE", "ResourceConflict", "the VPC has dependencies", false)
	err = p.Delete(ctx, st)
	if err == nil || !strings.Contains(err.Error(), "the VPC has dependencies") {
		t.Fatalf("err = %v", err)
	}
	if p.ClassifyError(err) != provider.ConditionallyRetryable {
		t.Errorf("classify = %v", p.ClassifyError(err))
	}
}

func TestImportReadsOnceAndSaysWhatIsMissing(t *testing.T) {
	p, fake, _ := fakeProvider(t)
	fake.Put("us-east-1", "AWS::EC2::VPC", "vpc-1", map[string]any{"VpcId": "vpc-1", "CidrBlock": "172.16.0.0/16"})
	st, err := p.Import(ctx, "aws.vpc", "us-east-1/vpc-1")
	if err != nil || attr(t, st, "CidrBlock") != "172.16.0.0/16" || st.ProviderID != "us-east-1/vpc-1" {
		t.Fatalf("Import = %+v, %v", st, err)
	}
	before := fake.Calls("GetResource")
	if _, err := p.Import(ctx, "aws.vpc", "us-east-1/vpc-missing"); err == nil || !strings.Contains(err.Error(), "us-east-1/vpc-missing") {
		t.Errorf("importing what does not exist: err = %v", err)
	}
	if n := fake.Calls("GetResource") - before; n != 1 {
		t.Errorf("import made %d reads, want 1", n)
	}
}

func TestAnUnknownTypeIsRefused(t *testing.T) {
	p, _, _ := fakeProvider(t)
	if _, err := p.Create(ctx, desired("aws.nope", nil)); err == nil || !strings.Contains(err.Error(), "aws.nope") {
		t.Fatalf("err = %v", err)
	}
}

// TestNothingIsWrittenToStdout. Stdout is the plugin protocol's pipe.
func TestNothingIsWrittenToStdout(t *testing.T) {
	if p := New("x", testCatalog(), testConfig("http://127.0.0.1:1", 1), Options{}); p.log != os.Stderr {
		t.Error("the provider's log is not stderr")
	}
}
