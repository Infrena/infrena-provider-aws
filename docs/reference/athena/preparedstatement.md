# aws.preparedstatement

**CloudFormation type:** `AWS::Athena::PreparedStatement`

Resource schema for AWS::Athena::PreparedStatement

Region attribute: `region`

**Import ID:** `<region>/StatementName|WorkGroup` (AWS::Athena::PreparedStatement)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the prepared statement. |
| `QueryStatement` | query_statement | `string` | required |  | The query string for the prepared statement. |
| `StatementName` | statement_name | `string` | required, replaces on change |  | The name of the prepared statement. |
| `WorkGroup` | work_group | `string` | required, replaces on change |  | The name of the workgroup to which the prepared statement belongs. |

Supports update: yes

Discovery: supported (parent resource required)
