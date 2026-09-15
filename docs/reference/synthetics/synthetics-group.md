# aws.synthetics.group

**CloudFormation type:** `AWS::Synthetics::Group`

Resource Type definition for AWS::Synthetics::Group

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Synthetics::Group)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  | Id of the group. |
| `Name` |  | `string` | required, replaces on change |  | Name of the group. |
| `ResourceArns` | resource_arns | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
