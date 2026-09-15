# aws.resourcespecificlogging

**CloudFormation type:** `AWS::IoT::ResourceSpecificLogging`

Resource-specific logging allows you to specify a logging level for a specific thing group.

Region attribute: `region`

**Import ID:** `<region>/TargetId` (AWS::IoT::ResourceSpecificLogging)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LogLevel` | log_level | `string` | required |  | The log level for a specific target. Valid values are: ERROR, WARN, INFO, DEBUG, or DISABLED. |
| `TargetId` | target_id | `string` | computed |  | Unique Id for a Target (TargetType:TargetName), this will be internally built to serve as primary identifier for a log target. |
| `TargetName` | target_name | `string` | required, replaces on change |  | The target name. |
| `TargetType` | target_type | `string` | required, replaces on change |  | The target type. Value must be THING_GROUP, CLIENT_ID, SOURCE_IP, PRINCIPAL_ID, or EVENT_TYPE. |

Supports update: yes

Discovery: supported
