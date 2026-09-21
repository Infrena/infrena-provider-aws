# aws.componenttype

**CloudFormation type:** `AWS::IoTTwinMaker::ComponentType`

Resource schema for AWS::IoTTwinMaker::ComponentType

Region attribute: `region`

**Import ID:** `<region>/WorkspaceId|ComponentTypeId` (AWS::IoTTwinMaker::ComponentType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the component type. |
| `ComponentTypeId` | component_type_id | `string` | required, replaces on change | aws.componenttype.ComponentTypeId | The ID of the component type. |
| `CompositeComponentTypes` | composite_component_types | `map` | optional, computed, provider-chosen |  | An map of the composite component types in the component type. Each composite component type's key must be unique to this map. |
| `CreationDateTime` | creation_date_time | `string` | computed |  | The date and time when the component type was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the component type. |
| `ExtendsFrom` | extends_from | `list` | optional, computed, provider-chosen |  | Specifies the parent component type to extend. |
| `Functions` |  | `map` | optional, computed, provider-chosen |  | a Map of functions in the component type. Each function's key must be unique to this map. |
| `IsAbstract` | is_abstract | `boolean` | computed |  | A Boolean value that specifies whether the component type is abstract. |
| `IsSchemaInitialized` | is_schema_initialized | `boolean` | computed |  | A Boolean value that specifies whether the component type has a schema initializer and that the schema initializer has run. |
| `IsSingleton` | is_singleton | `boolean` | optional, computed, provider-chosen |  | A Boolean value that specifies whether an entity can have more than one component of this type. |
| `PropertyDefinitions` | property_definitions | `map` | optional, computed, provider-chosen |  | An map of the property definitions in the component type. Each property definition's key must be unique to this map. |
| `PropertyGroups` | property_groups | `map` | optional, computed, provider-chosen |  | An map of the property groups in the component type. Each property group's key must be unique to this map. |
| `Status` |  | `map` | computed |  | The current status of the component type. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of key-value pairs to associate with a resource. |
| `UpdateDateTime` | update_date_time | `string` | computed |  | The last date and time when the component type was updated. |
| `WorkspaceId` | workspace_id | `string` | required, replaces on change | aws.iottwinmaker.workspace.WorkspaceId | The ID of the workspace that contains the component type. |

Supports update: yes

Discovery: supported (parent resource required)
