# aws.workforce

**CloudFormation type:** `AWS::SageMaker::Workforce`

Resource Type definition for AWS::SageMaker::Workforce. Use to create a private workforce that you can use to label your training data using Amazon SageMaker Ground Truth.

Region attribute: `region`

**Import ID:** `<region>/WorkforceArn` (AWS::SageMaker::Workforce)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CognitoConfig` | cognito_config | `map` | optional, computed, provider-chosen, replaces on change |  | The configuration of an Amazon Cognito workforce. A single Cognito workforce is created using and corresponds to a single Amazon Cognito user pool. |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  | The IP address type for the workforce. IPv4 only or dualstack (IPv4 and IPv6). |
| `OidcConfig` | oidc_config | `map` | optional, computed, provider-chosen |  | The configuration of an OIDC Identity Provider (IdP) private workforce. |
| `SourceIpConfig` | source_ip_config | `map` | optional, computed, provider-chosen |  | A list of IP address ranges used to access your training data. |
| `SubDomain` | sub_domain | `string` | computed |  | The subdomain for your OIDC Identity Provider. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs. |
| `WorkforceArn` | workforce_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the private workforce. |
| `WorkforceName` | workforce_name | `string` | required, replaces on change |  | The name of the private workforce. |
| `WorkforceVpcConfig` | workforce_vpc_config | `map` | optional, computed, provider-chosen |  | The VPC configuration for the workforce. |

Supports update: yes

Discovery: supported
