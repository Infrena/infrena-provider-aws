# aws.amplify.webhook

**CloudFormation type:** `AWS::Amplify::Webhook`

Resource Type definition for AWS::Amplify::Webhook. Creates a webhook on an Amplify app.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Amplify::Webhook)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppId` | app_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.amplify.app.AppId | The unique ID for an Amplify app. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) for the webhook. |
| `BranchName` | branch_name | `string` | required |  | The name for a branch that is part of an Amplify app. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description for a webhook. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags for the webhook. |
| `WebhookId` | webhook_id | `string` | computed |  | The unique ID for a webhook. |
| `WebhookUrl` | webhook_url | `string` | computed |  | The URL of the webhook. |

Supports update: yes

Discovery: supported (parent resource required)
