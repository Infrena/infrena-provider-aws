# aws.disk

**CloudFormation type:** `AWS::Lightsail::Disk`

Resource Type definition for AWS::Lightsail::Disk

Region attribute: `region`

**Import ID:** `<region>/DiskName` (AWS::Lightsail::Disk)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AddOns` | add_ons | `list` | optional, computed, provider-chosen |  | An array of objects representing the add-ons to enable for the new instance. |
| `AttachedTo` | attached_to | `string` | computed |  | Name of the attached Lightsail Instance |
| `AttachmentState` | attachment_state | `string` | computed |  | Attachment State of the Lightsail disk |
| `AvailabilityZone` | availability_zone | `string` | optional, computed, provider-chosen, replaces on change |  | The Availability Zone in which to create your instance. Use the following format: us-east-2a (case sensitive). Be sure to add the include Availability Zones parameter to your request. |
| `DiskArn` | disk_arn | `string` | computed |  |  |
| `DiskName` | disk_name | `string` | required, replaces on change |  | The names to use for your new Lightsail disk. |
| `Iops` |  | `integer` | computed |  | Iops of the Lightsail disk |
| `IsAttached` | is_attached | `boolean` | computed |  | Check is Disk is attached state |
| `Location` |  | `map` | optional, computed, provider-chosen |  | Location of a resource. |
| `Path` |  | `string` | computed |  | Path of the  attached Disk |
| `ResourceType` | resource_type | `string` | computed |  | Resource type of Lightsail instance. |
| `SizeInGb` | size_in_gb | `integer` | required, replaces on change |  | Size of the Lightsail disk |
| `State` |  | `string` | computed |  | State of the Lightsail disk |
| `SupportCode` | support_code | `string` | computed |  | Support code to help identify any issues |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
