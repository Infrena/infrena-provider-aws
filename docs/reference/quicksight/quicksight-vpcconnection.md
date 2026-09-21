# aws.quicksight.vpcconnection

**CloudFormation type:** `AWS::QuickSight::VPCConnection`

Definition of the AWS::QuickSight::VPCConnection Resource Type.

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|VPCConnectionId` (AWS::QuickSight::VPCConnection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the VPC connection.</p> |
| `AvailabilityStatus` | availability_status | `string` | optional, computed, provider-chosen |  |  |
| `AwsAccountId` | aws_account_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreatedTime` | created_time | `string` | computed |  | <p>The time that the VPC connection was created.</p> |
| `DnsResolvers` | dns_resolvers | `list` | optional, computed, provider-chosen |  |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | <p>The time that the VPC connection was last updated.</p> |
| `Name` |  | `string` | optional, computed, provider-chosen |  |  |
| `NetworkInterfaces` | network_interfaces | `list` | computed |  | <p>A list of network interfaces.</p> |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id |  |
| `Status` |  | `string` | computed |  |  |
| `SubnetIds` | subnet_ids | `list` | optional, computed, provider-chosen, write-only | aws.subnet.SubnetId |  |
| `Tags` |  | `map` | tags map |  |  |
| `VPCConnectionId` | vpc_connection_id | `string` | optional, computed, provider-chosen, replaces on change | aws.quicksight.vpcconnection.VPCConnectionId |  |
| `VPCId` | vpc_id | `string` | computed |  | <p>The Amazon EC2 VPC ID associated with the VPC connection.</p> |

Supports update: yes

Discovery: supported
