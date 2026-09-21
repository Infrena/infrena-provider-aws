# aws.querydefinition

**CloudFormation type:** `AWS::Logs::QueryDefinition`

The resource schema for AWSLogs QueryDefinition

Region attribute: `region`

**Import ID:** `<region>/QueryDefinitionId` (AWS::Logs::QueryDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LogGroupNames` | log_group_names | `list` | optional, computed, provider-chosen |  | Optionally define specific log groups as part of your query definition |
| `Name` |  | `string` | required |  | A name for the saved query definition |
| `Parameters` |  | `list` | optional, computed, provider-chosen |  | Use this parameter to include specific query parameters as part of your query definition. Query parameters are supported only for Logs Insights QL queries. Query parameters allow you to use placeholder variables in your query string that are substituted with values at execution time. Use the {{parameterName}} syntax in your query string to reference a parameter. |
| `QueryDefinitionId` | query_definition_id | `string` | computed |  | Unique identifier of a query definition |
| `QueryLanguage` | query_language | `string` | optional, computed, provider-chosen |  | Query language of the query string. Possible values are CWLI, SQL, PPL, with CWLI being the default. |
| `QueryString` | query_string | `string` | required |  | The query string to use for this definition |

Supports update: yes

Discovery: supported
