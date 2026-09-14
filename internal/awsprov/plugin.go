// Package awsprov is infrata's AWS provider plugin.
package awsprov

import (
	"errors"

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

// New builds one configured instance. Instance configuration and credentials arrive in the next
// change; until then no instance can be configured, and saying so is better than a half-built one.
func (pl *Plugin) New(cfg provider.Config) (provider.Provider, error) {
	return nil, errors.New("the aws plugin cannot configure instances yet")
}
