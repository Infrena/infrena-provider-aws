# aws.vpclattice.service

**CloudFormation type:** `AWS::VpcLattice::Service`

A service is any software application that can run on instances containers, or serverless functions within an account or virtual private cloud (VPC).

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::Service)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AuthType` | auth_type | `string` | optional, computed, provider-chosen |  |  |
| `CertificateArn` | certificate_arn | `string` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `CustomDomainName` | custom_domain_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DnsEntry` | dns_entry | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `IdleTimeoutSeconds` | idle_timeout_seconds | `integer` | optional, computed, provider-chosen |  |  |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
