# aws.networkconnector

**CloudFormation type:** `AWS::Lambda::NetworkConnector`

Resource Type definition for AWS::Lambda::NetworkConnector

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Lambda::NetworkConnector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the network connector. |
| `Configuration` |  | `map` | required |  | The network configuration for the connector. Specify a VpcEgressConfiguration to enable outbound traffic routing through your VPC. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A unique name for the network connector within your account and Region. Must be 1 to 64 alphanumeric characters, hyphens, or underscores. |
| `OperatorRole` | operator_role | `string` | optional, computed, provider-chosen |  | The ARN of the IAM role that Lambda assumes to manage elastic network interfaces in your VPC. This role must have permissions for ec2:CreateNetworkInterface and related describe operations. |
| `State` |  | `string` | computed |  | The current state of the network connector. |
| `Tags` |  | `map` | tags map |  | A list of tags to apply to the network connector. Use tags to categorize network connectors for cost allocation, access control, or operational management. |

Supports update: yes

Discovery: supported
