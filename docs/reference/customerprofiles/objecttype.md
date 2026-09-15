# aws.objecttype

**CloudFormation type:** `AWS::CustomerProfiles::ObjectType`

An ObjectType resource of Amazon Connect Customer Profiles

Region attribute: `region`

**Import ID:** `<region>/DomainName|ObjectTypeName` (AWS::CustomerProfiles::ObjectType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowProfileCreation` | allow_profile_creation | `boolean` | optional, computed, provider-chosen |  | Indicates whether a profile should be created when data is received. |
| `CreatedAt` | created_at | `string` | computed |  | The time of this integration got created. |
| `Description` |  | `string` | required |  | Description of the profile object type. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The unique name of the domain. |
| `EncryptionKey` | encryption_key | `string` | optional, computed, provider-chosen |  | The default encryption key |
| `ExpirationDays` | expiration_days | `integer` | optional, computed, provider-chosen |  | The default number of days until the data within the domain expires. |
| `Fields` |  | `list` | optional, computed, provider-chosen |  | A list of the name and ObjectType field. |
| `Keys` |  | `list` | optional, computed, provider-chosen |  | A list of unique keys that can be used to map data to the profile. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The time of this integration got last updated at. |
| `MaxAvailableProfileObjectCount` | max_available_profile_object_count | `integer` | computed |  | The maximum available number of profile objects |
| `MaxProfileObjectCount` | max_profile_object_count | `integer` | optional, computed, provider-chosen |  | The maximum number of profile objects for this object type |
| `ObjectTypeName` | object_type_name | `string` | required, replaces on change |  | The name of the profile object type. |
| `SourceLastUpdatedTimestampFormat` | source_last_updated_timestamp_format | `string` | optional, computed, provider-chosen |  | The format of your sourceLastUpdatedTimestamp that was previously set up. |
| `SourcePriority` | source_priority | `integer` | optional, computed, provider-chosen |  | Defines the priority order of object types. Lower value indicates higher priority. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the integration. |
| `TemplateId` | template_id | `string` | optional, computed, provider-chosen |  | A unique identifier for the object template. |

Supports update: yes

Discovery: supported (parent resource required)
