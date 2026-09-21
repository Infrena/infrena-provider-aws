# aws.s3files.mounttarget

**CloudFormation type:** `AWS::S3Files::MountTarget`

Resource Type definition for AWS::S3Files::MountTarget

Region attribute: `region`

**Import ID:** `<region>/MountTargetId` (AWS::S3Files::MountTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZoneId` | availability_zone_id | `string` | computed |  |  |
| `FileSystemId` | file_system_id | `string` | required, replaces on change | aws.s3files.filesystem.FileSystemId |  |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Ipv4Address` | ipv4_address | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Ipv6Address` | ipv6_address | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `MountTargetId` | mount_target_id | `string` | computed |  |  |
| `NetworkInterfaceId` | network_interface_id | `string` | computed |  |  |
| `OwnerId` | owner_id | `string` | computed |  |  |
| `SecurityGroups` | security_groups | `list` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `SubnetId` | subnet_id | `string` | required, replaces on change | aws.subnet.SubnetId |  |
| `VpcId` | vpc_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
