# aws.bedrockagentcore.dataset

**CloudFormation type:** `AWS::BedrockAgentCore::Dataset`

Definition of AWS::BedrockAgentCore::Dataset Resource Type

Region attribute: `region`

**Import ID:** `<region>/DatasetArn` (AWS::BedrockAgentCore::Dataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the dataset was created. |
| `DatasetArn` | dataset_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the dataset. |
| `DatasetId` | dataset_id | `string` | computed |  | The unique identifier of the dataset. |
| `DatasetName` | dataset_name | `string` | required, replaces on change |  | Human-readable name for the dataset. Unique within the account (case-insensitive). Immutable after creation. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the dataset. |
| `ExampleCount` | example_count | `integer` | computed |  | The number of examples in the dataset DRAFT. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | Optional AWS KMS key ARN for SSE-KMS on service S3 writes. |
| `SchemaType` | schema_type | `string` | required, replaces on change |  | Versioned schema type governing the structure of examples. Immutable after creation. |
| `Source` |  | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Source of initial examples. Provide either inline examples or an S3 URI pointing to a JSONL file. |
| `Status` |  | `string` | computed |  | The current status of the dataset. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to assign to the dataset. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the dataset was last updated. |

Supports update: yes

Discovery: supported
