# aws.ec2.host

**CloudFormation type:** `AWS::EC2::Host`

Resource Type definition for AWS::EC2::Host

Region attribute: `region`

**Import ID:** `<region>/HostId` (AWS::EC2::Host)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssetId` | asset_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the Outpost hardware asset. |
| `AutoPlacement` | auto_placement | `string` | optional, computed, provider-chosen |  | Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. |
| `AvailabilityZone` | availability_zone | `string` | required, replaces on change |  | The Availability Zone in which to allocate the Dedicated Host. |
| `HostId` | host_id | `string` | computed |  | ID of the host created. |
| `HostMaintenance` | host_maintenance | `string` | optional, computed, provider-chosen |  | Automatically allocates a new dedicated host and moves your instances on to it if a degradation is detected on your current host. |
| `HostRecovery` | host_recovery | `string` | optional, computed, provider-chosen |  | Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default. |
| `InstanceFamily` | instance_family | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. |
| `InstanceType` | instance_type | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only. |
| `OutpostArn` | outpost_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the Host. |

Supports update: yes

Discovery: supported
