# aws.crl

**CloudFormation type:** `AWS::RolesAnywhere::CRL`

Definition of AWS::RolesAnywhere::CRL Resource Type

Region attribute: `region`

**Import ID:** `<region>/CrlId` (AWS::RolesAnywhere::CRL)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CrlData` | crl_data | `string` | required |  |  |
| `CrlId` | crl_id | `string` | computed |  |  |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TrustAnchorArn` | trust_anchor_arn | `string` | optional, computed, provider-chosen | aws.trustanchor.TrustAnchorArn |  |

Supports update: yes

Discovery: supported
