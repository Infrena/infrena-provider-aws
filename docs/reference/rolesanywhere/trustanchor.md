# aws.trustanchor

**CloudFormation type:** `AWS::RolesAnywhere::TrustAnchor`

Definition of AWS::RolesAnywhere::TrustAnchor Resource Type.

Region attribute: `region`

**Import ID:** `<region>/TrustAnchorId` (AWS::RolesAnywhere::TrustAnchor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `NotificationSettings` | notification_settings | `list` | optional, computed, provider-chosen |  |  |
| `Source` |  | `map` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TrustAnchorArn` | trust_anchor_arn | `string` | computed |  |  |
| `TrustAnchorId` | trust_anchor_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
