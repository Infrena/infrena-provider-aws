# aws.contactflowmodule

**CloudFormation type:** `AWS::Connect::ContactFlowModule`

Resource Type definition for AWS::Connect::ContactFlowModule.

Region attribute: `region`

**Import ID:** `<region>/ContactFlowModuleArn` (AWS::Connect::ContactFlowModule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContactFlowModuleArn` | contact_flow_module_arn | `string` | computed |  | The identifier of the contact flow module (ARN). |
| `Content` |  | `string` | required |  | The content of the contact flow module in JSON format. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the contact flow module. |
| `ExternalInvocationConfiguration` | external_invocation_configuration | `map` | optional, computed, provider-chosen |  | Defines the external invocation configuration of the flow module resource |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance (ARN). |
| `Name` |  | `string` | required |  | The name of the contact flow module. |
| `Settings` |  | `string` | optional, computed, provider-chosen |  | The schema of the settings for contact flow module in JSON Schema V4 format. |
| `State` |  | `string` | optional, computed, provider-chosen |  | The state of the contact flow module. |
| `Status` |  | `string` | computed |  | The status of the contact flow module. |
| `Tags` |  | `map` | tags map |  | One or more tags. |

Supports update: yes

Discovery: supported (parent resource required)
