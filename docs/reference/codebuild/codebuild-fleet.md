# aws.codebuild.fleet

**CloudFormation type:** `AWS::CodeBuild::Fleet`

Resource Type definition for AWS::CodeBuild::Fleet

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CodeBuild::Fleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `BaseCapacity` | base_capacity | `integer` | optional, computed, provider-chosen |  |  |
| `ComputeConfiguration` | compute_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ComputeType` | compute_type | `string` | optional, computed, provider-chosen |  |  |
| `EnvironmentType` | environment_type | `string` | optional, computed, provider-chosen |  |  |
| `FleetProxyConfiguration` | fleet_proxy_configuration | `map` | optional, computed, provider-chosen |  |  |
| `FleetServiceRole` | fleet_service_role | `string` | optional, computed, provider-chosen |  |  |
| `FleetVpcConfig` | fleet_vpc_config | `map` | optional, computed, provider-chosen |  |  |
| `ImageId` | image_id | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `OverflowBehavior` | overflow_behavior | `string` | optional, computed, provider-chosen |  |  |
| `ScalingConfiguration` | scaling_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
