# aws.ingestconfiguration

**CloudFormation type:** `AWS::IVS::IngestConfiguration`

Resource Type definition for AWS::IVS::IngestConfiguration

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::IngestConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | IngestConfiguration ARN is automatically generated on creation and assigned as the unique identifier. |
| `IngestProtocol` | ingest_protocol | `string` | optional, computed, provider-chosen, replaces on change |  | Ingest Protocol. |
| `InsecureIngest` | insecure_ingest | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Whether ingest configuration allows insecure ingest. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | IngestConfiguration |
| `ParticipantId` | participant_id | `string` | computed |  | Participant Id is automatically generated on creation and assigned. |
| `StageArn` | stage_arn | `string` | optional, computed, provider-chosen | aws.ivs.stage.Arn | Stage ARN. A value other than an empty string indicates that stage is linked to IngestConfiguration. Default: "" (recording is disabled). |
| `State` |  | `string` | computed |  | State of IngestConfiguration which determines whether IngestConfiguration is in use or not. |
| `StreamKey` | stream_key | `string` | computed |  | Stream-key value. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the asset model. |
| `UserId` | user_id | `string` | optional, computed, provider-chosen, replaces on change |  | User defined indentifier for participant associated with IngestConfiguration. |

Supports update: yes

Discovery: supported
