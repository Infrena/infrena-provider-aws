# aws.prefixlist

**CloudFormation type:** `AWS::EC2::PrefixList`

Resource schema of AWS::EC2::PrefixList Type

Region attribute: `region`

**Import ID:** `<region>/PrefixListId` (AWS::EC2::PrefixList)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddressFamily` | address_family | `string` | required |  | Ip Version of Prefix List. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Prefix List. |
| `Entries` |  | `list` | optional, computed, provider-chosen |  | Entries of Prefix List. |
| `MaxEntries` | max_entries | `integer` | optional, computed, provider-chosen |  | Max Entries of Prefix List. |
| `OwnerId` | owner_id | `string` | computed |  | Owner Id of Prefix List. |
| `PrefixListId` | prefix_list_id | `string` | computed |  | Id of Prefix List. |
| `PrefixListName` | prefix_list_name | `string` | required |  | Name of Prefix List. |
| `Tags` |  | `map` | tags map |  | Tags for Prefix List |
| `Version` |  | `integer` | computed |  | Version of Prefix List. |

Supports update: yes

Discovery: supported
