# aws.endpointauthorization

**CloudFormation type:** `AWS::Redshift::EndpointAuthorization`

Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts.

Region attribute: `region`

**Import ID:** `<region>/ClusterIdentifier|Account` (AWS::Redshift::EndpointAuthorization)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Account` |  | `string` | required, replaces on change |  | The target AWS account ID to grant or revoke access for. |
| `AllowedAllVPCs` | allowed_all_vp_cs | `boolean` | computed |  | Indicates whether all VPCs in the grantee account are allowed access to the cluster. |
| `AllowedVPCs` | allowed_vp_cs | `list` | computed |  | The VPCs allowed access to the cluster. |
| `AuthorizeTime` | authorize_time | `string` | computed |  | The time (UTC) when the authorization was created. |
| `ClusterIdentifier` | cluster_identifier | `string` | required, replaces on change |  | The cluster identifier. |
| `ClusterStatus` | cluster_status | `string` | computed |  | The status of the cluster. |
| `EndpointCount` | endpoint_count | `integer` | computed |  | The number of Redshift-managed VPC endpoints created for the authorization. |
| `Force` |  | `boolean` | optional, computed, provider-chosen, write-only |  | Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted. |
| `Grantee` |  | `string` | computed |  | The AWS account ID of the grantee of the cluster. |
| `Grantor` |  | `string` | computed |  | The AWS account ID of the cluster owner. |
| `Status` |  | `string` | computed |  | The status of the authorization action. |
| `VpcIds` | vpc_ids | `list` | optional, computed, provider-chosen | aws.vpc.VpcId | The virtual private cloud (VPC) identifiers to grant or revoke access to. |

Supports update: yes

Discovery: supported
