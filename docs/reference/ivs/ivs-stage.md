# aws.ivs.stage

**CloudFormation type:** `AWS::IVS::Stage`

Resource Type definition for AWS::IVS::Stage.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::Stage)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActiveSessionId` | active_session_id | `string` | computed |  | ID of the active session within the stage. |
| `Arn` |  | `string` | computed |  | Stage ARN is automatically generated on creation and assigned as the unique identifier. |
| `AutoParticipantRecordingConfiguration` | auto_participant_recording_configuration | `map` | optional, computed, provider-chosen |  | Configuration object for individual participant recording, to attach to the new stage. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | Stage name |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
