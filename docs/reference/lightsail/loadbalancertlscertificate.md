# aws.loadbalancertlscertificate

**CloudFormation type:** `AWS::Lightsail::LoadBalancerTlsCertificate`

Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate

Region attribute: `region`

**Import ID:** `<region>/CertificateName|LoadBalancerName` (AWS::Lightsail::LoadBalancerTlsCertificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateAlternativeNames` | certificate_alternative_names | `list` | optional, computed, provider-chosen, replaces on change |  | An array of strings listing alternative domains and subdomains for your SSL/TLS certificate. |
| `CertificateDomainName` | certificate_domain_name | `string` | required, replaces on change |  | The domain name (e.g., example.com ) for your SSL/TLS certificate. |
| `CertificateName` | certificate_name | `string` | required, replaces on change |  | The SSL/TLS certificate name. |
| `HttpsRedirectionEnabled` | https_redirection_enabled | `boolean` | optional, computed, provider-chosen |  | A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer. |
| `IsAttached` | is_attached | `boolean` | optional, computed, provider-chosen |  | When true, the SSL/TLS certificate is attached to the Lightsail load balancer. |
| `LoadBalancerName` | load_balancer_name | `string` | required, replaces on change |  | The name of your load balancer. |
| `LoadBalancerTlsCertificateArn` | load_balancer_tls_certificate_arn | `string` | computed |  |  |
| `Status` |  | `string` | computed |  | The validation status of the SSL/TLS certificate. |

Supports update: yes

Discovery: supported
