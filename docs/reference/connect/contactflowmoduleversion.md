# aws.contactflowmoduleversion

**CloudFormation type:** `AWS::Connect::ContactFlowModuleVersion`

Resource Type definition for ContactFlowModuleVersion

Region attribute: `region`

**Import ID:** `<region>/ContactFlowModuleVersionARN` (AWS::Connect::ContactFlowModuleVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContactFlowModuleId` | contact_flow_module_id | `string` | required, replaces on change |  | The identifier of the contact flow module (ARN) this version is tied to. |
| `ContactFlowModuleVersionARN` | contact_flow_module_version_arn | `string` | computed |  | The identifier of the contact flow module version (ARN). |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the version. |
| `FlowModuleContentSha256` | flow_module_content_sha256 | `string` | computed |  | Indicates the checksum value of the latest published flow module content |
| `Version` |  | `integer` | computed |  | The version number of this revision |

Supports update: yes

Discovery: supported (parent resource required)
