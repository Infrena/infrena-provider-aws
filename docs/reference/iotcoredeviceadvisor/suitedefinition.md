# aws.suitedefinition

**CloudFormation type:** `AWS::IoTCoreDeviceAdvisor::SuiteDefinition`

An example resource schema demonstrating some basic constructs and validation rules.

Region attribute: `region`

**Import ID:** `<region>/SuiteDefinitionId` (AWS::IoTCoreDeviceAdvisor::SuiteDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `SuiteDefinitionArn` | suite_definition_arn | `string` | computed |  | The Amazon Resource name for the suite definition. |
| `SuiteDefinitionConfiguration` | suite_definition_configuration | `map` | required |  |  |
| `SuiteDefinitionId` | suite_definition_id | `string` | computed |  | The unique identifier for the suite definition. |
| `SuiteDefinitionVersion` | suite_definition_version | `string` | computed |  | The suite definition version of a test suite. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
