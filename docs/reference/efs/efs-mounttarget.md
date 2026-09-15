# aws.efs.mounttarget

**CloudFormation type:** `AWS::EFS::MountTarget`

The ``AWS::EFS::MountTarget`` resource is an Amazon EFS resource that creates a mount target for an EFS file system. You can then mount the file system on Amazon EC2 instances or other resources by using the mount target.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EFS::MountTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FileSystemId` | file_system_id | `string` | required, replaces on change | aws.efs.filesystem.FileSystemId | The ID of the file system for which to create the mount target. |
| `Id` |  | `string` | computed |  |  |
| `IpAddress` | ip_address | `string` | optional, computed, provider-chosen, replaces on change |  | If the ``IpAddressType`` for the mount target is IPv4 ( ``IPV4_ONLY`` or ``DUAL_STACK``), then specify the IPv4 address to use. If you do not specify an ``IpAddress``, then Amazon EFS selects an unused IP address from the subnet specified for ``SubnetId``. |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The IP address type for the mount target. The possible values are ``IPV4_ONLY`` (only IPv4 addresses), ``IPV6_ONLY`` (only IPv6 addresses), and ``DUAL_STACK`` (dual-stack, both IPv4 and IPv6 addresses). If you don’t specify an ``IpAddressType``, then ``IPV4_ONLY`` is used. |
| `Ipv6Address` | ipv6_address | `string` | optional, computed, provider-chosen, replaces on change |  | If the ``IPAddressType`` for the mount target is IPv6 (``IPV6_ONLY`` or ``DUAL_STACK``), then specify the IPv6 address to use. If you do not specify an ``Ipv6Address``, then Amazon EFS selects an unused IP address from the subnet specified for ``SubnetId``. |
| `SecurityGroups` | security_groups | `list` | required |  | VPC security group IDs, of the form ``sg-xxxxxxxx``. These must be for the same VPC as the subnet specified. The maximum number of security groups depends on account quota. For more information, see [Amazon VPC Quotas](https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html) in the *Amazon VPC User Guide* (see the *Security Groups* table). If you don't specify a security group, then Amazon EFS uses the default security group for the subnet's VPC. |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId | The ID of the subnet to add the mount target in. For One Zone file systems, use the subnet that is associated with the file system's Availability Zone. The subnet type must be the same type as the ``IpAddressType``. |

Supports update: yes

Discovery: supported (parent resource required)
