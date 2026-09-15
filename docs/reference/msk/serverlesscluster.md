# aws.serverlesscluster

**CloudFormation type:** `AWS::MSK::ServerlessCluster`

Resource Type definition for AWS::MSK::ServerlessCluster

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MSK::ServerlessCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ClientAuthentication` | client_authentication | `map` | required, replaces on change |  |  |
| `ClusterName` | cluster_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | A key-value pair to associate with a resource. |
| `VpcConfigs` | vpc_configs | `list` | required, replaces on change |  |  |

Supports update: no

Discovery: supported
