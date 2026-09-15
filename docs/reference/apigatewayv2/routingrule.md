# aws.routingrule

**CloudFormation type:** `AWS::ApiGatewayV2::RoutingRule`

Represents a routing rule. When the incoming request to a domain name matches the conditions for a rule, API Gateway invokes a stage of a target API. Supported only for REST APIs.

Region attribute: `region`

**Import ID:** `<region>/RoutingRuleArn` (AWS::ApiGatewayV2::RoutingRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  | The resulting action based on matching a routing rules condition. Only InvokeApi is supported. |
| `Conditions` |  | `list` | required |  | The conditions of the routing rule. |
| `DomainNameArn` | domain_name_arn | `string` | required, replaces on change | aws.apigatewayv2.domainname.DomainNameArn | The ARN of the domain name. |
| `Priority` |  | `integer` | required |  | The order in which API Gateway evaluates a rule. Priority is evaluated from the lowest value to the highest value. Rules can't have the same priority. Priority values 1-1,000,000 are supported. |
| `RoutingRuleArn` | routing_rule_arn | `string` | computed |  |  |
| `RoutingRuleId` | routing_rule_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
