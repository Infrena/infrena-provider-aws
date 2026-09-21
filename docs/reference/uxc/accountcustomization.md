# aws.accountcustomization

**CloudFormation type:** `AWS::UXC::AccountCustomization`

Resource schema for managing AWS account-level UX customization settings, including account color, visible services, and visible regions.

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::UXC::AccountCustomization)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountColor` | account_color | `string` | optional, computed, provider-chosen |  | The color theme assigned to the account for visual identification in the AWS Console. |
| `AccountId` | account_id | `string` | computed |  | The AWS account ID that this customization belongs to. This is automatically determined from the caller's identity. |
| `VisibleRegions` | visible_regions | `list` | optional, computed, provider-chosen |  | A list of AWS region identifiers visible to the account in the AWS Console. |
| `VisibleServices` | visible_services | `list` | optional, computed, provider-chosen |  | A list of AWS service identifiers visible to the account in the AWS Console. |

Supports update: yes

Discovery: not supported
