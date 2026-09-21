# aws.secrettargetattachment

**CloudFormation type:** `AWS::SecretsManager::SecretTargetAttachment`

Resource Type definition for AWS::SecretsManager::SecretTargetAttachment

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::SecretsManager::SecretTargetAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `SecretId` | secret_id | `string` | required, replaces on change | aws.secret.Id |  |
| `TargetId` | target_id | `string` | required |  |  |
| `TargetType` | target_type | `string` | required |  |  |

Supports update: yes

Discovery: supported
