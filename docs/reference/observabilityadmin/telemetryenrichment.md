# aws.telemetryenrichment

**CloudFormation type:** `AWS::ObservabilityAdmin::TelemetryEnrichment`

AWS::ObservabilityAdmin::TelemetryEnrichment cloudformation resource enables the resource tags for telemetry feature in CloudWatch to enrich infrastructure metrics with AWS resource tags. For more details: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/resource-tags-for-telemetry.html

Region attribute: `region`

**Import ID:** `<region>/Scope` (AWS::ObservabilityAdmin::TelemetryEnrichment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Scope` |  | `string` | required, replaces on change |  | Scope of the Telemetry Enrichment |
| `Status` |  | `string` | computed |  | Current status of the resource tags for telemetry feature (Running, Stopped, or Impaired). |

Supports update: yes

Discovery: supported
