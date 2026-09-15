# aws.quickconnect

**CloudFormation type:** `AWS::Connect::QuickConnect`

Resource Type definition for AWS::Connect::QuickConnect

Region attribute: `region`

**Import ID:** `<region>/QuickConnectArn` (AWS::Connect::QuickConnect)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the quick connect. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Name` |  | `string` | required |  | The name of the quick connect. |
| `QuickConnectArn` | quick_connect_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the quick connect. |
| `QuickConnectConfig` | quick_connect_config | `map` | required |  | Configuration settings for the quick connect. |
| `QuickConnectType` | quick_connect_type | `string` | computed |  | The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE). |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |

Supports update: yes

Discovery: supported (parent resource required)
