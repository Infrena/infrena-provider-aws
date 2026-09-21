# aws.emrserverless.application

**CloudFormation type:** `AWS::EMRServerless::Application`

Resource schema for AWS::EMRServerless::Application Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId` (AWS::EMRServerless::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | computed |  | The ID of the EMR Serverless Application. |
| `Architecture` |  | `string` | optional, computed, provider-chosen |  | The cpu architecture of an application. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the EMR Serverless Application. |
| `AutoStartConfiguration` | auto_start_configuration | `map` | optional, computed, provider-chosen |  | Configuration for Auto Start of Application |
| `AutoStopConfiguration` | auto_stop_configuration | `map` | optional, computed, provider-chosen |  | Configuration for Auto Stop of Application |
| `IdentityCenterConfiguration` | identity_center_configuration | `map` | optional, computed, provider-chosen |  | The IAM IdentityCenter configuration for trusted-identity-propagation on this application. Supported with release labels emr-7.8.0 and above. |
| `ImageConfiguration` | image_configuration | `map` | optional, computed, provider-chosen |  | The image configuration. |
| `InitialCapacity` | initial_capacity | `list` | optional, computed, provider-chosen |  | Initial capacity initialized when an Application is started. |
| `InteractiveConfiguration` | interactive_configuration | `map` | optional, computed, provider-chosen |  |  |
| `MaximumCapacity` | maximum_capacity | `map` | optional, computed, provider-chosen |  | Maximum allowed cumulative resources for an Application. No new resources will be created once the limit is hit. |
| `MonitoringConfiguration` | monitoring_configuration | `map` | optional, computed, provider-chosen |  | Monitoring configuration for batch and interactive JobRun. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | User friendly Application name. |
| `NetworkConfiguration` | network_configuration | `map` | optional, computed, provider-chosen |  | Network Configuration for customer VPC connectivity. |
| `ReleaseLabel` | release_label | `string` | required |  | EMR release label. |
| `RuntimeConfiguration` | runtime_configuration | `list` | optional, computed, provider-chosen |  | Runtime configuration for batch and interactive JobRun. |
| `SchedulerConfiguration` | scheduler_configuration | `map` | optional, computed, provider-chosen |  | The scheduler configuration for batch and streaming jobs running on this application. Supported with release labels emr-7.0.0 and above. |
| `Tags` |  | `map` | tags map |  | Tag map with key and value |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the application |
| `WorkerTypeSpecifications` | worker_type_specifications | `map` | optional, computed, provider-chosen |  | The key-value pairs that specify worker type to WorkerTypeSpecificationInput. This parameter must contain all valid worker types for a Spark or Hive application. Valid worker types include Driver and Executor for Spark applications and HiveDriver and TezTask for Hive applications. You can either set image details in this parameter for each worker type, or in imageConfiguration for all worker types. |

Supports update: yes

Discovery: supported
