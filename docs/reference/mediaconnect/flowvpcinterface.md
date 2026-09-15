# aws.flowvpcinterface

**CloudFormation type:** `AWS::MediaConnect::FlowVpcInterface`

Resource schema for AWS::MediaConnect::FlowVpcInterface

Region attribute: `region`

**Import ID:** `<region>/FlowArn|Name` (AWS::MediaConnect::FlowVpcInterface)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FlowArn` | flow_arn | `string` | required, replaces on change | aws.mediaconnect.flow.FlowArn | The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow. |
| `Name` |  | `string` | required, replaces on change |  | Immutable and has to be a unique against other VpcInterfaces in this Flow. |
| `NetworkInterfaceIds` | network_interface_ids | `list` | computed |  | IDs of the network interfaces created in customer's account by MediaConnect. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | Role Arn MediaConnect can assume to create ENIs in customer's account. |
| `SecurityGroupIds` | security_group_ids | `list` | required | aws.securitygroup.Id | Security Group IDs to be used on ENI. |
| `SubnetId` | subnet_id | `string` | required | aws.subnet.SubnetId | Subnet must be in the AZ of the Flow |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Key-value pairs that can be used to tag and organize this VPC network interface. |

Supports update: yes

Discovery: supported (parent resource required)
