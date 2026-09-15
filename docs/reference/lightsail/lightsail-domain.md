# aws.lightsail.domain

**CloudFormation type:** `AWS::Lightsail::Domain`

Resource Type definition for AWS::Lightsail::Domain

Region attribute: `region`

**Import ID:** `<region>/DomainName` (AWS::Lightsail::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the domain (read-only). |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the domain was created (read-only). |
| `DomainEntries` | domain_entries | `list` | optional, computed, provider-chosen |  | An array of key-value pairs containing information about the domain entries. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The name of the domain to manage in Lightsail. |
| `Location` |  | `map` | computed |  | The AWS Region and Availability Zone where the domain was created (read-only). |
| `ResourceType` | resource_type | `string` | computed |  | The Lightsail resource type (read-only). |
| `SupportCode` | support_code | `string` | computed |  | The support code. Include this code in your email to support when you have questions (read-only). |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
