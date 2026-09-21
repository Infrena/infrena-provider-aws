# aws.connectcampaignsv2.campaign

**CloudFormation type:** `AWS::ConnectCampaignsV2::Campaign`

Definition of AWS::ConnectCampaignsV2::Campaign Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ConnectCampaignsV2::Campaign)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Connect Campaign Arn |
| `ChannelSubtypeConfig` | channel_subtype_config | `map` | optional, computed, provider-chosen |  | The possible types of channel subtype config parameters |
| `CommunicationLimitsOverride` | communication_limits_override | `map` | optional, computed, provider-chosen |  | Communication limits config |
| `CommunicationTimeConfig` | communication_time_config | `map` | optional, computed, provider-chosen |  | Campaign communication time config |
| `ConnectCampaignFlowArn` | connect_campaign_flow_arn | `string` | optional, computed, provider-chosen |  | Arn |
| `ConnectInstanceId` | connect_instance_id | `string` | required, replaces on change |  | Amazon Connect Instance Id |
| `EntryLimitsConfig` | entry_limits_config | `map` | optional, computed, provider-chosen |  | Entry limits config for a campaign |
| `Name` |  | `string` | required |  | Campaign name |
| `Schedule` |  | `map` | optional, computed, provider-chosen |  | Campaign schedule |
| `Source` |  | `map` | optional, computed, provider-chosen |  | The possible source of the campaign |
| `Tags` |  | `map` | tags map |  | One or more tags. |
| `Type` | type_value | `string` | optional, computed, provider-chosen |  | Campaign type |

Supports update: yes

Discovery: supported (parent resource required)
