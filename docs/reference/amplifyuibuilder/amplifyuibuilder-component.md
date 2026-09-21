# aws.amplifyuibuilder.component

**CloudFormation type:** `AWS::AmplifyUIBuilder::Component`

Definition of AWS::AmplifyUIBuilder::Component Resource Type

Region attribute: `region`

**Import ID:** `<region>/AppId|EnvironmentName|Id` (AWS::AmplifyUIBuilder::Component)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppId` | app_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `BindingProperties` | binding_properties | `map` | optional, computed, provider-chosen |  |  |
| `Children` |  | `list` | optional, computed, provider-chosen |  |  |
| `CollectionProperties` | collection_properties | `map` | optional, computed, provider-chosen |  |  |
| `ComponentType` | component_type | `string` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `EnvironmentName` | environment_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Events` |  | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `Overrides` |  | `map` | optional, computed, provider-chosen |  |  |
| `Properties` |  | `map` | optional, computed, provider-chosen |  |  |
| `SchemaVersion` | schema_version | `string` | optional, computed, provider-chosen |  |  |
| `SourceId` | source_id | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `Variants` |  | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
