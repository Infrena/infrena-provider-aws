# aws.fms.notificationchannel

**CloudFormation type:** `AWS::FMS::NotificationChannel`

Designates the IAM role and Amazon Simple Notification Service (SNS) topic that AWS Firewall Manager uses to record SNS logs.

Region attribute: `region`

**Import ID:** `<region>/SnsTopicArn` (AWS::FMS::NotificationChannel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `SnsRoleName` | sns_role_name | `string` | required |  | A resource ARN. |
| `SnsTopicArn` | sns_topic_arn | `string` | required, replaces on change |  | A resource ARN. |

Supports update: yes

Discovery: supported
