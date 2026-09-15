# aws.glue.connection

**CloudFormation type:** `AWS::Glue::Connection`

Resource Type definition for AWS::Glue::Connection

Region attribute: `region`

**Import ID:** `<region>/CatalogId|Name` (AWS::Glue::Connection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CatalogId` | catalog_id | `string` | required, replaces on change | aws.catalog.CatalogId | The ID of the data catalog to create the catalog object in. Currently, this should be the AWS account ID. |
| `ConnectionInput` | connection_input | `map` | required |  | The connection properties used for this connection. |
| `Name` |  | `string` | computed |  | The name of the connection. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The collection of tags. Each tag element is associated with a given resource. |

Supports update: yes

Discovery: supported (parent resource required)
