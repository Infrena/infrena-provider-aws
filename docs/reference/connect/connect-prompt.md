# aws.connect.prompt

**CloudFormation type:** `AWS::Connect::Prompt`

Resource Type definition for AWS::Connect::Prompt

Region attribute: `region`

**Import ID:** `<region>/PromptArn` (AWS::Connect::Prompt)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the prompt. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Name` |  | `string` | required |  | The name of the prompt. |
| `PromptArn` | prompt_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the prompt. |
| `S3Uri` | s3_uri | `string` | optional, computed, provider-chosen, write-only |  | S3 URI of the customer's audio file for creating prompts resource.. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
