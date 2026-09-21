# aws.packagegroup

**CloudFormation type:** `AWS::CodeArtifact::PackageGroup`

The resource schema to create a CodeArtifact package group.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CodeArtifact::PackageGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the package group. |
| `ContactInfo` | contact_info | `string` | optional, computed, provider-chosen |  | The contact info of the package group. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The text description of the package group. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The name of the domain that contains the package group. |
| `DomainOwner` | domain_owner | `string` | optional, computed, provider-chosen |  | The 12-digit account ID of the AWS account that owns the domain. |
| `OriginConfiguration` | origin_configuration | `map` | optional, computed, provider-chosen |  | The package origin configuration of the package group. |
| `Pattern` |  | `string` | required, replaces on change |  | The package group pattern that is used to gather packages. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to the package group. |

Supports update: yes

Discovery: supported (parent resource required)
