# aws.launchconfiguration

**CloudFormation type:** `AWS::AutoScaling::LaunchConfiguration`

The AWS::AutoScaling::LaunchConfiguration resource specifies the launch configuration that can be used by an Auto Scaling group to configure Amazon EC2 instances.

Region attribute: `region`

**Import ID:** `<region>/LaunchConfigurationName` (AWS::AutoScaling::LaunchConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociatePublicIpAddress` | associate_public_ip_address | `boolean` | optional, computed, provider-chosen, replaces on change |  | For Auto Scaling groups that are running in a virtual private cloud (VPC), specifies whether to assign a public IP address to the group's instances. |
| `BlockDeviceMappings` | block_device_mappings | `list` | optional, computed, provider-chosen, replaces on change |  | Specifies how block devices are exposed to the instance. You can specify virtual devices and EBS volumes. |
| `ClassicLinkVPCId` | classic_link_vpc_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.vpc.VpcId | The ID of a ClassicLink-enabled VPC to link your EC2-Classic instances to. |
| `ClassicLinkVPCSecurityGroups` | classic_link_vpc_security_groups | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | The IDs of one or more security groups for the VPC that you specified in the ClassicLinkVPCId property. |
| `EbsOptimized` | ebs_optimized | `boolean` | optional, computed, provider-chosen, replaces on change |  | Specifies whether the launch configuration is optimized for EBS I/O (true) or not (false). |
| `IamInstanceProfile` | iam_instance_profile | `string` | optional, computed, provider-chosen, replaces on change |  | Provides the name or the Amazon Resource Name (ARN) of the instance profile associated with the IAM role for the instance. The instance profile contains the IAM role. |
| `ImageId` | image_id | `string` | required, replaces on change |  | Provides the unique ID of the Amazon Machine Image (AMI) that was assigned during registration. |
| `InstanceId` | instance_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of the Amazon EC2 instance you want to use to create the launch configuration. |
| `InstanceMonitoring` | instance_monitoring | `boolean` | optional, computed, provider-chosen, replaces on change |  | Controls whether instances in this group are launched with detailed (true) or basic (false) monitoring. |
| `InstanceType` | instance_type | `string` | required, replaces on change |  | Specifies the instance type of the EC2 instance. |
| `KernelId` | kernel_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Provides the ID of the kernel associated with the EC2 AMI. |
| `KeyName` | key_name | `string` | optional, computed, provider-chosen, replaces on change |  | Provides the name of the EC2 key pair. |
| `LaunchConfigurationName` | launch_configuration_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the launch configuration. This name must be unique per Region per account. |
| `MetadataOptions` | metadata_options | `map` | optional, computed, provider-chosen, replaces on change |  | MetadataOptions is a property of AWS::AutoScaling::LaunchConfiguration that describes metadata options for the instances. |
| `PlacementTenancy` | placement_tenancy | `string` | optional, computed, provider-chosen, replaces on change |  | The tenancy of the instance, either default or dedicated. |
| `RamDiskId` | ram_disk_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of the RAM disk to select. |
| `SecurityGroups` | security_groups | `list` | optional, computed, provider-chosen, replaces on change |  | A list that contains the security groups to assign to the instances in the Auto Scaling group. |
| `SpotPrice` | spot_price | `string` | optional, computed, provider-chosen, replaces on change |  | The maximum hourly price you are willing to pay for any Spot Instances launched to fulfill the request. |
| `UserData` | user_data | `string` | optional, computed, provider-chosen, replaces on change |  | The Base64-encoded user data to make available to the launched EC2 instances. |

Supports update: no

Discovery: supported
