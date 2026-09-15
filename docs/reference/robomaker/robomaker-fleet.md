# aws.robomaker.fleet

**CloudFormation type:** `AWS::RoboMaker::Fleet`

AWS::RoboMaker::Fleet resource creates an AWS RoboMaker fleet. Fleets contain robots and can receive deployments.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RoboMaker::Fleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the fleet. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
