# aws.certificateprovider

**CloudFormation type:** `AWS::IoT::CertificateProvider`

Use the AWS::IoT::CertificateProvider resource to declare an AWS IoT Certificate Provider.

Region attribute: `region`

**Import ID:** `<region>/CertificateProviderName` (AWS::IoT::CertificateProvider)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountDefaultForOperations` | account_default_for_operations | `list` | required |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CertificateProviderName` | certificate_provider_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `LambdaFunctionArn` | lambda_function_arn | `string` | required |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
