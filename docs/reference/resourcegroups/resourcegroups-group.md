# aws.resourcegroups.group

**CloudFormation type:** `AWS::ResourceGroups::Group`

Schema for ResourceGroups::Group

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::ResourceGroups::Group)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Resource Group ARN. |
| `Configuration` |  | `list` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the resource group |
| `Name` |  | `string` | required, replaces on change |  | The name of the resource group |
| `ResourceQuery` | resource_query | `map` | optional, computed, provider-chosen |  |  |
| `Resources` |  | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
