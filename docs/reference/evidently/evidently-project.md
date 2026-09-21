# aws.evidently.project

**CloudFormation type:** `AWS::Evidently::Project`

Resource Type definition for AWS::Evidently::Project

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Evidently::Project)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppConfigResource` | app_config_resource | `map` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `DataDelivery` | data_delivery | `map` | optional, computed, provider-chosen |  | Destinations for data. |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: not supported
