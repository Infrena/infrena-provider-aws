# aws.packaginggroup

**CloudFormation type:** `AWS::MediaPackage::PackagingGroup`

Resource schema for AWS::MediaPackage::PackagingGroup

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaPackage::PackagingGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the PackagingGroup. |
| `Authorization` |  | `map` | optional, computed, provider-chosen |  | CDN Authorization |
| `DomainName` | domain_name | `string` | computed |  | The fully qualified domain name for Assets in the PackagingGroup. |
| `EgressAccessLogs` | egress_access_logs | `map` | optional, computed, provider-chosen |  | The configuration parameters for egress access logging. |
| `Id` |  | `string` | required, replaces on change |  | The ID of the PackagingGroup. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
