# aws.securityagent.application

**CloudFormation type:** `AWS::SecurityAgent::Application`

Resource Type definition for AWS::SecurityAgent::Application

Region attribute: `region`

**Import ID:** `<region>/ApplicationId` (AWS::SecurityAgent::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | computed |  |  |
| `ApplicationName` | application_name | `string` | computed |  |  |
| `DefaultKmsKeyId` | default_kms_key_id | `string` | optional, computed, provider-chosen |  | Identifier of a KMS key. Can be a key ID, key ARN, alias name, or alias ARN. |
| `Domain` |  | `string` | computed |  |  |
| `IdCConfiguration` | id_c_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags for the application |

Supports update: yes

Discovery: supported
