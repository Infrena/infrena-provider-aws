# aws.cloudfront.truststore

**CloudFormation type:** `AWS::CloudFront::TrustStore`

A trust store.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::TrustStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CaCertificatesBundleSource` | ca_certificates_bundle_source | `map` | optional, computed, provider-chosen, write-only |  | A CA certificates bundle source. |
| `ETag` | e_tag | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  | The trust store's name. |
| `NumberOfCaCertificates` | number_of_ca_certificates | `integer` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | A complex type that contains zero or more ``Tag`` elements. |
| `UseClientCertificateOCSPEndpoint` | use_client_certificate_ocsp_endpoint | `boolean` | optional, computed, provider-chosen |  | A boolean. When true, performs real-time certificate revocation checks by querying the OCSP endpoint specified within the client certificate. |

Supports update: yes

Discovery: supported
