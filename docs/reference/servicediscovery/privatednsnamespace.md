# aws.privatednsnamespace

**CloudFormation type:** `AWS::ServiceDiscovery::PrivateDnsNamespace`

Resource Type definition for AWS::ServiceDiscovery::PrivateDnsNamespace

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceDiscovery::PrivateDnsNamespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the private namespace. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the namespace. |
| `HostedZoneId` | hosted_zone_id | `string` | computed |  | The ID for the Route 53 hosted zone that AWS Cloud Map creates when you create a namespace. |
| `Id` |  | `string` | computed |  | The ID of the private namespace. |
| `Name` |  | `string` | required, replaces on change |  | The name that you want to assign to this namespace. When you create a private DNS namespace, AWS Cloud Map automatically creates an Amazon Route 53 private hosted zone that has the same name as the namespace. |
| `Properties` |  | `map` | optional, computed, provider-chosen |  | Properties of the private DNS namespace. |
| `Tags` |  | `map` | tags map |  | The tags for the namespace. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters. |
| `Vpc` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of the Amazon VPC that you want to associate the namespace with. |

Supports update: yes

Discovery: supported
