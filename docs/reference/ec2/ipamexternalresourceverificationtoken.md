# aws.ipamexternalresourceverificationtoken

**CloudFormation type:** `AWS::EC2::IpamExternalResourceVerificationToken`

A verification token is an AWS-generated random value that you can use to prove ownership of an external resource. For example, you can use a verification token to validate that you control a public IP address range when you bring an IP address range to AWS (BYOIP).

Region attribute: `region`

**Import ID:** `<region>/IpamExternalResourceVerificationTokenArn` (AWS::EC2::IpamExternalResourceVerificationToken)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IpamArn` | ipam_arn | `string` | computed |  | The ARN of the IPAM that created the token. |
| `IpamExternalResourceVerificationTokenArn` | ipam_external_resource_verification_token_arn | `string` | computed |  | The ARN of the IPAM external resource verification token. |
| `IpamExternalResourceVerificationTokenId` | ipam_external_resource_verification_token_id | `string` | computed |  | The ID of the token. |
| `IpamId` | ipam_id | `string` | required, replaces on change | aws.ipam.IpamId | The ID of the IPAM that will create the token. |
| `IpamRegion` | ipam_region | `string` | computed |  | The Region of the IPAM that created the token. |
| `NotAfter` | not_after | `string` | computed |  | The token expiration. |
| `State` |  | `string` | computed |  | The token state. |
| `Status` |  | `string` | computed |  | The token status. |
| `Tags` |  | `map` | tags map |  | The tags for the token. |
| `TokenName` | token_name | `string` | computed |  | The token name. |
| `TokenValue` | token_value | `string` | computed |  | The token value. |

Supports update: yes

Discovery: supported
