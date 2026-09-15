# aws.s3outposts.endpoint

**CloudFormation type:** `AWS::S3Outposts::Endpoint`

Resource Type Definition for AWS::S3Outposts::Endpoint

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::S3Outposts::Endpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessType` | access_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of access for the on-premise network connectivity for the Outpost endpoint. To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the endpoint. |
| `CidrBlock` | cidr_block | `string` | computed |  | The VPC CIDR committed by this endpoint. |
| `CreationTime` | creation_time | `string` | computed |  | The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ) |
| `CustomerOwnedIpv4Pool` | customer_owned_ipv4_pool | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the customer-owned IPv4 pool for the Endpoint. IP addresses will be allocated from this pool for the endpoint. |
| `FailedReason` | failed_reason | `map` | optional, computed, provider-chosen, replaces on change |  | The failure reason, if any, for a create or delete endpoint operation. |
| `Id` |  | `string` | computed |  | The ID of the endpoint. |
| `NetworkInterfaces` | network_interfaces | `list` | computed |  | The network interfaces of the endpoint. |
| `OutpostId` | outpost_id | `string` | required, replaces on change |  | The id of the customer outpost on which the bucket resides. |
| `SecurityGroupId` | security_group_id | `string` | required, replaces on change | aws.securitygroup.Id | The ID of the security group to use with the endpoint. |
| `Status` |  | `string` | computed |  |  |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | The ID of the subnet in the selected VPC. The subnet must belong to the Outpost. |

Supports update: no

Discovery: supported
