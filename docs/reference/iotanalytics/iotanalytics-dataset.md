# aws.iotanalytics.dataset

**CloudFormation type:** `AWS::IoTAnalytics::Dataset`

Resource Type definition for AWS::IoTAnalytics::Dataset

Region attribute: `region`

**Import ID:** `<region>/DatasetName` (AWS::IoTAnalytics::Dataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  |  |
| `ContentDeliveryRules` | content_delivery_rules | `list` | optional, computed, provider-chosen |  |  |
| `DatasetName` | dataset_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `LateDataRules` | late_data_rules | `list` | optional, computed, provider-chosen |  |  |
| `RetentionPeriod` | retention_period | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  |  |
| `Triggers` |  | `list` | optional, computed, provider-chosen |  |  |
| `VersioningConfiguration` | versioning_configuration | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
