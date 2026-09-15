# aws.computeenvironment

**CloudFormation type:** `AWS::Batch::ComputeEnvironment`

Resource Type definition for AWS::Batch::ComputeEnvironment

Region attribute: `region`

**Import ID:** `<region>/ComputeEnvironmentArn` (AWS::Batch::ComputeEnvironment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ComputeEnvironmentArn` | compute_environment_arn | `string` | computed |  |  |
| `ComputeEnvironmentName` | compute_environment_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ComputeResources` | compute_resources | `map` | optional, computed, provider-chosen |  |  |
| `Context` |  | `string` | optional, computed, provider-chosen |  |  |
| `EcsSettings` | ecs_settings | `map` | optional, computed, provider-chosen |  |  |
| `EksConfiguration` | eks_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `ReplaceComputeEnvironment` | replace_compute_environment | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `ServiceRole` | service_role | `string` | optional, computed, provider-chosen |  |  |
| `State` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | A key-value pair to associate with a resource. |
| `Type` | type_value | `string` | required, replaces on change |  |  |
| `UnmanagedvCpus` | unmanagedv_cpus | `integer` | optional, computed, provider-chosen |  |  |
| `UpdatePolicy` | update_policy | `map` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported
