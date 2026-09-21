# aws.lens

**CloudFormation type:** `AWS::WellArchitected::Lens`

Definition of AWS::WellArchitected::Lens Resource Type

Region attribute: `region`

**Import ID:** `<region>/LensArn` (AWS::WellArchitected::Lens)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | computed |  | The description of the lens. |
| `JSONString` | json_string | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The JSON representation of a lens. |
| `LensArn` | lens_arn | `string` | computed |  | The ARN of the lens. |
| `LensId` | lens_id | `string` | computed |  | The unique identifier of the lens. |
| `LensVersion` | lens_version | `string` | optional, computed, provider-chosen, replaces on change |  | The version of the lens. |
| `Name` |  | `string` | computed |  | The full name of the lens. |
| `Owner` |  | `string` | computed |  | The Amazon Web Services account ID that owns the lens. |
| `Tags` |  | `map` | tags map |  | The tags assigned to the lens. |

Supports update: yes

Discovery: supported
