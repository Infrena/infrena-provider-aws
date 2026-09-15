# aws.cloudfront.distribution

**CloudFormation type:** `AWS::CloudFront::Distribution`

A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::Distribution)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DistributionConfig` | distribution_config | `map` | required |  | A distribution configuration. |
| `DomainName` | domain_name | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A complex type that contains zero or more ``Tag`` elements. |

Supports update: yes

Discovery: supported
