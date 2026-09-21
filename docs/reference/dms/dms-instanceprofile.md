# aws.dms.instanceprofile

**CloudFormation type:** `AWS::DMS::InstanceProfile`

Resource schema for AWS::DMS::InstanceProfile.

Region attribute: `region`

**Import ID:** `<region>/InstanceProfileArn` (AWS::DMS::InstanceProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen |  | The property describes an availability zone of the instance profile. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The optional description of the instance profile. |
| `InstanceProfileArn` | instance_profile_arn | `string` | computed |  | The property describes an ARN of the instance profile. |
| `InstanceProfileCreationTime` | instance_profile_creation_time | `string` | computed |  | The property describes a creating time of the instance profile. |
| `InstanceProfileIdentifier` | instance_profile_identifier | `string` | optional, computed, provider-chosen, write-only |  | The property describes an identifier for the instance profile. It is used for describing/deleting/modifying. Can be name/arn |
| `InstanceProfileName` | instance_profile_name | `string` | optional, computed, provider-chosen |  | The property describes a name for the instance profile. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  | The property describes kms key arn for the instance profile. |
| `NetworkType` | network_type | `string` | optional, computed, provider-chosen |  | The property describes a network type for the instance profile. |
| `PubliclyAccessible` | publicly_accessible | `boolean` | optional, computed, provider-chosen |  | The property describes the publicly accessible of the instance profile |
| `SubnetGroupIdentifier` | subnet_group_identifier | `string` | optional, computed, provider-chosen |  | The property describes a subnet group identifier for the instance profile. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `VpcSecurityGroups` | vpc_security_groups | `list` | optional, computed, provider-chosen |  | The property describes vps security groups for the instance profile. |

Supports update: yes

Discovery: supported
