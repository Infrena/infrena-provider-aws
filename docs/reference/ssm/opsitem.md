# aws.opsitem

**CloudFormation type:** `AWS::SSM::OpsItem`

Resource schema for AWS::SSM::OpsItem.

Region attribute: `region`

**Import ID:** `<region>/OpsItemArn` (AWS::SSM::OpsItem)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Category` |  | `string` | optional, computed, provider-chosen |  | The category of the OpsItem. |
| `CreatedBy` | created_by | `string` | computed |  | The user who created the OpsItem. |
| `CreatedTime` | created_time | `string` | computed |  | The time the OpsItem was created. |
| `Description` |  | `string` | required |  | The description of the OpsItem. |
| `LastModifiedBy` | last_modified_by | `string` | computed |  | The user who last modified the OpsItem. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The time the OpsItem was last modified. |
| `OpsItemArn` | ops_item_arn | `string` | computed |  | The ARN of the OpsItem. |
| `OpsItemId` | ops_item_id | `string` | computed |  | The ID of the OpsItem. |
| `OpsItemType` | ops_item_type | `string` | computed |  | The type of OpsItem. |
| `Priority` |  | `integer` | optional, computed, provider-chosen |  | The priority of the OpsItem. |
| `Severity` |  | `string` | optional, computed, provider-chosen |  | The severity of the OpsItem. |
| `Source` |  | `string` | required, replaces on change |  | The origin of the OpsItem. |
| `Status` |  | `string` | computed |  | The status of the OpsItem. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags for the OpsItem. |
| `Title` |  | `string` | required |  | The title of the OpsItem. |
| `Version` |  | `string` | computed |  | The version of the OpsItem. |

Supports update: yes

Discovery: supported
