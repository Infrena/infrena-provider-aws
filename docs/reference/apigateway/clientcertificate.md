# aws.clientcertificate

**CloudFormation type:** `AWS::ApiGateway::ClientCertificate`

The ``AWS::ApiGateway::ClientCertificate`` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.

Region attribute: `region`

**Import ID:** `<region>/ClientCertificateId` (AWS::ApiGateway::ClientCertificate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClientCertificateId` | client_certificate_id | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
