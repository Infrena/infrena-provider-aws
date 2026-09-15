# aws.networkinterfaceattachment

**CloudFormation type:** `AWS::EC2::NetworkInterfaceAttachment`

Attaches an elastic network interface (ENI) to an Amazon EC2 instance. You can use this resource type to attach additional network interfaces to an instance without interruption.

Region attribute: `region`

**Import ID:** `<region>/AttachmentId` (AWS::EC2::NetworkInterfaceAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachmentId` | attachment_id | `string` | computed |  |  |
| `DeleteOnTermination` | delete_on_termination | `boolean` | optional, computed, provider-chosen |  | Whether to delete the network interface when the instance terminates. By default, this value is set to ``true``. |
| `DeviceIndex` | device_index | `string` | required, replaces on change |  | The network interface's position in the attachment order. For example, the first attached network interface has a ``DeviceIndex`` of 0. |
| `EnaQueueCount` | ena_queue_count | `integer` | optional, computed, provider-chosen |  | The number of ENA queues created with the instance. |
| `EnaSrdSpecification` | ena_srd_specification | `map` | optional, computed, provider-chosen |  | ENA Express uses AWS Scalable Reliable Datagram (SRD) technology to increase the maximum bandwidth used per stream and minimize tail latency of network traffic between EC2 instances. With ENA Express, you can communicate between two EC2 instances in the same subnet within the same account, or in different accounts. Both sending and receiving instances must have ENA Express enabled. |
| `InstanceId` | instance_id | `string` | required, replaces on change | aws.ec2.instance.InstanceId | The ID of the instance to which you will attach the ENI. |
| `NetworkInterfaceId` | network_interface_id | `string` | required, replaces on change | aws.networkinterface.Id | The ID of the ENI that you want to attach. |

Supports update: yes

Discovery: supported
