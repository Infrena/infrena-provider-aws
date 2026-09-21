# aws.simulationapplication

**CloudFormation type:** `AWS::RoboMaker::SimulationApplication`

This schema is for testing purpose only.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RoboMaker::SimulationApplication)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CurrentRevisionId` | current_revision_id | `string` | optional, computed, provider-chosen |  | The current revision id. |
| `Environment` |  | `string` | optional, computed, provider-chosen |  | The URI of the Docker image for the robot application. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the simulation application. |
| `RenderingEngine` | rendering_engine | `map` | optional, computed, provider-chosen, write-only |  | Information about a rendering engine. |
| `RobotSoftwareSuite` | robot_software_suite | `map` | required |  | Information about a robot software suite. |
| `SimulationSoftwareSuite` | simulation_software_suite | `map` | required |  | Information about a simulation software suite. |
| `Sources` |  | `list` | optional, computed, provider-chosen, write-only |  | The sources of the simulation application. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
