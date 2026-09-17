# aws.domainconfiguration

**CloudFormation type:** `AWS::IoT::DomainConfiguration`

Create and manage a Domain Configuration

Region attribute: `region`

**Import ID:** `<region>/DomainConfigurationName` (AWS::IoT::DomainConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationProtocol` | application_protocol | `string` | optional, computed, provider-chosen |  |  |
| `Arn` |  | `string` | computed |  |  |
| `AuthenticationType` | authentication_type | `string` | optional, computed, provider-chosen |  |  |
| `AuthorizerConfig` | authorizer_config | `map` | optional, computed, provider-chosen |  |  |
| `ClientCertificateConfig` | client_certificate_config | `map` | optional, computed, provider-chosen |  |  |
| `DomainConfigurationName` | domain_configuration_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainConfigurationStatus` | domain_configuration_status | `string` | optional, computed, provider-chosen |  |  |
| `DomainName` | domain_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DomainType` | domain_type | `string` | computed |  |  |
| `ServerCertificateArns` | server_certificate_arns | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ServerCertificateConfig` | server_certificate_config | `map` | optional, computed, provider-chosen |  |  |
| `ServerCertificates` | server_certificates | `list` | computed |  |  |
| `ServiceType` | service_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TlsConfig` | tls_config | `map` | optional, computed, provider-chosen |  |  |
| `ValidationCertificateArn` | validation_certificate_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.iot.certificate.Arn |  |

Supports update: yes

Discovery: supported
