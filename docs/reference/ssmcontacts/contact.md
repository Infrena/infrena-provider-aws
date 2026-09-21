# aws.contact

**CloudFormation type:** `AWS::SSMContacts::Contact`

Resource Type definition for AWS::SSMContacts::Contact

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SSMContacts::Contact)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | required, replaces on change |  | Alias of the contact. String value with 20 to 256 characters. Only alphabetical, numeric characters, dash, or underscore allowed. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the contact. |
| `DisplayName` | display_name | `string` | required |  | Name of the contact. String value with 3 to 256 characters. Only alphabetical, space, numeric characters, dash, or underscore allowed. |
| `Plan` |  | `list` | optional, computed, provider-chosen, write-only |  | The stages that an escalation plan or engagement plan engages contacts and contact methods in. |
| `Tags` |  | `map` | tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  | Contact type, which specify type of contact. Currently supported values: “PERSONAL”, “SHARED”, “OTHER“. |

Supports update: yes

Discovery: supported
