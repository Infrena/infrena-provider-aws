// Package awsprov is infrata's AWS provider plugin.
package awsprov

import (
	"context"

	"github.com/infrata/infrata/pkg/provider"
	"github.com/infrata/infrata/pkg/schema"
)

// PluginName is the binary's suffix, what `plugin:` names, and every type's prefix.
const PluginName = "aws"

// Version is reported in the handshake. "0.0.0-dev" in every build a release did not stamp, so a
// broken -ldflags path cannot pass the release gate by coincidence. scripts/build-release stamps it.
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

// Definitions need no configuration and make no network call: every command loads them.
func (pl *Plugin) Definitions() []*schema.ResourceDefinition { return definitions() }

// New builds one configured instance. An error here is rendered against the `providers:` entry, so
// it must say what is wrong and what to set.
func (pl *Plugin) New(cfg provider.Config) (provider.Provider, error) {
	ic, err := parseConfig(cfg.Values)
	if err != nil {
		return nil, err
	}
	awsCfg, err := loadAWSConfig(context.Background(), cfg.Instance, ic)
	if err != nil {
		return nil, err
	}
	return newProvider(cfg.Instance, ic, awsCfg), nil
}
