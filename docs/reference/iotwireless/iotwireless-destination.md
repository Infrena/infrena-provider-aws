# aws.iotwireless.destination

**CloudFormation type:** `AWS::IoTWireless::Destination`

Destination's resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoTWireless::Destination)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Destination arn. Returned after successful create. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Destination description |
| `Expression` |  | `string` | required |  | Destination expression |
| `ExpressionType` | expression_type | `string` | required |  | Must be RuleName |
| `Name` |  | `string` | required, replaces on change |  | Unique name of destination |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | AWS role ARN that grants access |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the destination. |

Supports update: yes

Discovery: supported
