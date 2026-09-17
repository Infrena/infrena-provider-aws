# aws.tasktemplate

**CloudFormation type:** `AWS::Connect::TaskTemplate`

Resource Type definition for AWS::Connect::TaskTemplate.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Connect::TaskTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The identifier (arn) of the task template. |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen |  | the client token string in uuid format |
| `Constraints` |  | `map` | optional, computed, provider-chosen |  | The constraints for the task template |
| `ContactFlowArn` | contact_flow_arn | `string` | optional, computed, provider-chosen | aws.contactflow.ContactFlowArn | The identifier of the contact flow. |
| `Defaults` |  | `list` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the task template. |
| `Fields` |  | `list` | optional, computed, provider-chosen |  | The list of task template's fields |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier (arn) of the instance. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the task template. |
| `SelfAssignContactFlowArn` | self_assign_contact_flow_arn | `string` | optional, computed, provider-chosen | aws.contactflow.ContactFlowArn | The identifier of the contact flow. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the task template |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |

Supports update: yes

Discovery: supported (parent resource required)
