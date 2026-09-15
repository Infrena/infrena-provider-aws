# aws.vpcblockpublicaccessoptions

**CloudFormation type:** `AWS::EC2::VPCBlockPublicAccessOptions`

Resource Type definition for AWS::EC2::VPCBlockPublicAccessOptions

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::EC2::VPCBlockPublicAccessOptions)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | The identifier for the specified AWS account. |
| `ExclusionsAllowed` | exclusions_allowed | `string` | computed |  | Determines if exclusions are allowed. If you have enabled VPC BPA at the Organization level, exclusions may be not-allowed. Otherwise, they are allowed. |
| `InternetGatewayBlockMode` | internet_gateway_block_mode | `string` | required |  | The desired Block Public Access mode for Internet Gateways in your account. We do not allow to create in a off mode as this is the default value |

Supports update: yes

Discovery: not supported
