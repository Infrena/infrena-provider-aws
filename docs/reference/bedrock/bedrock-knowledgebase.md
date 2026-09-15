# aws.bedrock.knowledgebase

**CloudFormation type:** `AWS::Bedrock::KnowledgeBase`

Definition of AWS::Bedrock::KnowledgeBase Resource Type

Region attribute: `region`

**Import ID:** `<region>/KnowledgeBaseId` (AWS::Bedrock::KnowledgeBase)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time at which the knowledge base was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the Resource. |
| `FailureReasons` | failure_reasons | `list` | computed |  | A list of reasons that the API operation on the knowledge base failed. |
| `KnowledgeBaseArn` | knowledge_base_arn | `string` | computed |  | The ARN of the knowledge base. |
| `KnowledgeBaseConfiguration` | knowledge_base_configuration | `map` | required |  | Contains details about the embeddings model used for the knowledge base. |
| `KnowledgeBaseId` | knowledge_base_id | `string` | computed |  | The unique identifier of the knowledge base. |
| `Name` |  | `string` | required |  | The name of the knowledge base. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The ARN of the IAM role with permissions to invoke API operations on the knowledge base. The ARN must begin with AmazonBedrockExecutionRoleForKnowledgeBase_ |
| `Status` |  | `string` | computed |  | The status of a knowledge base. |
| `StorageConfiguration` | storage_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The vector store service in which the knowledge base is stored. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `UpdatedAt` | updated_at | `string` | computed |  | The time at which the knowledge base was last updated. |

Supports update: yes

Discovery: supported
