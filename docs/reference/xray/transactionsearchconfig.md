# aws.transactionsearchconfig

**CloudFormation type:** `AWS::XRay::TransactionSearchConfig`

This schema provides construct and validation rules for AWS-XRay TransactionSearchConfig resource parameters.

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::XRay::TransactionSearchConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | User account id, used as the primary identifier for the resource |
| `IndexingPercentage` | indexing_percentage | `float` | optional, computed, provider-chosen |  | Determines the percentage of traces indexed from CloudWatch Logs to X-Ray |

Supports update: yes

Discovery: supported
