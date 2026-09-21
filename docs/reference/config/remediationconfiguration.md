# aws.remediationconfiguration

**CloudFormation type:** `AWS::Config::RemediationConfiguration`

Resource Type definition for AWS::Config::RemediationConfiguration

Region attribute: `region`

**Import ID:** `<region>/ConfigRuleName` (AWS::Config::RemediationConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Automatic` |  | `boolean` | optional, computed, provider-chosen |  |  |
| `ConfigRuleName` | config_rule_name | `string` | required, replaces on change |  |  |
| `ExecutionControls` | execution_controls | `map` | optional, computed, provider-chosen |  |  |
| `MaximumAutomaticAttempts` | maximum_automatic_attempts | `integer` | optional, computed, provider-chosen |  |  |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  |  |
| `ResourceType` | resource_type | `string` | optional, computed, provider-chosen |  |  |
| `RetryAttemptSeconds` | retry_attempt_seconds | `integer` | optional, computed, provider-chosen |  |  |
| `TargetId` | target_id | `string` | required |  |  |
| `TargetType` | target_type | `string` | required |  |  |
| `TargetVersion` | target_version | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
