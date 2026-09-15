# aws.greengrassv2.deployment

**CloudFormation type:** `AWS::GreengrassV2::Deployment`

Resource for Greengrass V2 deployment.

Region attribute: `region`

**Import ID:** `<region>/DeploymentId` (AWS::GreengrassV2::Deployment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Components` |  | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `DeploymentId` | deployment_id | `string` | computed |  |  |
| `DeploymentName` | deployment_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DeploymentPolicies` | deployment_policies | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `IotJobConfiguration` | iot_job_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `ParentTargetArn` | parent_target_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `TargetArn` | target_arn | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
