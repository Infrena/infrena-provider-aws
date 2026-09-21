# aws.otelenrichment

**CloudFormation type:** `AWS::CloudWatch::OTelEnrichment`

AWS::CloudWatch::OTelEnrichment enables OTel metric enrichment in CloudWatch, allowing CloudWatch vended metrics to be available for PromQL querying enriched with AWS resource tags and metadata.

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::CloudWatch::OTelEnrichment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | The AWS account ID. This is the primary identifier for this singleton resource. |
| `Status` |  | `string` | computed |  | Current status of OTel enrichment (RUNNING or STOPPED). |

Supports update: yes

Discovery: supported
