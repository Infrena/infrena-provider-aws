# aws.gameservergroup

**CloudFormation type:** `AWS::GameLift::GameServerGroup`

The AWS::GameLift::GameServerGroup resource creates an Amazon GameLift (GameLift) GameServerGroup.

Region attribute: `region`

**Import ID:** `<region>/GameServerGroupArn` (AWS::GameLift::GameServerGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScalingGroupArn` | auto_scaling_group_arn | `string` | computed |  | A generated unique ID for the EC2 Auto Scaling group that is associated with this game server group. |
| `AutoScalingPolicy` | auto_scaling_policy | `map` | optional, computed, provider-chosen, write-only |  | Configuration settings to define a scaling policy for the Auto Scaling group that is optimized for game hosting. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource. |
| `BalancingStrategy` | balancing_strategy | `string` | optional, computed, provider-chosen |  | The fallback balancing method to use for the game server group when Spot Instances in a Region become unavailable or are not viable for game hosting. |
| `DeleteOption` | delete_option | `string` | optional, computed, provider-chosen, write-only |  | The type of delete to perform. |
| `GameServerGroupArn` | game_server_group_arn | `string` | computed |  | A generated unique ID for the game server group. |
| `GameServerGroupName` | game_server_group_name | `string` | required |  | An identifier for the new game server group. |
| `GameServerProtectionPolicy` | game_server_protection_policy | `string` | optional, computed, provider-chosen |  | A flag that indicates whether instances in the game server group are protected from early termination. |
| `InstanceDefinitions` | instance_definitions | `list` | required |  | A set of EC2 instance types to use when creating instances in the group. |
| `LaunchTemplate` | launch_template | `map` | optional, computed, provider-chosen, write-only |  | The EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource. |
| `MaxSize` | max_size | `float` | optional, computed, provider-chosen, write-only |  | The maximum number of instances allowed in the EC2 Auto Scaling group. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource. |
| `MinSize` | min_size | `float` | optional, computed, provider-chosen, write-only |  | The minimum number of instances allowed in the EC2 Auto Scaling group. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) for an IAM role that allows Amazon GameLift to access your EC2 Auto Scaling groups. |
| `Tags` |  | `map` | tags map |  | A list of labels to assign to the new game server group resource. |
| `VpcSubnets` | vpc_subnets | `list` | optional, computed, provider-chosen, write-only |  | A list of virtual private cloud (VPC) subnets to use with instances in the game server group. Updating this game server group property will not take effect for the created EC2 Auto Scaling group, please update the EC2 Auto Scaling group directly after creating the resource. |

Supports update: yes

Discovery: supported
