// Package awsprov is infrena's AWS provider plugin: instance configuration and credentials around the generic Cloud
// Control provider.
package awsprov

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
	"github.com/infrena/infrena-provider-aws/internal/ccprov"
	"github.com/infrena/infrena/pkg/provider"
	"github.com/infrena/infrena/pkg/schema"
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

// MaxConcurrency is how many operations infrena may run against this plugin at
// once (plugin protocol 5). It REPLACES the host's own per-provider default of
// 8 rather than merely capping it.
//
// The number is derived, and here is the derivation, because the SDK is right
// that a bare number reads as measured and the host will act on it.
//
// CLOUD CONTROL ALLOWS ABOUT 137 SIMULTANEOUS REQUESTS. One infrena operation
// is at most ONE simultaneous request: crud.go sends a single create, update or
// delete and then await() polls GetResourceRequestStatus one call at a time,
// sleeping with backoff in between, so an operation holds no more than one
// request open at any instant and spends most of its life holding none. N
// concurrent operations therefore peak at N simultaneous requests, and usually
// far fewer.
//
// So 24 peaks at about 17% of what the account allows. That margin is the
// point, and it is not timidity: the quota belongs to the ACCOUNT, not to this
// process. A colleague's apply, a CI run, the console and anything else in the
// account draw on the same budget, and infrena taking a fifth of it leaves room
// for five more of itself before anybody is throttled.
//
// It is three times the host's default, which is what makes it worth declaring
// at all — the SDK warns against restating 8, because a guess dressed as a
// claim is worth less than silence.
//
// IT IS A CEILING, NOT A TARGET. The global --parallelism (default 10) still
// bounds every provider together, so this changes nothing until somebody raises
// that; what it does is make raising it safe rather than a gamble.
//
// Changing it is this one constant. If throttling ever shows up in practice,
// lower it — ccprov classifies a throttle as retryable, so the symptom is a
// slower apply rather than a failed one, which is exactly the wrong signal to
// wait for.
func (pl *Plugin) MaxConcurrency() int { return 24 }

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
		return nil, fmt.Errorf("`discover_types` names %s, which the aws plugin does not serve; run `infrena explain <type>` to check a name",
			strings.Join(unknown, ", "))
	}
	awsCfg, err := loadAWSConfig(context.Background(), cfg.Instance, ic)
	if err != nil {
		return nil, err
	}
	return ccprov.New(cfg.Instance, cat, awsCfg, ccprov.Options{DiscoverRegions: ic.DiscoverRegions, DiscoverTypes: ic.DiscoverTypes}), nil
}
