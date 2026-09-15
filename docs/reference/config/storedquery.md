# aws.storedquery

**CloudFormation type:** `AWS::Config::StoredQuery`

Resource Type definition for AWS::Config::StoredQuery

Region attribute: `region`

**Import ID:** `<region>/QueryName` (AWS::Config::StoredQuery)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `QueryArn` | query_arn | `string` | computed |  |  |
| `QueryDescription` | query_description | `string` | optional, computed, provider-chosen |  |  |
| `QueryExpression` | query_expression | `string` | required |  |  |
| `QueryId` | query_id | `string` | computed |  |  |
| `QueryName` | query_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the stored query. |

Supports update: yes

Discovery: supported
