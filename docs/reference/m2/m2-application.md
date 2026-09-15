# aws.m2.application

**CloudFormation type:** `AWS::M2::Application`

Represents an application that runs on an AWS Mainframe Modernization Environment

Region attribute: `region`

**Import ID:** `<region>/ApplicationArn` (AWS::M2::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationArn` | application_arn | `string` | computed |  |  |
| `ApplicationId` | application_id | `string` | computed |  |  |
| `Definition` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EngineType` | engine_type | `string` | required, replaces on change |  |  |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID or the Amazon Resource Name (ARN) of the customer managed KMS Key used for encrypting application-related resources. |
| `Name` |  | `string` | required, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
