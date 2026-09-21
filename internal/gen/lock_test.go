package gen

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestAFreshLockUsesShortNamesOnlyWhenUnique. J2: the resource segment lowercased when no other supported type
// shares it, otherwise service.segment.
func TestAFreshLockUsesShortNamesOnlyWhenUnique(t *testing.T) {
	l := &Lock{}
	got, err := l.Assign([]string{"AWS::EC2::VPC", "AWS::EC2::Instance", "AWS::Lightsail::Instance", "AWS::S3::Bucket"})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{
		"AWS::EC2::VPC":            "aws.vpc",
		"AWS::EC2::Instance":       "aws.ec2.instance",
		"AWS::Lightsail::Instance": "aws.lightsail.instance",
		"AWS::S3::Bucket":          "aws.bucket",
	}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s -> %q, want %q", k, got[k], v)
		}
	}
	if len(l.Names) != 4 {
		t.Errorf("lock holds %d names, want the 4 assigned", len(l.Names))
	}
}

// TestALockedShortNameNeverMoves. J3: a new type whose segment clashes with a locked short name gets the qualified
// form, and the locked type keeps its name. The fixture's lock would otherwise be rewritten by the uniqueness rule.
func TestALockedShortNameNeverMoves(t *testing.T) {
	l := &Lock{Names: map[string]string{"AWS::EC2::VPC": "aws.vpc"}}
	got, err := l.Assign([]string{"AWS::EC2::VPC", "AWS::NetworkFirewall::VPC"})
	if err != nil {
		t.Fatal(err)
	}
	if got["AWS::EC2::VPC"] != "aws.vpc" || got["AWS::NetworkFirewall::VPC"] != "aws.networkfirewall.vpc" {
		t.Fatalf("assigned %v", got)
	}
}

// TestARemovedTypeKeepsItsNameReserved. A tombstone means a name never comes back meaning something else.
func TestARemovedTypeKeepsItsNameReserved(t *testing.T) {
	l := &Lock{Names: map[string]string{"AWS::Old::Widget": "aws.widget"}}
	got, err := l.Assign([]string{"AWS::New::Widget"})
	if err != nil {
		t.Fatal(err)
	}
	if got["AWS::New::Widget"] != "aws.new.widget" {
		t.Errorf("new widget = %q, want the qualified form: aws.widget belongs to a removed type", got["AWS::New::Widget"])
	}
	if _, kept := l.Names["AWS::Old::Widget"]; !kept {
		t.Error("the tombstone was dropped")
	}
	if _, returned := got["AWS::Old::Widget"]; returned {
		t.Error("a removed type was returned as current")
	}
}

func TestTwoTypesCannotShareAName(t *testing.T) {
	l := &Lock{Names: map[string]string{"AWS::A::Thing": "aws.b.thing"}}
	if _, err := l.Assign([]string{"AWS::A::Thing", "AWS::B::Thing"}); err == nil || !strings.Contains(err.Error(), "aws.b.thing") {
		t.Fatalf("err = %v", err)
	}
}

func TestTheLockRoundTripsSorted(t *testing.T) {
	path := filepath.Join(t.TempDir(), "names.lock.json")
	l, err := LoadLock(path)
	if err != nil || len(l.Names) != 0 {
		t.Fatalf("missing lock = %+v, %v", l, err)
	}
	if _, err := l.Assign([]string{"AWS::S3::Bucket", "AWS::EC2::VPC"}); err != nil {
		t.Fatal(err)
	}
	if err := l.Save(path); err != nil {
		t.Fatal(err)
	}
	raw, _ := os.ReadFile(path)
	if strings.Index(string(raw), "AWS::EC2::VPC") > strings.Index(string(raw), "AWS::S3::Bucket") || !strings.HasSuffix(string(raw), "\n") {
		t.Errorf("lock not sorted or no trailing newline:\n%s", raw)
	}
	again, err := LoadLock(path)
	if err != nil || again.Names["AWS::S3::Bucket"] != "aws.bucket" {
		t.Fatalf("reloaded = %+v, %v", again, err)
	}
}
