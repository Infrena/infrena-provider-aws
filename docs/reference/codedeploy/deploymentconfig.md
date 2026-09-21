# aws.deploymentconfig

**CloudFormation type:** `AWS::CodeDeploy::DeploymentConfig`

Resource Type definition for AWS::CodeDeploy::DeploymentConfig

Region attribute: `region`

**Import ID:** `<region>/DeploymentConfigName` (AWS::CodeDeploy::DeploymentConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ComputePlatform` | compute_platform | `string` | optional, computed, provider-chosen, replaces on change |  | The destination platform type for the deployment (Lambda, Server, or ECS). |
| `DeploymentConfigName` | deployment_config_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the deployment configuration. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment configuration name. For more information, see Name Type. |
| `MinimumHealthyHosts` | minimum_healthy_hosts | `map` | optional, computed, provider-chosen, replaces on change |  | The minimum number of healthy instances that should be available at any time during the deployment. There are two parameters expected in the input: type and value. |
| `TrafficRoutingConfig` | traffic_routing_config | `map` | optional, computed, provider-chosen, replaces on change |  | The configuration that specifies how the deployment traffic is routed. |
| `ZonalConfig` | zonal_config | `map` | optional, computed, provider-chosen, replaces on change |  | The zonal deployment config that specifies how the zonal deployment behaves |

Supports update: no

Discovery: supported
