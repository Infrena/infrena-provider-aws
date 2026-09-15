# aws.contactflowversion

**CloudFormation type:** `AWS::Connect::ContactFlowVersion`

Resource Type Definition for ContactFlowVersion

Region attribute: `region`

**Import ID:** `<region>/ContactFlowVersionARN` (AWS::Connect::ContactFlowVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContactFlowId` | contact_flow_id | `string` | required, replaces on change |  | The ARN of the contact flow this version is tied to. |
| `ContactFlowVersionARN` | contact_flow_version_arn | `string` | computed |  | The identifier of the contact flow version (ARN). |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the version. |
| `FlowContentSha256` | flow_content_sha256 | `string` | computed |  | Indicates the checksum value of the latest published flow content |
| `Version` |  | `integer` | computed |  | The version number of this revision |

Supports update: yes

Discovery: supported (parent resource required)
