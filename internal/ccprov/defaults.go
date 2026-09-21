package ccprov

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// ec2Defaults answers what AWS itself set up in one region.
//
// It exists because Cloud Control cannot answer it. AWS's published CloudFormation schemas carry no IsDefault on
// AWS::EC2::VPC and no DefaultForAz on AWS::EC2::Subnet, and nothing else in either read distinguishes a default
// (AWS::EC2::VPC's DefaultSecurityGroup and DefaultNetworkAcl are read-only properties of EVERY VPC, not just the
// default one). EC2 is the only API that states it, so discovery asks EC2 rather than inferring it from a CIDR or a
// name, which would be a guess.
type ec2Defaults interface {
	// DefaultVPC is the region's default VPC id, or "" when the account has none in that region.
	DefaultVPC(ctx context.Context, region string) (string, error)
	// DefaultSubnets are the ids of the subnets AWS marks default for their availability zone.
	DefaultSubnets(ctx context.Context, region string) (map[string]bool, error)
}

// ec2API is the real thing, one client per region, built on first use like the Cloud Control ones.
type ec2API struct{ clients *clients }

func (a *ec2API) DefaultVPC(ctx context.Context, region string) (string, error) {
	// One call: the filter leaves at most one VPC, so there is nothing to page through.
	out, err := a.clients.ec2(region).DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		Filters: []types.Filter{{Name: aws.String("is-default"), Values: []string{"true"}}},
	})
	if err != nil {
		return "", err
	}
	for _, v := range out.Vpcs {
		if aws.ToBool(v.IsDefault) {
			return aws.ToString(v.VpcId), nil
		}
	}
	return "", nil
}

func (a *ec2API) DefaultSubnets(ctx context.Context, region string) (map[string]bool, error) {
	pages := ec2.NewDescribeSubnetsPaginator(a.clients.ec2(region), &ec2.DescribeSubnetsInput{
		Filters: []types.Filter{{Name: aws.String("default-for-az"), Values: []string{"true"}}},
	})
	out := map[string]bool{}
	for pages.HasMorePages() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		page, err := pages.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		for _, s := range page.Subnets {
			if aws.ToBool(s.DefaultForAz) {
				out[aws.ToString(s.SubnetId)] = true
			}
		}
	}
	return out, nil
}

// regionFacts is the region context a check may consult about a resource. The first question costs one EC2 call and
// the rest of the run is answered from the cache, so a type whose evidence is in its own properties — a security
// group's GroupName, a role's Path — never causes one.
//
// Each answer carries whether it is KNOWN. EC2 that could not be asked is not EC2 saying no: an unknown answer leaves
// the flag unset rather than claiming the resource is a user's.
type regionFacts interface {
	defaultVPC(ctx context.Context) (id string, known bool)
	defaultSubnets(ctx context.Context) (ids map[string]bool, known bool)
}

// defaults caches one discovery run's answers, per region. A failure is cached too: EC2 is asked once per region per
// run whatever the answer, and a run over a thousand resources cannot turn into a thousand calls.
type defaults struct {
	api      ec2Defaults
	log      io.Writer
	instance string
	vpc      map[string]answer[string]
	subnets  map[string]answer[map[string]bool]
}

type answer[T any] struct {
	value T
	known bool
}

func newDefaults(p *Provider) *defaults {
	return &defaults{
		api: p.ec2, log: p.log, instance: p.instance,
		vpc: map[string]answer[string]{}, subnets: map[string]answer[map[string]bool]{},
	}
}

// in is the facts for one region.
func (d *defaults) in(region string) regionFacts { return &regionDefaults{defaults: d, region: region} }

type regionDefaults struct {
	*defaults
	region string
}

func (r *regionDefaults) defaultVPC(ctx context.Context) (string, bool) {
	if a, asked := r.vpc[r.region]; asked {
		return a.value, a.known
	}
	id, err := r.api.DefaultVPC(ctx, r.region)
	if err != nil {
		r.report(r.region, "which VPC is the default", err)
	}
	a := answer[string]{value: id, known: err == nil}
	r.vpc[r.region] = a
	return a.value, a.known
}

func (r *regionDefaults) defaultSubnets(ctx context.Context) (map[string]bool, bool) {
	if a, asked := r.subnets[r.region]; asked {
		return a.value, a.known
	}
	ids, err := r.api.DefaultSubnets(ctx, r.region)
	if err != nil {
		r.report(r.region, "which subnets are default for their availability zone", err)
	}
	a := answer[map[string]bool]{value: ids, known: err == nil}
	r.subnets[r.region] = a
	return a.value, a.known
}

// report says on stderr that a region's AWS-owned resources could not be identified. Never an error: the flag is
// advisory, and losing it must not lose the discovery it was attached to.
func (d *defaults) report(region, question string, err error) {
	fmt.Fprintf(d.log, "aws instance %q: asking EC2 %s in %s failed: %v\n"+
		"discovery cannot identify which resources AWS owns there, so none are marked as cloud-owned: "+
		"check ec2:DescribeVpcs and ec2:DescribeSubnets\n",
		d.instance, question, region, err)
}
