# aws.configurationseteventdestination

**CloudFormation type:** `AWS::SES::ConfigurationSetEventDestination`

Resource Type definition for AWS::SES::ConfigurationSetEventDestination

Region attribute: `region`

**Import ID:** `<region>/Id|ConfigurationSetName` (AWS::SES::ConfigurationSetEventDestination)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConfigurationSetName` | configuration_set_name | `string` | required, replaces on change |  | The name of the configuration set that contains the event destination. |
| `EventDestination` | event_destination | `map` | required |  | The event destination object. |
| `Id` |  | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
