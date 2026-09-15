# aws.secretsmanager.resourcepolicy

**CloudFormation type:** `AWS::SecretsManager::ResourcePolicy`

Resource Type definition for AWS::SecretsManager::ResourcePolicy

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::SecretsManager::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BlockPublicPolicy` | block_public_policy | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether to block resource-based policies that allow broad access to the secret. |
| `Id` |  | `string` | computed |  | The Arn of the secret. |
| `ResourcePolicy` | resource_policy | `string` | required |  | A JSON-formatted string for an AWS resource-based policy. |
| `SecretId` | secret_id | `string` | required, replaces on change | aws.secret.Id | The ARN or name of the secret to attach the resource-based policy. |

Supports update: yes

Discovery: supported
