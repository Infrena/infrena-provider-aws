package gen

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// Overlay is the hand-maintained part of generation: what CloudFormation schemas cannot say (spec §3.2).
type Overlay struct {
	Global          []string                        `yaml:"global"`
	DiscoverDefault []string                        `yaml:"discover_default"`
	Aliases         map[string]map[string][]string  `yaml:"aliases"`
	Sensitive       map[string][]string             `yaml:"sensitive"`
	Requirements    map[string][]OverlayRequirement `yaml:"requirements"`
	References      OverlayReferences               `yaml:"references"`
}

// OverlayReferences is the human review of derived reference edges (resource references design §4.2), written with
// CloudFormation type names.
type OverlayReferences struct {
	// CrossServiceTargets lists the target types an exact match from ANOTHER service may reach as tier 1. Any other
	// cross-service exact match is tier 2, because generic names (ResourceId, NotificationArns) match exactly and
	// wrongly. Same-service exact matches are tier 1 without being listed.
	CrossServiceTargets []string `yaml:"cross_service_targets"`
	// ApproveTargets lists target types whose tier-2 edges are approved, all of them at once.
	ApproveTargets []string `yaml:"approve_targets"`
	// RejectTargets lists target types every edge to which from ANOTHER service is refused, whatever its tier and even
	// if approved. Edges from the target's own service keep their normal status: the fabrications this exists for come
	// from generic names in other services (RDS's SourceDbiResourceId), while ApiGateway::Method.ResourceId is real.
	RejectTargets []string `yaml:"reject_targets"`
	// Reject lists Type.Property edges refused whatever their tier.
	Reject []string `yaml:"reject"`
}

// OverlayRequirement is a pre-flight hint, written with CloudFormation type names.
type OverlayRequirement struct {
	Name        string   `yaml:"name"`
	Types       []string `yaml:"types"`
	Description string   `yaml:"description"`
}

// LoadOverlay reads the overlay, refusing unknown keys: a misspelt section would silently apply nothing.
func LoadOverlay(path string) (*Overlay, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(bytes.NewReader(raw))
	dec.KnownFields(true)
	var o Overlay
	if err := dec.Decode(&o); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	return &o, nil
}

// IsGlobal matches a CloudFormation type against the global patterns: an exact name, or a prefix ending in "*".
func (o *Overlay) IsGlobal(cfn string) bool {
	for _, p := range o.Global {
		if prefix, ok := strings.CutSuffix(p, "*"); ok && strings.HasPrefix(cfn, prefix) {
			return true
		}
		if p == cfn {
			return true
		}
	}
	return false
}
