# aws.namedquery

**CloudFormation type:** `AWS::Athena::NamedQuery`

Resource schema for AWS::Athena::NamedQuery

Region attribute: `region`

**Import ID:** `<region>/NamedQueryId` (AWS::Athena::NamedQuery)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Database` |  | `string` | required, replaces on change |  | The database to which the query belongs. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The query description. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The query name. |
| `NamedQueryId` | named_query_id | `string` | computed |  | The unique ID of the query. |
| `QueryString` | query_string | `string` | required, replaces on change |  | The contents of the query with all query statements. |
| `WorkGroup` | work_group | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the workgroup that contains the named query. |

Supports update: no

Discovery: supported
