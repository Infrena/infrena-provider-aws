# aws.collectionindex

**CloudFormation type:** `AWS::OpenSearchServerless::CollectionIndex`

Resource Type definition for AWS::OpenSearchServerless::CollectionIndex

Region attribute: `region`

**Import ID:** `<region>/Id|IndexName` (AWS::OpenSearchServerless::CollectionIndex)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | required, replaces on change |  | The identifier of the collection |
| `IndexName` | index_name | `string` | required, replaces on change |  | The name of the collection index |
| `IndexSchema` | index_schema | `string` | optional, computed, provider-chosen |  | The Mappings for the collection index |

Supports update: yes

Discovery: not supported
