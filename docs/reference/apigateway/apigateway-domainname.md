# aws.apigateway.domainname

**CloudFormation type:** `AWS::ApiGateway::DomainName`

The ``AWS::ApiGateway::DomainName`` resource specifies a public custom domain name for your API in API Gateway.

Region attribute: `region`

**Import ID:** `<region>/DomainName` (AWS::ApiGateway::DomainName)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateArn` | certificate_arn | `string` | optional, computed, provider-chosen |  |  |
| `DistributionDomainName` | distribution_domain_name | `string` | computed |  |  |
| `DistributionHostedZoneId` | distribution_hosted_zone_id | `string` | computed |  |  |
| `DomainName` | domain_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainNameArn` | domain_name_arn | `string` | computed |  |  |
| `EndpointAccessMode` | endpoint_access_mode | `string` | optional, computed, provider-chosen |  |  |
| `EndpointConfiguration` | endpoint_configuration | `map` | optional, computed, provider-chosen |  | The ``EndpointConfiguration`` property type specifies the endpoint types and IP address types of an Amazon API Gateway domain name. |
| `MutualTlsAuthentication` | mutual_tls_authentication | `map` | optional, computed, provider-chosen |  |  |
| `OwnershipVerificationCertificateArn` | ownership_verification_certificate_arn | `string` | optional, computed, provider-chosen |  | The ARN of the public certificate issued by ACM to validate ownership of your custom domain. Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the RegionalCertificateArn. |
| `RegionalCertificateArn` | regional_certificate_arn | `string` | optional, computed, provider-chosen |  |  |
| `RegionalDomainName` | regional_domain_name | `string` | computed |  |  |
| `RegionalHostedZoneId` | regional_hosted_zone_id | `string` | computed |  |  |
| `RoutingMode` | routing_mode | `string` | optional, computed, provider-chosen |  |  |
| `SecurityPolicy` | security_policy | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
