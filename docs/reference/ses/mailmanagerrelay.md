# aws.mailmanagerrelay

**CloudFormation type:** `AWS::SES::MailManagerRelay`

Definition of AWS::SES::MailManagerRelay Resource Type

Region attribute: `region`

**Import ID:** `<region>/RelayId` (AWS::SES::MailManagerRelay)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Authentication` |  | `string` | required |  |  |
| `RelayArn` | relay_arn | `string` | computed |  |  |
| `RelayId` | relay_id | `string` | computed |  |  |
| `RelayName` | relay_name | `string` | optional, computed, provider-chosen |  |  |
| `ServerName` | server_name | `string` | required |  |  |
| `ServerPort` | server_port | `float` | required |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
