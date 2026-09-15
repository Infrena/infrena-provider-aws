# aws.macie.session

**CloudFormation type:** `AWS::Macie::Session`

The AWS::Macie::Session resource specifies a new Amazon Macie session. A session is an object that represents the Amazon Macie service. A session is required for Amazon Macie to become operational.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId` (AWS::Macie::Session)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutomatedDiscoveryStatus` | automated_discovery_status | `string` | computed |  | The status of automated sensitive data discovery for the Macie session. |
| `AwsAccountId` | aws_account_id | `string` | computed |  | AWS account ID of customer |
| `FindingPublishingFrequency` | finding_publishing_frequency | `string` | optional, computed, provider-chosen |  | A enumeration value that specifies how frequently finding updates are published. |
| `ServiceRole` | service_role | `string` | computed |  | Service role used by Macie |
| `Status` |  | `string` | optional, computed, provider-chosen |  | A enumeration value that specifies the status of the Macie Session. |

Supports update: yes

Discovery: supported
