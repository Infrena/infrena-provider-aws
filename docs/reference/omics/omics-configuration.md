# aws.omics.configuration

**CloudFormation type:** `AWS::Omics::Configuration`

Resource schema for AWS::Omics::Configuration

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Omics::Configuration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Unique resource identifier for the configuration. |
| `CreationTime` | creation_time | `string` | computed |  | Configuration creation timestamp. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Optional description for the configuration. |
| `Name` |  | `string` | required, replaces on change |  | User-friendly name for the configuration. |
| `RunConfigurations` | run_configurations | `map` | required, replaces on change |  | Required run-specific configurations. |
| `Status` |  | `string` | computed |  | Current configuration status. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | A map of resource tags |
| `Uuid` |  | `string` | computed |  | Unique identifier for the configuration. |

Supports update: no

Discovery: supported
