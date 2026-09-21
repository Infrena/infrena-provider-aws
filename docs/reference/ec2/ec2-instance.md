# aws.ec2.instance

**CloudFormation type:** `AWS::EC2::Instance`

Resource Type definition for AWS::EC2::Instance

Region attribute: `region`

**Import ID:** `<region>/InstanceId` (AWS::EC2::Instance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalInfo` | additional_info | `string` | optional, computed, provider-chosen, write-only |  | This property is reserved for internal use. If you use it, the stack fails with this error: Bad property set: [Testing this property] (Service: AmazonEC2; Status Code: 400; Error Code: InvalidParameterCombination; Request ID: 0XXXXXX-49c7-4b40-8bcc-76885dcXXXXX). |
| `Affinity` |  | `string` | optional, computed, provider-chosen |  | Indicates whether the instance is associated with a dedicated host. If you want the instance to always restart on the same host on which it was launched, specify host. If you want the instance to restart on any available host, but try to launch onto the last host it ran on (on a best-effort basis), specify default. |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The Availability Zone of the instance. |
| `BlockDeviceMappings` | block_device_mappings | `list` | optional, computed, provider-chosen |  | The block device mapping entries that defines the block devices to attach to the instance at launch. |
| `CpuOptions` | cpu_options | `map` | optional, computed, provider-chosen, replaces on change |  | The CPU options for the instance. |
| `CreditSpecification` | credit_specification | `map` | optional, computed, provider-chosen |  | The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited. |
| `DisableApiTermination` | disable_api_termination | `boolean` | optional, computed, provider-chosen |  | If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can. |
| `EbsOptimized` | ebs_optimized | `boolean` | optional, computed, provider-chosen |  | Indicates whether the instance is optimized for Amazon EBS I/O. |
| `ElasticGpuSpecifications` | elastic_gpu_specifications | `list` | optional, computed, provider-chosen, replaces on change |  | An elastic GPU to associate with the instance. Amazon Elastic Graphics is no longer available. |
| `ElasticInferenceAccelerators` | elastic_inference_accelerators | `list` | optional, computed, provider-chosen, replaces on change |  | An elastic inference accelerator to associate with the instance. Amazon Elastic Inference is no longer available. |
| `EnclaveOptions` | enclave_options | `map` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the instance is enabled for AWS Nitro Enclaves. |
| `HibernationOptions` | hibernation_options | `map` | optional, computed, provider-chosen, replaces on change |  | Indicates whether an instance is enabled for hibernation. |
| `HostId` | host_id | `string` | optional, computed, provider-chosen | aws.ec2.host.HostId | If you specify host for the Affinity property, the ID of a dedicated host that the instance is associated with. If you don't specify an ID, Amazon EC2 launches the instance onto any available, compatible dedicated host in your account. |
| `HostResourceGroupArn` | host_resource_group_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the host resource group in which to launch the instances. If you specify a host resource group ARN, omit the Tenancy parameter or set it to host. |
| `IamInstanceProfile` | iam_instance_profile | `string` | optional, computed, provider-chosen |  | The IAM instance profile. |
| `ImageId` | ami, image_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template. |
| `InstanceId` | instance_id | `string` | computed |  | The EC2 Instance ID. |
| `InstanceInitiatedShutdownBehavior` | instance_initiated_shutdown_behavior | `string` | optional, computed, provider-chosen |  | Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown). |
| `InstanceType` | instance_type | `string` | optional, computed, provider-chosen |  | The instance type. |
| `Ipv6AddressCount` | ipv6_address_count | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | [EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet. |
| `Ipv6Addresses` | ipv6_addresses | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | [EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface. |
| `KernelId` | kernel_id | `string` | optional, computed, provider-chosen |  | The ID of the kernel. |
| `KeyName` | key_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the key pair. |
| `LaunchTemplate` | launch_template | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | The launch template to use to launch the instances. |
| `LicenseSpecifications` | license_specifications | `list` | optional, computed, provider-chosen, replaces on change |  | The license configurations. |
| `MetadataOptions` | metadata_options | `map` | optional, computed, provider-chosen |  | The metadata options for the instance |
| `Monitoring` |  | `boolean` | optional, computed, provider-chosen |  | Specifies whether detailed monitoring is enabled for the instance. |
| `NetworkInterfaces` | network_interfaces | `list` | optional, computed, provider-chosen, replaces on change |  | The network interfaces to associate with the instance. |
| `PlacementGroupName` | placement_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of an existing placement group that you want to launch the instance into (cluster \| partition \| spread). |
| `PrivateDnsName` | private_dns_name | `string` | computed |  | The private DNS name of the specified instance. For example: ip-10-24-34-0.ec2.internal. |
| `PrivateDnsNameOptions` | private_dns_name_options | `map` | optional, computed, provider-chosen |  | The options for the instance hostname. |
| `PrivateIp` | private_ip | `string` | computed |  | The private IP address of the specified instance. For example: 10.24.34.0. |
| `PrivateIpAddress` | private_ip_address | `string` | optional, computed, provider-chosen, replaces on change |  | [EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet. |
| `PropagateTagsToVolumeOnCreation` | propagate_tags_to_volume_on_creation | `boolean` | optional, computed, provider-chosen, write-only |  | Indicates whether to assign the tags from the instance to all of the volumes attached to the instance at launch. If you specify true and you assign tags to the instance, those tags are automatically assigned to all of the volumes that you attach to the instance at launch. If you specify false, those tags are not assigned to the attached volumes. |
| `PublicDnsName` | public_dns_name | `string` | computed |  | The public DNS name of the specified instance. For example: ec2-107-20-50-45.compute-1.amazonaws.com. |
| `PublicIp` | public_ip | `string` | computed |  | The public IP address of the specified instance. For example: 192.0.2.0. |
| `RamdiskId` | ramdisk_id | `string` | optional, computed, provider-chosen |  | The ID of the RAM disk to select. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | The IDs of the security groups. |
| `SecurityGroups` | security_groups | `list` | optional, computed, provider-chosen, replaces on change |  | the names of the security groups. For a nondefault VPC, you must use security group IDs instead. |
| `SourceDestCheck` | source_dest_check | `boolean` | optional, computed, provider-chosen |  | Specifies whether to enable an instance launched in a VPC to perform NAT. |
| `SsmAssociations` | ssm_associations | `list` | optional, computed, provider-chosen |  | The SSM document and parameter values in AWS Systems Manager to associate with this instance. |
| `State` |  | `map` | computed |  | The current state of the instance |
| `SubnetId` | subnet_id | `string` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId | [EC2-VPC] The ID of the subnet to launch the instance into. |
| `Tags` |  | `map` | tags map |  | The tags to add to the instance. |
| `Tenancy` |  | `string` | optional, computed, provider-chosen |  | The tenancy of the instance (if the instance is running in a VPC). An instance with a tenancy of dedicated runs on single-tenant hardware. |
| `UserData` | user_data | `string` | optional, computed, provider-chosen |  | The user data to make available to the instance. |
| `Volumes` |  | `list` | optional, computed, provider-chosen |  | The volumes to attach to the instance. |
| `VpcId` | vpc_id | `string` | computed |  | The ID of the VPC that the instance is running in. |

Supports update: yes

Discovery: supported
