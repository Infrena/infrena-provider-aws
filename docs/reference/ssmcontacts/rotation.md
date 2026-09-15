# aws.rotation

**CloudFormation type:** `AWS::SSMContacts::Rotation`

Resource Type definition for AWS::SSMContacts::Rotation.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SSMContacts::Rotation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the rotation. |
| `ContactIds` | contact_ids | `list` | required |  | Members of the rotation |
| `Name` |  | `string` | required |  | Name of the Rotation |
| `Recurrence` |  | `map` | required |  | Information about when an on-call rotation is in effect and how long the rotation period lasts. |
| `StartTime` | start_time | `string` | required |  | Start time of the first shift of Oncall Schedule |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TimeZoneId` | time_zone_id | `string` | required |  | TimeZone Identifier for the Oncall Schedule |

Supports update: yes

Discovery: supported
