# aws.flowentitlement

**CloudFormation type:** `AWS::MediaConnect::FlowEntitlement`

Resource schema for AWS::MediaConnect::FlowEntitlement

Region attribute: `region`

**Import ID:** `<region>/EntitlementArn` (AWS::MediaConnect::FlowEntitlement)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataTransferSubscriberFeePercent` | data_transfer_subscriber_fee_percent | `integer` | optional, computed, provider-chosen, replaces on change |  | Percentage from 0-100 of the data transfer cost to be billed to the subscriber. |
| `Description` |  | `string` | required |  | A description of the entitlement. |
| `Encryption` |  | `map` | optional, computed, provider-chosen |  | Information about the encryption of the flow. |
| `EntitlementArn` | entitlement_arn | `string` | computed |  | The ARN of the entitlement. |
| `EntitlementStatus` | entitlement_status | `string` | optional, computed, provider-chosen |  | An indication of whether the entitlement is enabled. |
| `FlowArn` | flow_arn | `string` | required, replaces on change | aws.mediaconnect.flow.FlowArn | The ARN of the flow. |
| `Name` |  | `string` | required, replaces on change |  | The name of the entitlement. |
| `Subscribers` |  | `list` | required |  | The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Key-value pairs that can be used to tag and organize this flow entitlement. |

Supports update: yes

Discovery: supported (parent resource required)
