# aws.warmpool

**CloudFormation type:** `AWS::AutoScaling::WarmPool`

Resource schema for AWS::AutoScaling::WarmPool.

Region attribute: `region`

**Import ID:** `<region>/AutoScalingGroupName` (AWS::AutoScaling::WarmPool)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoScalingGroupName` | auto_scaling_group_name | `string` | required, replaces on change |  |  |
| `InstanceReusePolicy` | instance_reuse_policy | `map` | optional, computed, provider-chosen |  |  |
| `MaxGroupPreparedCapacity` | max_group_prepared_capacity | `integer` | optional, computed, provider-chosen |  |  |
| `MinSize` | min_size | `integer` | optional, computed, provider-chosen |  |  |
| `PoolState` | pool_state | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
