# aws.instancesnapshot

**CloudFormation type:** `AWS::Lightsail::InstanceSnapshot`

Resource Type definition for AWS::Lightsail::InstanceSnapshot

Region attribute: `region`

**Import ID:** `<region>/InstanceSnapshotName` (AWS::Lightsail::InstanceSnapshot)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the snapshot. |
| `FromInstanceArn` | from_instance_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the instance from which the snapshot was created. |
| `FromInstanceName` | from_instance_name | `string` | computed |  | The instance from which the snapshot was created. |
| `InstanceName` | instance_name | `string` | required, replaces on change |  | The instance from which the snapshot was created. |
| `InstanceSnapshotName` | instance_snapshot_name | `string` | required, replaces on change |  | The name of the snapshot. |
| `IsFromAutoSnapshot` | is_from_auto_snapshot | `boolean` | computed |  | A Boolean value indicating whether the snapshot was created from an automatic snapshot. |
| `Location` |  | `map` | computed |  | The region name and Availability Zone where you created the snapshot. |
| `ResourceType` | resource_type | `string` | computed |  | The type of resource (usually InstanceSnapshot). |
| `SizeInGb` | size_in_gb | `integer` | computed |  | The size in GB of the SSD |
| `State` |  | `string` | computed |  | The state the snapshot is in. |
| `SupportCode` | support_code | `string` | computed |  | Support code to help identify any issues |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
