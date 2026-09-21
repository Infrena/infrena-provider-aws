# aws.bedrock.session

**CloudFormation type:** `AWS::Bedrock::Session`

Definition of AWS::Bedrock::Session Resource Type

Region attribute: `region`

**Import ID:** `<region>/SessionArn` (AWS::Bedrock::Session)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp for when the session was created. |
| `EncryptionKeyArn` | encryption_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the KMS key to use to encrypt the session data. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp for when the session was last modified. |
| `SessionArn` | session_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the session. |
| `SessionId` | session_id | `string` | computed |  | The unique identifier of the session in UUID format. |
| `SessionMetadata` | session_metadata | `map` | optional, computed, provider-chosen |  | A map of key-value pairs containing attributes to be persisted across the session. |
| `SessionStatus` | session_status | `string` | computed |  | The current status of the session. |
| `Tags` |  | `map` | tags map |  | A list of tags associated with the session. |

Supports update: yes

Discovery: supported
