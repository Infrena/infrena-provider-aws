# aws.partnership

**CloudFormation type:** `AWS::B2BI::Partnership`

Definition of AWS::B2BI::Partnership Resource Type

Region attribute: `region`

**Import ID:** `<region>/PartnershipId` (AWS::B2BI::Partnership)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Capabilities` |  | `list` | required |  |  |
| `CapabilityOptions` | capability_options | `map` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Email` |  | `string` | required, replaces on change |  |  |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `PartnershipArn` | partnership_arn | `string` | computed |  |  |
| `PartnershipId` | partnership_id | `string` | computed |  |  |
| `Phone` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ProfileId` | profile_id | `string` | required, replaces on change | aws.b2bi.profile.ProfileId |  |
| `Tags` |  | `map` | tags map |  |  |
| `TradingPartnerId` | trading_partner_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
