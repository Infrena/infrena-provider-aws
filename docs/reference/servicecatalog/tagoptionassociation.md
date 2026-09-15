# aws.tagoptionassociation

**CloudFormation type:** `AWS::ServiceCatalog::TagOptionAssociation`

Resource Type definition for AWS::ServiceCatalog::TagOptionAssociation

Region attribute: `region`

**Import ID:** `<region>/TagOptionId|ResourceId` (AWS::ServiceCatalog::TagOptionAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ResourceId` | resource_id | `string` | optional, computed, provider-chosen, replaces on change |  | The CloudformationProduct or Portfolio identifier. |
| `TagOptionId` | tag_option_id | `string` | optional, computed, provider-chosen, replaces on change | aws.tagoption.Id | The TagOption identifier. |

Supports update: no

Discovery: supported (parent resource required)
