# aws.dlpsetting

**CloudFormation type:** `AWS::QuickSight::DLPSetting`

Definition of the AWS::QuickSight::DLPSetting Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|DlpSettingId` (AWS::QuickSight::DLPSetting)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AwsAccountId` | aws_account_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DlpSettingId` | dlp_setting_id | `string` | required, replaces on change | aws.dlpsetting.DlpSettingId |  |
| `Enabled` |  | `boolean` | required |  |  |
| `Name` |  | `string` | required |  |  |
| `ProviderConfig` | provider_config | `string` | required |  |  |
| `ProviderOutageAction` | provider_outage_action | `string` | required |  |  |
| `ProviderType` | provider_type | `string` | required |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
