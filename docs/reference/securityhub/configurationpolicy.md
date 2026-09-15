# aws.configurationpolicy

**CloudFormation type:** `AWS::SecurityHub::ConfigurationPolicy`

The AWS::SecurityHub::ConfigurationPolicy resource represents the Central Configuration Policy in your account.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SecurityHub::ConfigurationPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the configuration policy. |
| `ConfigurationPolicy` | configuration_policy | `map` | required |  | An object that defines how Security Hub is configured. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time, in UTC and ISO 8601 format. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the configuration policy. |
| `Id` |  | `string` | computed |  | The universally unique identifier (UUID) of the configuration policy. |
| `Name` |  | `string` | required |  | The name of the configuration policy. |
| `ServiceEnabled` | service_enabled | `boolean` | computed |  | Indicates whether the service that the configuration policy applies to is enabled in the policy. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time, in UTC and ISO 8601 format. |

Supports update: yes

Discovery: supported
