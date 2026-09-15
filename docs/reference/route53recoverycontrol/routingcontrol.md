# aws.routingcontrol

**CloudFormation type:** `AWS::Route53RecoveryControl::RoutingControl`

AWS Route53 Recovery Control Routing Control resource schema .

Region attribute: `region`

**Import ID:** `<region>/RoutingControlArn` (AWS::Route53RecoveryControl::RoutingControl)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusterArn` | cluster_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.route53recoverycontrol.cluster.ClusterArn | Arn associated with Control Panel |
| `ControlPanelArn` | control_panel_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.controlpanel.ControlPanelArn | The Amazon Resource Name (ARN) of the control panel. |
| `Name` |  | `string` | required |  | The name of the routing control. You can use any non-white space character in the name. |
| `RoutingControlArn` | routing_control_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the routing control. |
| `Status` |  | `string` | computed |  | The deployment status of the routing control. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION. |

Supports update: yes

Discovery: supported (parent resource required)
