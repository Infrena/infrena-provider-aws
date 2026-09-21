# aws.publicdnsnamespace

**CloudFormation type:** `AWS::ServiceDiscovery::PublicDnsNamespace`

Resource Type definition for AWS::ServiceDiscovery::PublicDnsNamespace

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceDiscovery::PublicDnsNamespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the public namespace. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the namespace. |
| `HostedZoneId` | hosted_zone_id | `string` | computed |  | The ID for the Route 53 hosted zone that AWS Cloud Map creates when you create a namespace. |
| `Id` |  | `string` | computed |  | The ID of the public namespace. |
| `Name` |  | `string` | required, replaces on change |  | The name that you want to assign to this namespace. |
| `Properties` |  | `map` | optional, computed, provider-chosen |  | Properties for the public DNS namespace. |
| `Tags` |  | `map` | tags map |  | The tags for the namespace. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters. |

Supports update: yes

Discovery: supported
