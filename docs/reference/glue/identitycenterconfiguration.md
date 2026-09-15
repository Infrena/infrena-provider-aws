# aws.identitycenterconfiguration

**CloudFormation type:** `AWS::Glue::IdentityCenterConfiguration`

Resource Type definition for AWS::Glue::IdentityCenterConfiguration

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::Glue::IdentityCenterConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | The identifier for the specified AWS account. |
| `ApplicationArn` | application_arn | `string` | computed |  | The Glue IAM identity center application arn |
| `InstanceArn` | instance_arn | `string` | required, replaces on change |  | The IAM identity center instance arn |
| `Scopes` |  | `list` | optional, computed, provider-chosen |  | The downstream scopes that Glue identity center configuration can access |
| `UserBackgroundSessionsEnabled` | user_background_sessions_enabled | `boolean` | optional, computed, provider-chosen |  | Enable or disable user background sessions for Glue Identity Center |

Supports update: yes

Discovery: not supported
