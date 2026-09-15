# aws.instanceconnectendpoint

**CloudFormation type:** `AWS::EC2::InstanceConnectEndpoint`

Resource Type definition for AWS::EC2::InstanceConnectEndpoint

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::InstanceConnectEndpoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZone` | availability_zone | `string` | computed |  | The Availability Zone of the EC2 Instance Connect Endpoint |
| `AvailabilityZoneId` | availability_zone_id | `string` | computed |  | The ID of the Availability Zone of the EC2 Instance Connect Endpoint |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The client token of the instance connect endpoint. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time that the EC2 Instance Connect Endpoint was created |
| `Id` |  | `string` | computed |  | The ID of the EC2 Instance Connect Endpoint. |
| `InstanceConnectEndpointArn` | instance_connect_endpoint_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the EC2 Instance Connect Endpoint |
| `NetworkInterfaceIds` | network_interface_ids | `list` | computed |  | The ID of the elastic network interface that Amazon EC2 automatically created when creating the EC2 Instance Connect Endpoint |
| `OwnerId` | owner_id | `string` | computed |  | The ID of the AWS account that created the EC2 Instance Connect Endpoint |
| `PreserveClientIp` | preserve_client_ip | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether your client's IP address is preserved as the source when you connect to a resource. |
| `PublicDnsNames` | public_dns_names | `map` | computed |  | The public DNS names of the endpoint, including IPv4-only and dualstack DNS names. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | The security groups associated with the endpoint. |
| `State` |  | `string` | computed |  | The current state of the EC2 Instance Connect Endpoint |
| `StateMessage` | state_message | `string` | computed |  | The message for the current state of the EC2 Instance Connect Endpoint. Can include a failure message |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | The ID of the subnet in which the EC2 Instance Connect Endpoint was created. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags assigned to the EC2 Instance Connect Endpoint. |
| `VpcId` | vpc_id | `string` | computed |  | The ID of the VPC in which the EC2 Instance Connect Endpoint was created |

Supports update: yes

Discovery: supported
