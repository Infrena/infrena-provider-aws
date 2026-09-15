# aws.cloudtrail.dashboard

**CloudFormation type:** `AWS::CloudTrail::Dashboard`

The Amazon CloudTrail dashboard resource allows customers to manage managed dashboards and create custom dashboards. You can manually refresh custom and managed dashboards. For custom dashboards, you can also set up an automatic refresh schedule and modify dashboard widgets.

Region attribute: `region`

**Import ID:** `<region>/DashboardArn` (AWS::CloudTrail::Dashboard)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedTimestamp` | created_timestamp | `string` | computed |  | The timestamp of the dashboard creation. |
| `DashboardArn` | dashboard_arn | `string` | computed |  | The ARN of the dashboard. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the dashboard. |
| `RefreshSchedule` | refresh_schedule | `map` | optional, computed, provider-chosen |  | Configures the automatic refresh schedule for the dashboard. Includes the frequency unit (DAYS or HOURS) and value, as well as the status (ENABLED or DISABLED) of the refresh schedule. |
| `Status` |  | `string` | computed |  | The status of the dashboard. Values are CREATING, CREATED, UPDATING, UPDATED and DELETING. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TerminationProtectionEnabled` | termination_protection_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether the dashboard is protected from termination. |
| `Type` | type_value | `string` | computed |  | The type of the dashboard. Values are CUSTOM and MANAGED. |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  | The timestamp showing when the dashboard was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp. |
| `Widgets` |  | `list` | optional, computed, provider-chosen |  | List of widgets on the dashboard |

Supports update: yes

Discovery: supported
