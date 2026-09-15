# aws.volumeattachment

**CloudFormation type:** `AWS::EC2::VolumeAttachment`

Attaches an Amazon EBS volume to a running instance and exposes it to the instance with the specified device name.

Region attribute: `region`

**Import ID:** `<region>/VolumeId|InstanceId` (AWS::EC2::VolumeAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Device` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The device name |
| `EbsCardIndex` | ebs_card_index | `integer` | optional, computed, provider-chosen, replaces on change |  | The index of the EBS card. Some instance types support multiple EBS cards. The default EBS card index is 0. |
| `InstanceId` | instance_id | `string` | required, replaces on change | aws.ec2.instance.InstanceId | The ID of the instance to which the volume attaches |
| `VolumeId` | volume_id | `string` | required, replaces on change | aws.ec2.volume.VolumeId | The ID of the Amazon EBS volume |

Supports update: no

Discovery: supported
