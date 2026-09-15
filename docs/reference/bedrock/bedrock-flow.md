# aws.bedrock.flow

**CloudFormation type:** `AWS::Bedrock::Flow`

Definition of AWS::Bedrock::Flow Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Bedrock::Flow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn representation of the Flow |
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp. |
| `CustomerEncryptionKeyArn` | customer_encryption_key_arn | `string` | optional, computed, provider-chosen |  | A KMS key ARN |
| `Definition` |  | `map` | optional, computed, provider-chosen |  | Flow definition |
| `DefinitionS3Location` | definition_s3_location | `map` | optional, computed, provider-chosen, write-only |  | A bucket, key and optional version pointing to an S3 object containing a UTF-8 encoded JSON string Definition with the same schema as the Definition property of this resource |
| `DefinitionString` | definition_string | `string` | optional, computed, provider-chosen, write-only |  | A JSON string containing a Definition with the same schema as the Definition property of this resource |
| `DefinitionSubstitutions` | definition_substitutions | `map` | optional, computed, provider-chosen, write-only |  | When supplied with DefinitionString or DefinitionS3Location, substrings in the definition matching ${keyname} will be replaced with the associated value from this map |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the flow |
| `ExecutionRoleArn` | execution_role_arn | `string` | required | aws.role.Arn | ARN of a IAM role |
| `Id` |  | `string` | computed |  | Identifier for a Flow |
| `Name` |  | `string` | required |  | Name for the flow |
| `Status` |  | `string` | computed |  | Schema Type for Flow APIs |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `TestAliasTags` | test_alias_tags | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp. |
| `Validations` |  | `list` | computed |  | List of flow validations |
| `Version` |  | `string` | computed |  | Draft Version. |

Supports update: yes

Discovery: supported
