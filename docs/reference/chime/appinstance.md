# aws.appinstance

**CloudFormation type:** `AWS::Chime::AppInstance`

Resource Type definition for AWS::Chime::AppInstance

Region attribute: `region`

**Import ID:** `<region>/AppInstanceArn` (AWS::Chime::AppInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppInstanceArn` | app_instance_arn | `string` | computed |  | The Amazon Resource Number (ARN) of the AppInstance. |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The time at which an AppInstance was created, as an ISO 8601 timestamp. |
| `LastUpdatedTimestamp` | last_updated_timestamp | `string` | computed |  | The time an AppInstance was last updated, as an ISO 8601 timestamp. |
| `Metadata` |  | `string` | optional, computed, provider-chosen |  | The metadata of the AppInstance. Limited to a 1KB string in UTF-8. |
| `Name` |  | `string` | required |  | The name of the AppInstance. |
| `Tags` |  | `map` | tags map |  | Tags assigned to the AppInstance. |

Supports update: yes

Discovery: supported
