# aws.globalaccelerator.listener

**CloudFormation type:** `AWS::GlobalAccelerator::Listener`

Resource Type definition for AWS::GlobalAccelerator::Listener

Region attribute: `region`

**Import ID:** `<region>/ListenerArn` (AWS::GlobalAccelerator::Listener)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceleratorArn` | accelerator_arn | `string` | required, replaces on change | aws.accelerator.AcceleratorArn | The Amazon Resource Name (ARN) of the accelerator. |
| `ClientAffinity` | client_affinity | `string` | optional, computed, provider-chosen |  | Client affinity lets you direct all requests from a user to the same endpoint. |
| `ListenerArn` | listener_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the listener. |
| `PortRanges` | port_ranges | `list` | required |  |  |
| `Protocol` |  | `string` | required |  | The protocol for the listener. |

Supports update: yes

Discovery: supported
