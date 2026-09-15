# aws.vpcconnector

**CloudFormation type:** `AWS::AppRunner::VpcConnector`

The AWS::AppRunner::VpcConnector resource specifies an App Runner VpcConnector.

Region attribute: `region`

**Import ID:** `<region>/VpcConnectorArn` (AWS::AppRunner::VpcConnector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `SecurityGroups` | security_groups | `list` | optional, computed, provider-chosen, replaces on change |  | A list of IDs of security groups that App Runner should use for access to AWS resources under the specified subnets. If not specified, App Runner uses the default security group of the Amazon VPC. The default security group allows all outbound traffic. |
| `Subnets` |  | `list` | required, replaces on change |  | A list of IDs of subnets that App Runner should use when it associates your service with a custom Amazon VPC. Specify IDs of subnets of a single Amazon VPC. App Runner determines the Amazon VPC from the subnets you specify. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, write-only, tags map |  | A list of metadata items that you can associate with your VPC connector resource. A tag is a key-value pair. |
| `VpcConnectorArn` | vpc_connector_arn | `string` | computed |  | The Amazon Resource Name (ARN) of this VPC connector. |
| `VpcConnectorName` | vpc_connector_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the VPC connector. If you don't specify a name, AWS CloudFormation generates a name for your VPC connector. |
| `VpcConnectorRevision` | vpc_connector_revision | `integer` | computed |  | The revision of this VPC connector. It's unique among all the active connectors ("Status": "ACTIVE") that share the same Name. |

Supports update: no

Discovery: supported
