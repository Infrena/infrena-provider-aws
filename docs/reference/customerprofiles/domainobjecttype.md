# aws.domainobjecttype

**CloudFormation type:** `AWS::CustomerProfiles::DomainObjectType`

Resource Type definition for AWS::CustomerProfiles::DomainObjectType

Region attribute: `region`

**Import ID:** `<region>/DomainName|ObjectTypeName` (AWS::CustomerProfiles::DomainObjectType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the domain object type was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the domain object type. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The unique name of the domain. |
| `EncryptionKey` | encryption_key | `string` | optional, computed, provider-chosen |  | The default encryption key |
| `Fields` |  | `map` | required, replaces on change |  | A map of the name and ObjectType field. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp of when the domain object type was most recently edited. |
| `ObjectTypeName` | object_type_name | `string` | required, replaces on change |  | The name of the domain object type. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported (parent resource required)
