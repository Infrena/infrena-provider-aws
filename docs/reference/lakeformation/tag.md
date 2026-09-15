# aws.tag

**CloudFormation type:** `AWS::LakeFormation::Tag`

A resource schema representing a Lake Formation Tag.

Region attribute: `region`

**Import ID:** `<region>/TagKey` (AWS::LakeFormation::Tag)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CatalogId` | catalog_id | `string` | optional, computed, provider-chosen, replaces on change |  | The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your Lake Formation environment. |
| `TagKey` | tag_key | `string` | required, replaces on change |  | The key-name for the LF-tag. |
| `TagValues` | tag_values | `list` | required |  | A list of possible values an attribute can take. |

Supports update: yes

Discovery: supported
