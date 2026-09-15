// Command gen-docs generates reference documentation from the embedded catalog.
// It writes one Markdown file per type to docs/reference/<service>/<name>.md,
// and an index to docs/reference/README.md. Byte-stable: sorted, no timestamps.
//
//	go run ./cmd/gen-docs
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/infrena/infrena-provider-aws/internal/catalog"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "gen-docs:", err)
		os.Exit(1)
	}
}

func run() error {
	cat, err := catalog.Embedded()
	if err != nil {
		return err
	}

	// Write reference pages for each type
	for _, t := range cat.Types {
		if err := writeTypePage(t); err != nil {
			return fmt.Errorf("write %s: %w", t.Name, err)
		}
	}

	// Write index
	if err := writeIndex(cat.Types); err != nil {
		return fmt.Errorf("write index: %w", err)
	}

	return nil
}

func writeTypePage(t *catalog.Type) error {
	service := serviceFromCFN(t.CFN)
	name := typeNameToDashCase(t.Name)

	dir := filepath.Join("docs", "reference", service)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	path := filepath.Join(dir, name+".md")
	content := formatTypePage(t)

	return os.WriteFile(path, []byte(content), 0o644)
}

func formatTypePage(t *catalog.Type) string {
	var sb strings.Builder

	// Title and type info
	sb.WriteString("# " + t.Name + "\n\n")
	sb.WriteString("**CloudFormation type:** `" + t.CFN + "`\n\n")

	if t.Description != "" {
		sb.WriteString(t.Description + "\n\n")
	}

	// Region info
	if t.Global() {
		sb.WriteString("Global type (no region attribute)\n\n")
	} else {
		sb.WriteString("Region attribute: `" + t.RegionAttr + "`\n\n")
	}

	// Import ID
	scope := "<region>"
	if t.Global() {
		scope = "global"
	}
	identifier := strings.Join(t.Identifier, "|")
	sb.WriteString("**Import ID:** `" + scope + "/" + identifier + "` (" + t.CFN + ")\n\n")

	// Attributes table
	sb.WriteString("## Attributes\n\n")
	sb.WriteString("| Attribute | Also written as | Kind | Flags | References | Description |\n")
	sb.WriteString("|-----------|-----------------|------|-------|------------|-------------|\n")

	// Sort attributes by name for stable output
	attrs := make([]*catalog.Attribute, len(t.Attributes))
	copy(attrs, t.Attributes)
	sort.Slice(attrs, func(i, j int) bool { return attrs[i].Name < attrs[j].Name })

	for _, a := range attrs {
		flags := formatFlags(a, t)
		refs := formatReferences(a)
		desc := escapeTableCell(a.Description)
		aliases := escapeTableCell(strings.Join(a.Aliases, ", "))
		sb.WriteString("| `" + a.Name + "` | " + aliases + " | `" + a.Kind + "` | " + flags + " | " + refs + " | " + desc + " |\n")
	}

	sb.WriteString("\n")

	// Update and discovery support
	if t.HasUpdate {
		sb.WriteString("Supports update: yes\n\n")
	} else {
		sb.WriteString("Supports update: no\n\n")
	}

	if t.HasList {
		if t.ListNeedsModel {
			sb.WriteString("Discovery: supported (parent resource required)\n")
		} else {
			sb.WriteString("Discovery: supported\n")
		}
	} else {
		sb.WriteString("Discovery: not supported\n")
	}

	return sb.String()
}

func formatFlags(a *catalog.Attribute, t *catalog.Type) string {
	var flags []string

	if a.Required {
		flags = append(flags, "required")
	}
	if a.Optional {
		flags = append(flags, "optional")
	}
	if a.Computed {
		flags = append(flags, "computed")
	}
	if a.Optional && a.Computed {
		flags = append(flags, "provider-chosen")
	}
	if a.ForceNew {
		flags = append(flags, "replaces on change")
	}
	if a.Sensitive {
		flags = append(flags, "sensitive")
	}

	// Check if in WriteOnly
	for _, wo := range t.WriteOnly {
		if wo == a.Name {
			flags = append(flags, "write-only")
			break
		}
	}

	// Check if TagsAsMap
	if t.TagsAsMap != "" && t.TagsAsMap == a.Name {
		flags = append(flags, "tags map")
	}

	if len(flags) == 0 {
		return ""
	}
	return strings.Join(flags, ", ")
}

func formatReferences(a *catalog.Attribute) string {
	if a.References == nil {
		return ""
	}
	return a.References.Type + "." + a.References.Attribute
}

func escapeTableCell(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "|", "\\|")
	if s == "" {
		return ""
	}
	return s
}

func serviceFromCFN(cfnType string) string {
	// AWS::EC2::VPC -> ec2
	parts := strings.Split(cfnType, "::")
	if len(parts) < 2 {
		return "unknown"
	}
	return strings.ToLower(parts[1])
}

func typeNameToDashCase(name string) string {
	// aws.vpc -> vpc
	// aws.my.type -> my-type
	parts := strings.Split(name, ".")
	if len(parts) < 2 {
		return name
	}
	// Skip "aws" prefix
	rest := parts[1:]
	return strings.Join(rest, "-")
}

func writeIndex(types []*catalog.Type) error {
	dir := filepath.Join("docs", "reference")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	// Group types by service
	byService := make(map[string][]*catalog.Type)
	for _, t := range types {
		service := serviceFromCFN(t.CFN)
		byService[service] = append(byService[service], t)
	}

	// Sort services
	var services []string
	for service := range byService {
		services = append(services, service)
	}
	sort.Strings(services)

	var sb strings.Builder
	sb.WriteString("# AWS Provider Reference\n\n")
	sb.WriteString("Reference documentation for all AWS resource types provided by infrena-provider-aws.\n\n")

	for _, service := range services {
		types := byService[service]
		// Sort types by name
		sort.Slice(types, func(i, j int) bool { return types[i].Name < types[j].Name })

		sb.WriteString("## " + capitalizeService(service) + " (" + fmt.Sprintf("%d", len(types)) + ")\n\n")

		for _, t := range types {
			name := typeNameToDashCase(t.Name)
			path := service + "/" + name + ".md"
			sb.WriteString("- [" + t.Name + "](" + path + ")\n")
		}
		sb.WriteString("\n")
	}

	path := filepath.Join(dir, "README.md")
	return os.WriteFile(path, []byte(sb.String()), 0o644)
}

func capitalizeService(service string) string {
	// ec2 -> EC2, iam -> IAM, etc.
	upper := strings.ToUpper(service)
	switch upper {
	case "EC2":
		return "EC2"
	case "IAM":
		return "IAM"
	case "RDS":
		return "RDS"
	case "S3":
		return "S3"
	case "VPC":
		return "VPC"
	case "ECS":
		return "ECS"
	case "EKS":
		return "EKS"
	case "KMS":
		return "KMS"
	case "SNS":
		return "SNS"
	case "SQS":
		return "SQS"
	case "LAMBDA":
		return "Lambda"
	case "DYNAMODB":
		return "DynamoDB"
	case "ELASTICACHE":
		return "ElastiCache"
	case "APIGATEWAY":
		return "API Gateway"
	case "CLOUDFRONT":
		return "CloudFront"
	case "ROUTE53":
		return "Route 53"
	case "AUTOSCALING":
		return "Auto Scaling"
	case "ELASTICLOADBALANCING":
		return "Elastic Load Balancing"
	default:
		// Capitalize first letter
		if len(upper) == 0 {
			return upper
		}
		return strings.ToUpper(string(upper[0])) + strings.ToLower(upper[1:])
	}
}
