# aws.logdeliveryconfiguration

**CloudFormation type:** `AWS::Cognito::LogDeliveryConfiguration`

Resource Type definition for AWS::Cognito::LogDeliveryConfiguration

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Cognito::LogDeliveryConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `LogConfigurations` | log_configurations | `list` | optional, computed, provider-chosen |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |

Supports update: yes

Discovery: not supported
