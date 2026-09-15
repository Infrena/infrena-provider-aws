# aws.resourcecollection

**CloudFormation type:** `AWS::DevOpsGuru::ResourceCollection`

This resource schema represents the ResourceCollection resource in the Amazon DevOps Guru.

Region attribute: `region`

**Import ID:** `<region>/ResourceCollectionType` (AWS::DevOpsGuru::ResourceCollection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ResourceCollectionFilter` | resource_collection_filter | `map` | required |  | Information about a filter used to specify which AWS resources are analyzed for anomalous behavior by DevOps Guru. |
| `ResourceCollectionType` | resource_collection_type | `string` | computed |  | The type of ResourceCollection |

Supports update: yes

Discovery: supported
