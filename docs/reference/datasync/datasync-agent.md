# aws.datasync.agent

**CloudFormation type:** `AWS::DataSync::Agent`

Resource schema for AWS::DataSync::Agent.

Region attribute: `region`

**Import ID:** `<region>/AgentArn` (AWS::DataSync::Agent)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActivationKey` | activation_key | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Activation key of the Agent. |
| `AgentArn` | agent_arn | `string` | computed |  | The DataSync Agent ARN. |
| `AgentName` | agent_name | `string` | optional, computed, provider-chosen |  | The name configured for the agent. Text reference used to identify the agent in the console. |
| `EndpointType` | endpoint_type | `string` | computed |  | The service endpoints that the agent will connect to. |
| `SecurityGroupArns` | security_group_arns | `list` | optional, computed, provider-chosen, replaces on change |  | The ARNs of the security group used to protect your data transfer task subnets. |
| `SubnetArns` | subnet_arns | `list` | optional, computed, provider-chosen, replaces on change |  | The ARNs of the subnets in which DataSync will create elastic network interfaces for each data transfer task. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `VpcEndpointId` | vpc_endpoint_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the VPC endpoint that the agent has access to. |

Supports update: yes

Discovery: supported
