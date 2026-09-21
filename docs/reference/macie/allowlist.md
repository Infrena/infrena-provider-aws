# aws.allowlist

**CloudFormation type:** `AWS::Macie::AllowList`

Macie AllowList resource schema

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Macie::AllowList)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | AllowList ARN. |
| `Criteria` |  | `map` | required |  | The regex or s3 object to use for the AllowList. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of AllowList. |
| `Id` |  | `string` | computed |  | AllowList ID. |
| `Name` |  | `string` | required |  | Name of AllowList. |
| `Status` |  | `string` | computed |  | The status for the AllowList |
| `Tags` |  | `map` | tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
