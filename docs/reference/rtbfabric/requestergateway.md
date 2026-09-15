# aws.requestergateway

**CloudFormation type:** `AWS::RTBFabric::RequesterGateway`

Resource Type definition for AWS::RTBFabric::RequesterGateway Resource Type.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RTBFabric::RequesterGateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActiveLinksCount` | active_links_count | `integer` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DomainName` | domain_name | `string` | computed |  |  |
| `GatewayId` | gateway_id | `string` | computed |  |  |
| `RequesterGatewayStatus` | requester_gateway_status | `string` | computed |  |  |
| `SecurityGroupIds` | security_group_ids | `list` | required | aws.securitygroup.Id | The ID of one or more security groups in order to create a requester gateway. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | The ID of one or more subnets in order to create a requester gateway. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the Requester Gateway. |
| `TotalLinksCount` | total_links_count | `integer` | computed |  |  |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  |  |
| `VpcId` | vpc_id | `string` | required | aws.vpc.VpcId |  |

Supports update: yes

Discovery: supported
