# aws.appinstanceuser

**CloudFormation type:** `AWS::Chime::AppInstanceUser`

Resource Type definition for AWS::Chime::AppInstanceUser

Region attribute: `region`

**Import ID:** `<region>/AppInstanceUserArn` (AWS::Chime::AppInstanceUser)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppInstanceArn` | app_instance_arn | `string` | required, replaces on change | aws.appinstance.AppInstanceArn |  |
| `AppInstanceUserArn` | app_instance_user_arn | `string` | computed |  |  |
| `AppInstanceUserId` | app_instance_user_id | `string` | required, replaces on change | aws.appinstanceuser.AppInstanceUserId |  |
| `ExpirationSettings` | expiration_settings | `map` | optional, computed, provider-chosen |  |  |
| `Metadata` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
