# aws.wafv2.loggingconfiguration

**CloudFormation type:** `AWS::WAFv2::LoggingConfiguration`

A WAFv2 Logging Configuration Resource Provider

Region attribute: `region`

**Import ID:** `<region>/ResourceArn` (AWS::WAFv2::LoggingConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LogDestinationConfigs` | log_destination_configs | `list` | required |  | The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL. |
| `LoggingFilter` | logging_filter | `map` | optional, computed, provider-chosen |  | Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation. |
| `ManagedByFirewallManager` | managed_by_firewall_manager | `boolean` | computed |  | Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration. |
| `RedactedFields` | redacted_fields | `list` | optional, computed, provider-chosen |  | The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx. |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs. |

Supports update: yes

Discovery: supported
