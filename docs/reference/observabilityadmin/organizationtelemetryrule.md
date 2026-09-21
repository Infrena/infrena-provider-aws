# aws.organizationtelemetryrule

**CloudFormation type:** `AWS::ObservabilityAdmin::OrganizationTelemetryRule`

The AWS::ObservabilityAdmin::OrganizationTelemetryRule resource defines a CloudWatch Observability Admin Organization Telemetry Rule.

Region attribute: `region`

**Import ID:** `<region>/RuleArn` (AWS::ObservabilityAdmin::OrganizationTelemetryRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RegionStatuses` | region_statuses | `list` | computed |  | Per-region replication status of the rule |
| `Rule` |  | `map` | required |  | The telemetry rule |
| `RuleArn` | rule_arn | `string` | computed |  | The arn of the organization telemetry rule |
| `RuleName` | rule_name | `string` | required, replaces on change |  | The name of the organization telemetry rule |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource |

Supports update: yes

Discovery: supported
