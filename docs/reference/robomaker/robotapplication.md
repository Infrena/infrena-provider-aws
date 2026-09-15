# aws.robotapplication

**CloudFormation type:** `AWS::RoboMaker::RobotApplication`

This schema is for testing purpose only.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RoboMaker::RobotApplication)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CurrentRevisionId` | current_revision_id | `string` | optional, computed, provider-chosen |  | The revision ID of robot application. |
| `Environment` |  | `string` | optional, computed, provider-chosen |  | The URI of the Docker image for the robot application. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the robot application. |
| `RobotSoftwareSuite` | robot_software_suite | `map` | required |  | The robot software suite used by the robot application. |
| `Sources` |  | `list` | optional, computed, provider-chosen, write-only |  | The sources of the robot application. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
