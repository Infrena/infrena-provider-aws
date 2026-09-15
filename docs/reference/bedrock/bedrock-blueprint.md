# aws.bedrock.blueprint

**CloudFormation type:** `AWS::Bedrock::Blueprint`

Definition of AWS::Bedrock::Blueprint Resource Type

Region attribute: `region`

**Import ID:** `<region>/BlueprintArn` (AWS::Bedrock::Blueprint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BlueprintArn` | blueprint_arn | `string` | computed |  | ARN of a Blueprint |
| `BlueprintName` | blueprint_name | `string` | required, replaces on change |  | Name of the Blueprint |
| `BlueprintStage` | blueprint_stage | `string` | computed |  | Stage of the Blueprint |
| `CreationTime` | creation_time | `string` | computed |  | Creation timestamp |
| `KmsEncryptionContext` | kms_encryption_context | `map` | optional, computed, provider-chosen |  | KMS encryption context |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | KMS key identifier |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | Last modified timestamp |
| `Schema` |  | `map` | required |  | Schema of the blueprint |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | List of Tags |
| `Type` | type_value | `string` | required, replaces on change |  | Modality Type |

Supports update: yes

Discovery: supported
