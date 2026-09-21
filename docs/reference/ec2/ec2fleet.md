# aws.ec2fleet

**CloudFormation type:** `AWS::EC2::EC2Fleet`

Resource Type definition for AWS::EC2::EC2Fleet

Region attribute: `region`

**Import ID:** `<region>/FleetId` (AWS::EC2::EC2Fleet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Context` |  | `string` | optional, computed, provider-chosen |  |  |
| `ExcessCapacityTerminationPolicy` | excess_capacity_termination_policy | `string` | optional, computed, provider-chosen |  |  |
| `FleetId` | fleet_id | `string` | computed |  |  |
| `LaunchTemplateConfigs` | launch_template_configs | `list` | required, replaces on change |  |  |
| `OnDemandOptions` | on_demand_options | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `ReplaceUnhealthyInstances` | replace_unhealthy_instances | `boolean` | optional, computed, provider-chosen, replaces on change |  |  |
| `ReservedCapacityOptions` | reserved_capacity_options | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SpotOptions` | spot_options | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `TagSpecifications` | tag_specifications | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `TargetCapacitySpecification` | target_capacity_specification | `map` | required |  |  |
| `TerminateInstancesWithExpiration` | terminate_instances_with_expiration | `boolean` | optional, computed, provider-chosen, replaces on change |  |  |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ValidFrom` | valid_from | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ValidUntil` | valid_until | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
