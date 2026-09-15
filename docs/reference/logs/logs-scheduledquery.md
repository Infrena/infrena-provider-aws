# aws.logs.scheduledquery

**CloudFormation type:** `AWS::Logs::ScheduledQuery`

Creates a new Scheduled Query that allows you to define a Logs Insights query that will run on a schedule and configure actions to take with the query results.

Region attribute: `region`

**Import ID:** `<region>/ScheduledQueryArn` (AWS::Logs::ScheduledQuery)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `float` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DestinationConfiguration` | destination_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ExecutionRoleArn` | execution_role_arn | `string` | required | aws.role.Arn |  |
| `LastExecutionStatus` | last_execution_status | `string` | computed |  |  |
| `LastTriggeredTime` | last_triggered_time | `float` | computed |  |  |
| `LastUpdatedTime` | last_updated_time | `float` | computed |  |  |
| `LogGroupIdentifiers` | log_group_identifiers | `list` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `QueryLanguage` | query_language | `string` | required |  |  |
| `QueryString` | query_string | `string` | required |  |  |
| `ScheduleEndTime` | schedule_end_time | `float` | optional, computed, provider-chosen |  |  |
| `ScheduleExpression` | schedule_expression | `string` | required |  |  |
| `ScheduleStartTime` | schedule_start_time | `float` | optional, computed, provider-chosen |  |  |
| `ScheduledQueryArn` | scheduled_query_arn | `string` | computed |  |  |
| `StartTimeOffset` | start_time_offset | `integer` | optional, computed, provider-chosen |  |  |
| `State` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Timezone` |  | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
