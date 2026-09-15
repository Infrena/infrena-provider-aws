# aws.customerprofiles.integration

**CloudFormation type:** `AWS::CustomerProfiles::Integration`

The resource schema for creating an Amazon Connect Customer Profiles Integration.

Region attribute: `region`

**Import ID:** `<region>/DomainName|Uri` (AWS::CustomerProfiles::Integration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time of this integration got created |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The unique name of the domain. |
| `EventTriggerNames` | event_trigger_names | `list` | optional, computed, provider-chosen |  | A list of unique names for active event triggers associated with the integration. |
| `FlowDefinition` | flow_definition | `map` | optional, computed, provider-chosen, write-only |  |  |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The time of this integration got last updated at |
| `ObjectTypeName` | object_type_name | `string` | optional, computed, provider-chosen |  | The name of the ObjectType defined for the 3rd party data in Profile Service |
| `ObjectTypeNames` | object_type_names | `list` | optional, computed, provider-chosen |  | The mapping between 3rd party event types and ObjectType names |
| `Scope` |  | `string` | optional, computed, provider-chosen |  | Scope of the integration, such as 'PROFILE' or 'DOMAIN' |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the integration |
| `Uri` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The URI of the S3 bucket or any other type of data source. |

Supports update: yes

Discovery: supported (parent resource required)
