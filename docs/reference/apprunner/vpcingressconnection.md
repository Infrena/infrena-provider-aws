# aws.vpcingressconnection

**CloudFormation type:** `AWS::AppRunner::VpcIngressConnection`

The AWS::AppRunner::VpcIngressConnection resource is an App Runner resource that specifies an App Runner VpcIngressConnection.

Region attribute: `region`

**Import ID:** `<region>/VpcIngressConnectionArn` (AWS::AppRunner::VpcIngressConnection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DomainName` | domain_name | `string` | computed |  | The Domain name associated with the VPC Ingress Connection. |
| `IngressVpcConfiguration` | ingress_vpc_configuration | `map` | required |  | The configuration of customer’s VPC and related VPC endpoint |
| `ServiceArn` | service_arn | `string` | required, replaces on change | aws.apprunner.service.ServiceArn | The Amazon Resource Name (ARN) of the service. |
| `Status` |  | `string` | computed |  | The current status of the VpcIngressConnection. |
| `Tags` |  | `map` | replaces on change, write-only, tags map |  |  |
| `VpcIngressConnectionArn` | vpc_ingress_connection_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the VpcIngressConnection. |
| `VpcIngressConnectionName` | vpc_ingress_connection_name | `string` | optional, computed, provider-chosen, replaces on change |  | The customer-provided Vpc Ingress Connection name. |

Supports update: yes

Discovery: supported
