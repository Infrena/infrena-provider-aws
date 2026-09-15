# aws.emrcontainers.endpoint

**CloudFormation type:** `AWS::EMRContainers::Endpoint`

Resource Schema of AWS::EMRContainers::Endpoint Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::EMRContainers::Endpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the managed endpoint. |
| `AuthProxyUrl` | auth_proxy_url | `string` | computed |  | The auth proxy URL for Spark Connect connections. |
| `CertificateAuthority` | certificate_authority | `map` | computed |  | The certificate authority for the managed endpoint. |
| `ConfigurationOverrides` | configuration_overrides | `map` | optional, computed, provider-chosen, replaces on change |  | The configuration overrides for the managed endpoint. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time when the managed endpoint was created. |
| `ExecutionRoleArn` | execution_role_arn | `string` | required, replaces on change | aws.role.Arn | The execution role ARN for the managed endpoint. |
| `FailureReason` | failure_reason | `string` | computed |  | The reason for a failed managed endpoint. |
| `Id` |  | `string` | computed |  | The ID of the managed endpoint. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the managed endpoint. |
| `ReleaseLabel` | release_label | `string` | required, replaces on change |  | The Amazon EMR release label. |
| `SecurityGroup` | security_group | `string` | computed |  | The security group associated with the managed endpoint. |
| `ServerUrl` | server_url | `string` | computed |  | The server URL of the managed endpoint. |
| `SessionIdleTimeoutInMinutes` | session_idle_timeout_in_minutes | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | The idle timeout in minutes for sessions on the managed endpoint. |
| `State` |  | `string` | computed |  | The state of the managed endpoint. |
| `StateDetails` | state_details | `string` | computed |  | Additional details about the state of the managed endpoint. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this managed endpoint. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the managed endpoint. |
| `VirtualClusterId` | virtual_cluster_id | `string` | required, replaces on change | aws.virtualcluster.Id | The ID of the virtual cluster for which the managed endpoint is created. |

Supports update: yes

Discovery: supported
