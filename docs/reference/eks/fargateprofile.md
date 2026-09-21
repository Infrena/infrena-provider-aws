# aws.fargateprofile

**CloudFormation type:** `AWS::EKS::FargateProfile`

Resource Schema for AWS::EKS::FargateProfile

Region attribute: `region`

**Import ID:** `<region>/ClusterName|FargateProfileName` (AWS::EKS::FargateProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  | Name of the Cluster |
| `FargateProfileName` | fargate_profile_name | `string` | optional, computed, provider-chosen, replaces on change |  | Name of FargateProfile |
| `PodExecutionRoleArn` | pod_execution_role_arn | `string` | required, replaces on change | aws.role.Arn | The IAM policy arn for pods |
| `Selectors` |  | `list` | required, replaces on change |  |  |
| `Subnets` |  | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
