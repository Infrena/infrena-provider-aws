# aws.contactflow

**CloudFormation type:** `AWS::Connect::ContactFlow`

Resource Type definition for AWS::Connect::ContactFlow

Region attribute: `region`

**Import ID:** `<region>/ContactFlowArn` (AWS::Connect::ContactFlow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContactFlowArn` | contact_flow_arn | `string` | computed |  | The identifier of the contact flow (ARN). |
| `Content` |  | `string` | required |  | The content of the contact flow in JSON format. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the contact flow. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance (ARN). |
| `Name` |  | `string` | required |  | The name of the contact flow. |
| `State` |  | `string` | optional, computed, provider-chosen |  | The state of the contact flow. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the contact flow. |

Supports update: yes

Discovery: supported (parent resource required)
