# aws.aps.workspace

**CloudFormation type:** `AWS::APS::Workspace`

Resource Type definition for AWS::APS::Workspace

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::APS::Workspace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AlertManagerDefinition` | alert_manager_definition | `string` | optional, computed, provider-chosen |  | The AMP Workspace alert manager definition data |
| `Alias` |  | `string` | optional, computed, provider-chosen |  | AMP Workspace alias. |
| `Arn` |  | `string` | computed |  | Workspace arn. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | KMS Key ARN used to encrypt and decrypt AMP workspace data. |
| `LoggingConfiguration` | logging_configuration | `map` | optional, computed, provider-chosen |  | Logging configuration |
| `PrometheusEndpoint` | prometheus_endpoint | `string` | computed |  | AMP Workspace prometheus endpoint |
| `QueryLoggingConfiguration` | query_logging_configuration | `map` | optional, computed, provider-chosen |  | Query logging configuration |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `WorkspaceConfiguration` | workspace_configuration | `map` | optional, computed, provider-chosen |  | Workspace configuration |
| `WorkspaceId` | workspace_id | `string` | computed |  | Required to identify a specific APS Workspace. |

Supports update: yes

Discovery: supported
