# aws.knowledgebasepolicy

**CloudFormation type:** `AWS::Bedrock::KnowledgeBasePolicy`

Definition of AWS::Bedrock::KnowledgeBasePolicy Resource Type

Region attribute: `region`

**Import ID:** `<region>/KnowledgeBaseId` (AWS::Bedrock::KnowledgeBasePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `KnowledgeBaseId` | knowledge_base_id | `string` | required, replaces on change | aws.bedrock.knowledgebase.KnowledgeBaseId | The unique identifier of the knowledge base |
| `PolicyDocument` | policy_document | `map` | required |  | The IAM policy document defining access permissions for the knowledge base |
| `RevisionId` | revision_id | `string` | computed |  | The revision identifier for the policy, used for optimistic concurrency control |

Supports update: yes

Discovery: not supported
