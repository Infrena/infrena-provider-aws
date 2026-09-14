package awsprov

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/infrata/infrata/pkg/value"
)

const (
	keyAssumeRoleARN   = "assume_role_arn"
	keyDiscoverRegions = "discover_regions"
	keyProfile         = "profile"
)

// instanceConfig is one `providers:` entry's own configuration: credentials, and the regions
// discovery scans. Never a region for CRUD, which comes from each resource.
type instanceConfig struct {
	Profile         string
	AssumeRoleARN   string
	DiscoverRegions []string
}

// parseConfig fails closed: a key this plugin does not read is refused, naming what it accepts. A
// misspelled key silently ignored falls back to the default credential chain — another account.
func parseConfig(values map[string]value.Value) (instanceConfig, error) {
	var ic instanceConfig
	var unknown []string
	for k := range values {
		switch k {
		case keyAssumeRoleARN, keyDiscoverRegions, keyProfile:
		default:
			unknown = append(unknown, strconv.Quote(k))
		}
	}
	if len(unknown) > 0 {
		sort.Strings(unknown)
		return ic, fmt.Errorf("unknown configuration %s; the aws provider accepts only %s, %s and %s "+
			"(a resource's region belongs under defaults: {region: …}, not here)",
			strings.Join(unknown, ", "), keyAssumeRoleARN, keyDiscoverRegions, keyProfile)
	}

	var err error
	if ic.Profile, err = optionalString(values, keyProfile); err != nil {
		return ic, err
	}
	if ic.AssumeRoleARN, err = optionalString(values, keyAssumeRoleARN); err != nil {
		return ic, err
	}
	if ic.AssumeRoleARN != "" && !strings.HasPrefix(ic.AssumeRoleARN, "arn:") {
		return ic, fmt.Errorf("`%s` must be a role ARN such as arn:aws:iam::123456789012:role/deploy, got %q",
			keyAssumeRoleARN, ic.AssumeRoleARN)
	}
	if v, ok := values[keyDiscoverRegions]; ok {
		items, isList := v.Raw.([]value.Value)
		if v.Kind != value.KindList || !isList {
			return ic, fmt.Errorf("`%s` must be a list of regions, e.g. [us-east-1, eu-west-1], got %s", keyDiscoverRegions, v.Kind)
		}
		seen := map[string]bool{}
		for i, item := range items {
			region, isString := item.AsString()
			if !isString || region == "" {
				return ic, fmt.Errorf("`%s` item %d must be a region name such as us-east-1", keyDiscoverRegions, i+1)
			}
			if !seen[region] {
				seen[region] = true
				ic.DiscoverRegions = append(ic.DiscoverRegions, region)
			}
		}
	}
	return ic, nil
}

func optionalString(values map[string]value.Value, key string) (string, error) {
	v, ok := values[key]
	if !ok {
		return "", nil
	}
	text, isString := v.AsString()
	if !isString {
		return "", fmt.Errorf("`%s` must be a string, got %s", key, v.Kind)
	}
	if text == "" {
		return "", fmt.Errorf("`%s` is empty: give a value, or omit the key to use the default credential chain", key)
	}
	return text, nil
}
