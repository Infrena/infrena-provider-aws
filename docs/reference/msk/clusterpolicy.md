# aws.clusterpolicy

**CloudFormation type:** `AWS::MSK::ClusterPolicy`

Resource Type definition for AWS::MSK::ClusterPolicy

Region attribute: `region`

**Import ID:** `<region>/ClusterArn` (AWS::MSK::ClusterPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusterArn` | cluster_arn | `string` | required, replaces on change | aws.msk.cluster.Arn | The arn of the cluster for the resource policy. |
| `CurrentVersion` | current_version | `string` | computed |  | The current version of the policy attached to the specified cluster |
| `Policy` |  | `map` | required |  | A policy document containing permissions to add to the specified cluster. |

Supports update: yes

Discovery: supported (parent resource required)
