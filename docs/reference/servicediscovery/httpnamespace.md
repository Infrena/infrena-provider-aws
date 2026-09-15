# aws.httpnamespace

**CloudFormation type:** `AWS::ServiceDiscovery::HttpNamespace`

Resource Type definition for AWS::ServiceDiscovery::HttpNamespace

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceDiscovery::HttpNamespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the namespace. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the namespace. |
| `Id` |  | `string` | computed |  | The ID of the namespace. |
| `Name` |  | `string` | required, replaces on change |  | The name that you want to assign to this namespace. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the namespace. Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters. |

Supports update: yes

Discovery: supported
