# aws.networkfirewall.loggingconfiguration

**CloudFormation type:** `AWS::NetworkFirewall::LoggingConfiguration`

Resource type definition for AWS::NetworkFirewall::LoggingConfiguration

Region attribute: `region`

**Import ID:** `<region>/FirewallArn` (AWS::NetworkFirewall::LoggingConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EnableMonitoringDashboard` | enable_monitoring_dashboard | `boolean` | optional, computed, provider-chosen |  |  |
| `FirewallArn` | firewall_arn | `string` | required, replaces on change | aws.firewall.FirewallArn | A resource ARN. |
| `FirewallName` | firewall_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `LoggingConfiguration` | logging_configuration | `map` | required |  |  |

Supports update: yes

Discovery: not supported
