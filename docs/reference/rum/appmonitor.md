# aws.appmonitor

**CloudFormation type:** `AWS::RUM::AppMonitor`

Resource Type definition for AWS::RUM::AppMonitor

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::RUM::AppMonitor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppMonitorConfiguration` | app_monitor_configuration | `map` | optional, computed, provider-chosen |  | AppMonitor configuration |
| `CustomEvents` | custom_events | `map` | optional, computed, provider-chosen |  | AppMonitor custom events configuration |
| `CwLogEnabled` | cw_log_enabled | `boolean` | optional, computed, provider-chosen |  | Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false |
| `DeobfuscationConfiguration` | deobfuscation_configuration | `map` | optional, computed, provider-chosen |  | A structure that contains the configuration for how an app monitor can deobfuscate stack traces. |
| `Domain` |  | `string` | optional, computed, provider-chosen |  | The top-level internet domain name for which your application has administrative authority. The CreateAppMonitor requires either the domain or the domain list. |
| `DomainList` | domain_list | `list` | optional, computed, provider-chosen |  | The top-level internet domain names for which your application has administrative authority. The CreateAppMonitor requires either the domain or the domain list. |
| `Id` |  | `string` | computed |  | The unique ID of the new app monitor. |
| `Name` |  | `string` | required, replaces on change |  | A name for the app monitor |
| `Platform` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ResourcePolicy` | resource_policy | `map` | optional, computed, provider-chosen |  | A structure that defines resource policy attached to your app monitor. |
| `Tags` |  | `map` | tags map |  | Assigns one or more tags (key-value pairs) to the app monitor. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor. |

Supports update: yes

Discovery: supported
