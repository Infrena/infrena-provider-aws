// Package awsprov is infrata's AWS provider plugin: instance configuration and credentials around the generic Cloud
// Control provider.
package awsprov

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/infrata/infrata-provider-aws/internal/catalog"
	"github.com/infrata/infrata-provider-aws/internal/ccprov"
	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/schema"
)

// PluginName is the binary's suffix, what `plugin:` names, and every type's prefix.
const PluginName = ccprov.PluginName

// Version is reported in the handshake. "0.0.0-dev" in every build a release did not stamp, so a broken -ldflags path
// cannot pass the release gate by coincidence. scripts/build-release stamps it.
var Version = "0.0.0-dev"

// Plugin is the AWS provider before configuration.
type Plugin struct{}

// NewPlugin returns the AWS plugin.
func NewPlugin() *Plugin { return &Plugin{} }

var _ provider.Plugin = (*Plugin)(nil)

// Name is the plugin's name.
func (pl *Plugin) Name() string { return PluginName }

// Version reports this build's version.
func (pl *Plugin) Version() string { return Version }

// Definitions are the embedded catalog's. They need no configuration and make no network call.
func (pl *Plugin) Definitions() []*schema.ResourceDefinition {
	cat, err := catalog.Embedded()
	if err != nil {
		// Only a broken build gets here; the catalog tests fail first.
		fmt.Fprintln(os.Stderr, "the aws plugin's embedded catalog is unreadable:", err)
		return nil
	}
	return cat.Definitions()
}

// New builds one configured instance. An error here is rendered against the `providers:` entry, so it says what is
// wrong and what to set. It makes no network call: every compiling command constructs instances.
func (pl *Plugin) New(cfg provider.Config) (provider.Provider, error) {
	ic, err := parseConfig(cfg.Values)
	if err != nil {
		return nil, err
	}
	cat, err := catalog.Embedded()
	if err != nil {
		return nil, err
	}
	var unknown []string
	for _, name := range ic.DiscoverTypes {
		if _, ok := cat.Lookup(name); !ok {
			unknown = append(unknown, strconv.Quote(name))
		}
	}
	if len(unknown) > 0 {
		return nil, fmt.Errorf("`discover_types` names %s, which the aws plugin does not serve; run `infrata explain <type>` to check a name",
			strings.Join(unknown, ", "))
	}
	awsCfg, err := loadAWSConfig(context.Background(), cfg.Instance, ic)
	if err != nil {
		return nil, err
	}
	return ccprov.New(cfg.Instance, cat, awsCfg, ccprov.Options{DiscoverRegions: ic.DiscoverRegions, DiscoverTypes: ic.DiscoverTypes}), nil
}
