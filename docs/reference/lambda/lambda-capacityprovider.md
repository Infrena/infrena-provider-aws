# aws.lambda.capacityprovider

**CloudFormation type:** `AWS::Lambda::CapacityProvider`

Creates a capacity provider that manages compute resources for Lambda functions

Region attribute: `region`

**Import ID:** `<region>/CapacityProviderName` (AWS::Lambda::CapacityProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CapacityProviderName` | capacity_provider_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CapacityProviderScalingConfig` | capacity_provider_scaling_config | `map` | optional, computed, provider-chosen |  | Configuration that defines how the capacity provider scales compute instances based on demand and policies. |
| `InstanceRequirements` | instance_requirements | `map` | optional, computed, provider-chosen, replaces on change |  | Specifications that define the characteristics and constraints for compute instances used by the capacity provider. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the KMS key used to encrypt the capacity provider's resources. |
| `PermissionsConfig` | permissions_config | `map` | required, replaces on change |  | Configuration that specifies the permissions required for the capacity provider to manage compute resources. |
| `PropagateTags` | propagate_tags | `map` | optional, computed, provider-chosen |  | Configuration that defines how tags are propagated to managed resources. |
| `State` |  | `string` | computed |  | The current state of the capacity provider. Indicates whether the provider is being created, is active and ready for use, has failed, or is being deleted. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A key-value pair that provides metadata for the capacity provider. |
| `TelemetryConfig` | telemetry_config | `map` | optional, computed, provider-chosen |  | Configuration that specifies the telemetry collection for the capacity provider. |
| `VpcConfig` | vpc_config | `map` | required, replaces on change |  | VPC configuration that specifies the network settings for compute instances managed by the capacity provider. |

Supports update: yes

Discovery: supported
