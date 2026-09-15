# aws.acmeexternalaccountbinding

**CloudFormation type:** `AWS::CertificateManager::AcmeExternalAccountBinding`

Resource Type definition for AWS::CertificateManager::AcmeExternalAccountBinding

Region attribute: `region`

**Import ID:** `<region>/AcmeExternalAccountBindingArn` (AWS::CertificateManager::AcmeExternalAccountBinding)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcmeEndpointArn` | acme_endpoint_arn | `string` | required, replaces on change | aws.acmeendpoint.AcmeEndpointArn | The ARN of the ACME endpoint this binding is associated with. |
| `AcmeExternalAccountBindingArn` | acme_external_account_binding_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the external account binding. |
| `Expiration` |  | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | The expiration configuration for the external account binding. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The IAM role ARN for cross-account access. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with the external account binding. |

Supports update: yes

Discovery: supported (parent resource required)
