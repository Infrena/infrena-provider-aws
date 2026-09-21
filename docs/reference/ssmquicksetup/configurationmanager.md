# aws.configurationmanager

**CloudFormation type:** `AWS::SSMQuickSetup::ConfigurationManager`

Definition of AWS::SSMQuickSetup::ConfigurationManager Resource Type

Region attribute: `region`

**Import ID:** `<region>/ManagerArn` (AWS::SSMQuickSetup::ConfigurationManager)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConfigurationDefinitions` | configuration_definitions | `list` | required |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastModifiedAt` | last_modified_at | `string` | computed |  |  |
| `ManagerArn` | manager_arn | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `StatusSummaries` | status_summaries | `list` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
