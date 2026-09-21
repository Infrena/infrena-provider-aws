# aws.fms.resourceset

**CloudFormation type:** `AWS::FMS::ResourceSet`

Creates an AWS Firewall Manager resource set.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::FMS::ResourceSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  | A Base62 ID |
| `Name` |  | `string` | required |  |  |
| `ResourceTypeList` | resource_type_list | `list` | required |  |  |
| `Resources` |  | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
