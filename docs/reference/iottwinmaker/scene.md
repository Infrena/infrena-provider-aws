# aws.scene

**CloudFormation type:** `AWS::IoTTwinMaker::Scene`

Resource schema for AWS::IoTTwinMaker::Scene

Region attribute: `region`

**Import ID:** `<region>/WorkspaceId|SceneId` (AWS::IoTTwinMaker::Scene)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the scene. |
| `Capabilities` |  | `list` | optional, computed, provider-chosen |  | A list of capabilities that the scene uses to render. |
| `ContentLocation` | content_location | `string` | required |  | The relative path that specifies the location of the content definition file. |
| `CreationDateTime` | creation_date_time | `string` | computed |  | The date and time when the scene was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the scene. |
| `GeneratedSceneMetadata` | generated_scene_metadata | `map` | computed |  | A key-value pair of generated scene metadata for the scene. |
| `SceneId` | scene_id | `string` | required, replaces on change | aws.scene.SceneId | The ID of the scene. |
| `SceneMetadata` | scene_metadata | `map` | optional, computed, provider-chosen |  | A key-value pair of scene metadata for the scene. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `UpdateDateTime` | update_date_time | `string` | computed |  | The date and time of the current update. |
| `WorkspaceId` | workspace_id | `string` | required, replaces on change | aws.iottwinmaker.workspace.WorkspaceId | The ID of the scene. |

Supports update: yes

Discovery: supported (parent resource required)
