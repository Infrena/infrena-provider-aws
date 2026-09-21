# aws.robotapplicationversion

**CloudFormation type:** `AWS::RoboMaker::RobotApplicationVersion`

AWS::RoboMaker::RobotApplicationVersion resource creates an AWS RoboMaker RobotApplicationVersion. This helps you control which code your robot uses.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RoboMaker::RobotApplicationVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Application` |  | `string` | required, replaces on change |  |  |
| `ApplicationVersion` | application_version | `string` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CurrentRevisionId` | current_revision_id | `string` | optional, computed, provider-chosen, replaces on change |  | The revision ID of robot application. |

Supports update: no

Discovery: not supported
