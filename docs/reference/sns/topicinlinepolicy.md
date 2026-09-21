# aws.topicinlinepolicy

**CloudFormation type:** `AWS::SNS::TopicInlinePolicy`

Schema for AWS::SNS::TopicInlinePolicy

Region attribute: `region`

**Import ID:** `<region>/TopicArn` (AWS::SNS::TopicInlinePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `map` | required |  | A policy document that contains permissions to add to the specified SNS topics. |
| `TopicArn` | topic_arn | `string` | required, replaces on change | aws.sns.topic.TopicArn | The Amazon Resource Name (ARN) of the topic to which you want to add the policy. |

Supports update: yes

Discovery: not supported
