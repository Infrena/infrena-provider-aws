# aws.monitoringsubscription

**CloudFormation type:** `AWS::CloudFront::MonitoringSubscription`

A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.

Global type (no region attribute)

**Import ID:** `global/DistributionId` (AWS::CloudFront::MonitoringSubscription)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DistributionId` | distribution_id | `string` | required, replaces on change | aws.cloudfront.distribution.Id | The ID of the distribution that you are enabling metrics for. |
| `MonitoringSubscription` | monitoring_subscription | `map` | required, replaces on change |  | A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution. |

Supports update: no

Discovery: not supported
