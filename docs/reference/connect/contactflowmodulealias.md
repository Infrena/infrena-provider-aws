# aws.contactflowmodulealias

**CloudFormation type:** `AWS::Connect::ContactFlowModuleAlias`

Resource Type definition for ContactFlowModuleAlias

Region attribute: `region`

**Import ID:** `<region>/ContactFlowModuleAliasARN` (AWS::Connect::ContactFlowModuleAlias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AliasId` | alias_id | `string` | computed |  | The unique identifier of the alias. |
| `ContactFlowModuleAliasARN` | contact_flow_module_alias_arn | `string` | computed |  | The identifier of the contact flow module alias (ARN). This is constructed from the ContactFlowModuleArn and AliasId. |
| `ContactFlowModuleId` | contact_flow_module_id | `string` | required, replaces on change |  | The identifier of the contact flow module (ARN) this alias is tied to. |
| `ContactFlowModuleVersion` | contact_flow_module_version | `integer` | required |  | The version number of the contact flow module this alias points to. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the alias. |
| `Name` |  | `string` | required |  | The name of the alias. |

Supports update: yes

Discovery: supported (parent resource required)
