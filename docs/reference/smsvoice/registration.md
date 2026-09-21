# aws.registration

**CloudFormation type:** `AWS::SMSVOICE::Registration`

A registration that has been created.

Region attribute: `region`

**Import ID:** `<region>/RegistrationArn` (AWS::SMSVOICE::Registration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The time when the registration was created, in UNIX epoch time format. |
| `CurrentVersionNumber` | current_version_number | `integer` | computed |  | The current version number of the registration. |
| `RegistrationArn` | registration_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the registration. |
| `RegistrationId` | registration_id | `string` | computed |  | The unique identifier for the registration. |
| `RegistrationStatus` | registration_status | `string` | computed |  | The status of the registration. |
| `RegistrationType` | registration_type | `string` | required, replaces on change |  | The type of registration form to create. |
| `Tags` |  | `map` | tags map |  | An array of tags (key and value pairs) to associate with the registration. |

Supports update: yes

Discovery: supported
