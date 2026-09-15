# aws.simulationapplicationversion

**CloudFormation type:** `AWS::RoboMaker::SimulationApplicationVersion`

AWS::RoboMaker::SimulationApplicationVersion resource creates an AWS RoboMaker SimulationApplicationVersion. This helps you control which code your simulation uses.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RoboMaker::SimulationApplicationVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Application` |  | `string` | required, replaces on change |  |  |
| `ApplicationVersion` | application_version | `string` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CurrentRevisionId` | current_revision_id | `string` | optional, computed, provider-chosen, replaces on change |  | The revision ID of robot application. |

Supports update: no

Discovery: not supported
