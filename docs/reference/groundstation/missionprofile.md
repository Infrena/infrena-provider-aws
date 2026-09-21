# aws.missionprofile

**CloudFormation type:** `AWS::GroundStation::MissionProfile`

AWS Ground Station Mission Profile resource type for CloudFormation.

Region attribute: `aws_region`

**Import ID:** `<region>/Id|Arn` (AWS::GroundStation::MissionProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ContactPostPassDurationSeconds` | contact_post_pass_duration_seconds | `integer` | optional, computed, provider-chosen |  | Post-pass time needed after the contact. |
| `ContactPrePassDurationSeconds` | contact_pre_pass_duration_seconds | `integer` | optional, computed, provider-chosen |  | Pre-pass time needed before the contact. |
| `DataflowEdges` | dataflow_edges | `list` | required |  |  |
| `Id` |  | `string` | computed |  |  |
| `MinimumViableContactDurationSeconds` | minimum_viable_contact_duration_seconds | `integer` | required |  | Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts. |
| `Name` |  | `string` | required |  | A name used to identify a mission profile. |
| `Region` |  | `string` | computed |  |  |
| `StreamsKmsKey` | streams_kms_key | `map` | optional, computed, provider-chosen |  | The ARN of a KMS Key used for encrypting data during transmission from the source to destination locations. |
| `StreamsKmsRole` | streams_kms_role | `string` | optional, computed, provider-chosen |  | The ARN of the KMS Key or Alias Key role used to define permissions on KMS Key usage. |
| `Tags` |  | `map` | tags map |  |  |
| `TelemetrySinkConfigArn` | telemetry_sink_config_arn | `string` | optional, computed, provider-chosen | aws.config.Arn | ARN of a Config resource of type TelemetrySinkConfig used for telemetry data sink configuration. |
| `TrackingConfigArn` | tracking_config_arn | `string` | required | aws.config.Arn |  |

Supports update: yes

Discovery: supported
