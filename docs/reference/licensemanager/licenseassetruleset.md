# aws.licenseassetruleset

**CloudFormation type:** `AWS::LicenseManager::LicenseAssetRuleSet`

Resource schema for AWS::LicenseManager::LicenseAssetRuleSet.

Region attribute: `region`

**Import ID:** `<region>/LicenseAssetRulesetArn` (AWS::LicenseManager::LicenseAssetRuleSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | License asset ruleset description. |
| `LicenseAssetRulesetArn` | license_asset_ruleset_arn | `string` | computed |  | Amazon Resource Name (ARN) of the license asset ruleset. |
| `Name` |  | `string` | required |  | License asset ruleset name. |
| `Rules` |  | `list` | required |  | License asset rules. |
| `Tags` |  | `map` | tags map |  | Tags to add to the license asset ruleset. |

Supports update: yes

Discovery: supported
