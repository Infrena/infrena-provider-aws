# aws.servicesetting

**CloudFormation type:** `AWS::SSM::ServiceSetting`

Resource Type definition for AWS::SSM::ServiceSetting. ServiceSetting is an account-level setting for an AWS service that defines how a user interacts with or uses a service or feature.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SSM::ServiceSetting)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the service setting. |
| `LastModifiedDate` | last_modified_date | `string` | computed |  | The last time the service setting was modified. |
| `LastModifiedUser` | last_modified_user | `string` | computed |  | The ARN of the last modified user. |
| `SettingId` | setting_id | `string` | required, replaces on change |  | The ID of the service setting, such as /ssm/parameter-store/high-throughput-enabled. |
| `SettingValue` | setting_value | `string` | required |  | The value of the service setting. |
| `Status` |  | `string` | computed |  | The status of the service setting. The value can be Default, Customized or PendingUpdate. |

Supports update: yes

Discovery: supported (parent resource required)
