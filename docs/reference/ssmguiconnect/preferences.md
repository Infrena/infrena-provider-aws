# aws.preferences

**CloudFormation type:** `AWS::SSMGuiConnect::Preferences`

Definition of AWS::SSMGuiConnect::Preferences Resource Type

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::SSMGuiConnect::Preferences)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | The AWS Account Id that the preference is associated with, used as the unique identifier for this resource. |
| `ConnectionRecordingPreferences` | connection_recording_preferences | `map` | optional, computed, provider-chosen |  | The set of preferences used for recording RDP connections in the requesting AWS account and AWS Region. This includes details such as which S3 bucket recordings are stored in. |

Supports update: yes

Discovery: supported
