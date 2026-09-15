# aws.accountauditconfiguration

**CloudFormation type:** `AWS::IoT::AccountAuditConfiguration`

Configures the Device Defender audit settings for this account. Settings include how audit notifications are sent and which audit checks are enabled or disabled.

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::IoT::AccountAuditConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | required, replaces on change |  | Your 12-digit account ID (used as the primary identifier for the CloudFormation resource). |
| `AuditCheckConfigurations` | audit_check_configurations | `map` | required |  | Specifies which audit checks are enabled and disabled for this account. |
| `AuditNotificationTargetConfigurations` | audit_notification_target_configurations | `map` | optional, computed, provider-chosen |  | Information about the targets to which audit notifications are sent. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The ARN of the role that grants permission to AWS IoT to access information about your devices, policies, certificates and other items as required when performing an audit. |

Supports update: yes

Discovery: supported
