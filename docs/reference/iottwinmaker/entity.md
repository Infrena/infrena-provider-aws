# aws.entity

**CloudFormation type:** `AWS::IoTTwinMaker::Entity`

Resource schema for AWS::IoTTwinMaker::Entity

Region attribute: `region`

**Import ID:** `<region>/WorkspaceId|EntityId` (AWS::IoTTwinMaker::Entity)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the entity. |
| `Components` |  | `map` | optional, computed, provider-chosen |  | A map that sets information about a component type. |
| `CompositeComponents` | composite_components | `map` | optional, computed, provider-chosen |  | A map that sets information about a composite component. |
| `CreationDateTime` | creation_date_time | `string` | computed |  | The date and time when the entity was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the entity. |
| `EntityId` | entity_id | `string` | optional, computed, provider-chosen, replaces on change | aws.entity.EntityId | The ID of the entity. |
| `EntityName` | entity_name | `string` | required |  | The name of the entity. |
| `HasChildEntities` | has_child_entities | `boolean` | computed |  | A Boolean value that specifies whether the entity has child entities or not. |
| `ParentEntityId` | parent_entity_id | `string` | optional, computed, provider-chosen | aws.entity.EntityId | The ID of the parent entity. |
| `Status` |  | `map` | computed |  | The current status of the entity. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `UpdateDateTime` | update_date_time | `string` | computed |  | The last date and time when the entity was updated. |
| `WorkspaceId` | workspace_id | `string` | required, replaces on change | aws.iottwinmaker.workspace.WorkspaceId | The ID of the workspace. |

Supports update: yes

Discovery: supported (parent resource required)
