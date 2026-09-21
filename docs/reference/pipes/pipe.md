# aws.pipe

**CloudFormation type:** `AWS::Pipes::Pipe`

Definition of AWS::Pipes::Pipe Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Pipes::Pipe)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `CurrentState` | current_state | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DesiredState` | desired_state | `string` | optional, computed, provider-chosen |  |  |
| `Enrichment` |  | `string` | optional, computed, provider-chosen |  |  |
| `EnrichmentParameters` | enrichment_parameters | `map` | optional, computed, provider-chosen |  |  |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | optional, computed, provider-chosen |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |
| `LogConfiguration` | log_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn |  |
| `Source` |  | `string` | required, replaces on change |  |  |
| `SourceParameters` | source_parameters | `map` | optional, computed, provider-chosen, write-only |  |  |
| `StateReason` | state_reason | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `Target` |  | `string` | required |  |  |
| `TargetParameters` | target_parameters | `map` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported
