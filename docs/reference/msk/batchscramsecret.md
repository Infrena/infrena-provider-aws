# aws.batchscramsecret

**CloudFormation type:** `AWS::MSK::BatchScramSecret`

Resource Type definition for AWS::MSK::BatchScramSecret

Region attribute: `region`

**Import ID:** `<region>/ClusterArn` (AWS::MSK::BatchScramSecret)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusterArn` | cluster_arn | `string` | required, replaces on change | aws.msk.cluster.Arn |  |
| `SecretArnList` | secret_arn_list | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
