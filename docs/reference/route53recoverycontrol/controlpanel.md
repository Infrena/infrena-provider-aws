# aws.controlpanel

**CloudFormation type:** `AWS::Route53RecoveryControl::ControlPanel`

AWS Route53 Recovery Control Control Panel resource schema .

Region attribute: `region`

**Import ID:** `<region>/ControlPanelArn` (AWS::Route53RecoveryControl::ControlPanel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClusterArn` | cluster_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.route53recoverycontrol.cluster.ClusterArn | Cluster to associate with the Control Panel |
| `ControlPanelArn` | control_panel_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the cluster. |
| `DefaultControlPanel` | default_control_panel | `boolean` | computed |  | A flag that Amazon Route 53 Application Recovery Controller sets to true to designate the default control panel for a cluster. When you create a cluster, Amazon Route 53 Application Recovery Controller creates a control panel, and sets this flag for that control panel. If you create a control panel yourself, this flag is set to false. |
| `Name` |  | `string` | required |  | The name of the control panel. You can use any non-white space character in the name. |
| `RoutingControlCount` | routing_control_count | `integer` | computed |  | Count of associated routing controls |
| `Status` |  | `string` | computed |  | The deployment status of control panel. Status can be one of the following: PENDING, DEPLOYED, PENDING_DELETION. |
| `Tags` |  | `map` | replaces on change, tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
