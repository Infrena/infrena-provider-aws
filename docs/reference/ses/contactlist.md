# aws.contactlist

**CloudFormation type:** `AWS::SES::ContactList`

Resource schema for AWS::SES::ContactList.

Region attribute: `region`

**Import ID:** `<region>/ContactListName` (AWS::SES::ContactList)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContactListName` | contact_list_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the contact list. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the contact list. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the contact list. |
| `Topics` |  | `list` | optional, computed, provider-chosen |  | The topics associated with the contact list. |

Supports update: yes

Discovery: supported
