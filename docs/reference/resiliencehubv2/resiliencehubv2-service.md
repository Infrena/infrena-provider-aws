# aws.resiliencehubv2.service

**CloudFormation type:** `AWS::ResilienceHubV2::Service`

Creates a resilience-managed service with associated systems, input sources, assertions, and service functions.

Region attribute: `region`

**Import ID:** `<region>/ServiceArn` (AWS::ResilienceHubV2::Service)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Assertions` |  | `list` | optional, computed, provider-chosen |  | Assertions associated with this service. |
| `AssociatedSystems` | associated_systems | `list` | optional, computed, provider-chosen |  | Systems associated with this service. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the service was created. |
| `DependencyDiscovery` | dependency_discovery | `string` | optional, computed, provider-chosen |  | Dependency discovery state. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the service. |
| `EffectivePolicyValues` | effective_policy_values | `map` | computed |  | Effective policy values computed from the associated policy. |
| `InputSources` | input_sources | `list` | optional, computed, provider-chosen |  | Input sources for this service. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The KMS key ID for encrypting service data. |
| `Name` |  | `string` | required, replaces on change |  | The name of the service. |
| `PermissionModel` | permission_model | `map` | optional, computed, provider-chosen |  |  |
| `PolicyArn` | policy_arn | `string` | optional, computed, provider-chosen | aws.resiliencehubv2.policy.PolicyArn | The ARN of the resilience policy to associate. |
| `Regions` |  | `list` | required, replaces on change |  | AWS regions for the service. |
| `ReportConfiguration` | report_configuration | `map` | optional, computed, provider-chosen |  | Configuration for automatic report generation on a Service. |
| `ServiceArn` | service_arn | `string` | computed |  | The ARN of the service. |
| `Tags` |  | `map` | tags map |  | Tags assigned to the service. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the service was last updated. |

Supports update: yes

Discovery: supported
