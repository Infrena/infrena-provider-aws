# aws.robot

**CloudFormation type:** `AWS::RoboMaker::Robot`

AWS::RoboMaker::Robot resource creates an AWS RoboMaker Robot.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RoboMaker::Robot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Architecture` |  | `string` | required, replaces on change |  | The target architecture of the robot. |
| `Arn` |  | `string` | computed |  |  |
| `Fleet` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the fleet. |
| `GreengrassGroupId` | greengrass_group_id | `string` | required, replaces on change |  | The Greengrass group id. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name for the robot. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
