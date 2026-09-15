# aws.connectcampaigns.campaign

**CloudFormation type:** `AWS::ConnectCampaigns::Campaign`

Definition of AWS::ConnectCampaigns::Campaign Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ConnectCampaigns::Campaign)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Connect Campaign Arn |
| `ConnectInstanceArn` | connect_instance_arn | `string` | required, replaces on change |  | Amazon Connect Instance Arn |
| `DialerConfig` | dialer_config | `map` | required |  | The possible types of dialer config parameters |
| `Name` |  | `string` | required |  | Amazon Connect Campaign Name |
| `OutboundCallConfig` | outbound_call_config | `map` | required |  | The configuration used for outbound calls. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |

Supports update: yes

Discovery: supported
