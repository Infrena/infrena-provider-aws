# aws.route53recoverycontrol.cluster

**CloudFormation type:** `AWS::Route53RecoveryControl::Cluster`

AWS Route53 Recovery Control Cluster resource schema

Region attribute: `region`

**Import ID:** `<region>/ClusterArn` (AWS::Route53RecoveryControl::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusterArn` | cluster_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the cluster. |
| `ClusterEndpoints` | cluster_endpoints | `list` | computed |  | Endpoints for the cluster. |
| `Name` |  | `string` | required, replaces on change |  | Name of a Cluster. You can use any non-white space character in the name |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen |  | Cluster supports IPv4 endpoints and Dual-stack IPv4 and IPv6 endpoints. NetworkType can be IPV4 or DUALSTACK. |
| `Status` |  | `string` | computed |  | Deployment status of a resource. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
