# aws.continuousdeploymentpolicy

**CloudFormation type:** `AWS::CloudFront::ContinuousDeploymentPolicy`

Creates a continuous deployment policy that routes a subset of production traffic from a primary distribution to a staging distribution.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::ContinuousDeploymentPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContinuousDeploymentPolicyConfig` | continuous_deployment_policy_config | `map` | required |  | Contains the configuration for a continuous deployment policy. |
| `Id` |  | `string` | computed |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |

Supports update: yes

Discovery: supported
