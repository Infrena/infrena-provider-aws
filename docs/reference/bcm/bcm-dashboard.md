# aws.bcm.dashboard

**CloudFormation type:** `AWS::BCM::Dashboard`

Definition of AWS::BCM::Dashboard Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BCM::Dashboard)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Type` | type_value | `string` | computed |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |
| `Widgets` |  | `list` | required |  |  |

Supports update: yes

Discovery: supported
