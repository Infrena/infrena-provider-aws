# aws.form

**CloudFormation type:** `AWS::AmplifyUIBuilder::Form`

Definition of AWS::AmplifyUIBuilder::Form Resource Type

Region attribute: `region`

**Import ID:** `<region>/AppId|EnvironmentName|Id` (AWS::AmplifyUIBuilder::Form)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppId` | app_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Cta` |  | `map` | optional, computed, provider-chosen |  |  |
| `DataType` | data_type | `map` | optional, computed, provider-chosen |  |  |
| `EnvironmentName` | environment_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Fields` |  | `map` | optional, computed, provider-chosen |  |  |
| `FormActionType` | form_action_type | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `LabelDecorator` | label_decorator | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `SchemaVersion` | schema_version | `string` | optional, computed, provider-chosen |  |  |
| `SectionalElements` | sectional_elements | `map` | optional, computed, provider-chosen |  |  |
| `Style` |  | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
