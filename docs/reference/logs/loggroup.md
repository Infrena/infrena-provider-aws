# aws.loggroup

**CloudFormation type:** `AWS::Logs::LogGroup`

The ``AWS::Logs::LogGroup`` resource specifies a log group. A log group defines common properties for log streams, such as their retention and access control rules. Each log stream must belong to one log group.

Region attribute: `region`

**Import ID:** `<region>/LogGroupName` (AWS::Logs::LogGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `BearerTokenAuthenticationEnabled` | bearer_token_authentication_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether bearer token authentication is enabled for this log group. When enabled, bearer token authentication is allowed on operations until it is explicitly disabled. |
| `DataProtectionPolicy` | data_protection_policy | `map` | optional, computed, provider-chosen |  | Creates a data protection policy and assigns it to the log group. A data protection policy can help safeguard sensitive data that's ingested by the log group by auditing and masking the sensitive log data. When a user who does not have permission to view masked data views a log event that includes masked data, the sensitive data is replaced by asterisks. |
| `DeletionProtectionEnabled` | deletion_protection_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether deletion protection is enabled for this log group. When enabled, deletion protection blocks all deletion operations until it is explicitly disabled. |
| `FieldIndexPolicies` | field_index_policies | `list` | optional, computed, provider-chosen |  | Creates or updates a *field index policy* for the specified log group. Only log groups in the Standard log class support field index policies. For more information about log classes, see [Log classes](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatch_Logs_Log_Classes.html). |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the KMS key to use when encrypting log data. |
| `LogGroupClass` | log_group_class | `string` | optional, computed, provider-chosen |  | Specifies the log group class for this log group. There are two classes: |
| `LogGroupName` | log_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the log group. If you don't specify a name, CFNlong generates a unique ID for the log group. |
| `ResourcePolicyDocument` | resource_policy_document | `map` | optional, computed, provider-chosen |  | Creates or updates a resource policy for the specified log group that allows other services to put log events to this account. A LogGroup can have 1 resource policy. |
| `RetentionInDays` | retention_in_days | `integer` | optional, computed, provider-chosen |  | The number of days to retain the log events in the specified log group. Possible values are: 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1096, 1827, 2192, 2557, 2922, 3288, and 3653. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to the log group. |

Supports update: yes

Discovery: supported
