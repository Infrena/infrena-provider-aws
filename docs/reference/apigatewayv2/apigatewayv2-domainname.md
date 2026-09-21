# aws.apigatewayv2.domainname

**CloudFormation type:** `AWS::ApiGatewayV2::DomainName`

The ``AWS::ApiGatewayV2::DomainName`` resource specifies a custom domain name for your API in Amazon API Gateway (API Gateway). 

Region attribute: `region`

**Import ID:** `<region>/DomainName` (AWS::ApiGatewayV2::DomainName)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DomainName` | domain_name | `string` | required, replaces on change |  | The custom domain name for your API in Amazon API Gateway. Uppercase letters and the underscore (``_``) character are not supported. |
| `DomainNameArn` | domain_name_arn | `string` | computed |  |  |
| `DomainNameConfigurations` | domain_name_configurations | `list` | optional, computed, provider-chosen |  | The domain name configurations. |
| `MutualTlsAuthentication` | mutual_tls_authentication | `map` | optional, computed, provider-chosen |  | If specified, API Gateway performs two-way authentication between the client and the server. Clients must present a trusted certificate to access your API. |
| `RegionalDomainName` | regional_domain_name | `string` | computed |  |  |
| `RegionalHostedZoneId` | regional_hosted_zone_id | `string` | computed |  |  |
| `RoutingMode` | routing_mode | `string` | optional, computed, provider-chosen |  | The routing mode API Gateway uses to route traffic to your APIs. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The collection of tags associated with a domain name. |

Supports update: yes

Discovery: supported
