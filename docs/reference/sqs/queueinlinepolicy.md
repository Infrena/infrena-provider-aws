# aws.queueinlinepolicy

**CloudFormation type:** `AWS::SQS::QueueInlinePolicy`

Schema for SQS QueueInlinePolicy

Region attribute: `region`

**Import ID:** `<region>/Queue` (AWS::SQS::QueueInlinePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `map` | required |  | A policy document that contains permissions to add to the specified SQS queue |
| `Queue` |  | `string` | required, replaces on change |  | The URL of the SQS queue. |

Supports update: yes

Discovery: not supported
