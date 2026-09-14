package awsprov

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// stsFallbackRegion is used only when nothing configured a region: resources carry their own, so a
// user may have no default region, and an STS client without one cannot resolve an endpoint.
const stsFallbackRegion = "us-east-1"

// loadAWSConfig loads credentials the way the AWS CLI does, then applies this instance's overrides.
//
// It resolves nothing over the network: credentials are retrieved on first use. Every compiling
// command constructs instances, `validate` included, and an SSO or assume-role call there would put
// a network round trip in front of a syntax check.
func loadAWSConfig(ctx context.Context, instance string, ic instanceConfig) (aws.Config, error) {
	var opts []func(*config.LoadOptions) error
	if ic.Profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(ic.Profile))
	}
	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		var missing config.SharedConfigProfileNotExistError
		if errors.As(err, &missing) {
			return aws.Config{}, fmt.Errorf("provider instance %q names profile %q, which is not in your AWS config files: "+
				"run \"aws configure list-profiles\" to see what exists, or correct `profile`", instance, ic.Profile)
		}
		return aws.Config{}, fmt.Errorf("provider instance %q: loading AWS configuration: %w", instance, err)
	}
	if ic.AssumeRoleARN != "" {
		client := sts.NewFromConfig(cfg, func(o *sts.Options) {
			if o.Region == "" {
				o.Region = stsFallbackRegion
			}
		})
		cfg.Credentials = aws.NewCredentialsCache(stscreds.NewAssumeRoleProvider(client, ic.AssumeRoleARN))
	}
	return cfg, nil
}
