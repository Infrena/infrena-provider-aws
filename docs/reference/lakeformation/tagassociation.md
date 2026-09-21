# aws.tagassociation

**CloudFormation type:** `AWS::LakeFormation::TagAssociation`

A resource schema representing a Lake Formation Tag Association. While tag associations are not explicit Lake Formation resources, this CloudFormation resource can be used to associate tags with Lake Formation entities.

Region attribute: `region`

**Import ID:** `<region>/ResourceIdentifier|TagsIdentifier` (AWS::LakeFormation::TagAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LFTags` | lf_tags | `list` | required, replaces on change |  | List of Lake Formation Tags to associate with the Lake Formation Resource |
| `Resource` |  | `map` | required, replaces on change |  | Resource to tag with the Lake Formation Tags |
| `ResourceIdentifier` | resource_identifier | `string` | computed |  | Unique string identifying the resource. Used as primary identifier, which ideally should be a string |
| `TagsIdentifier` | tags_identifier | `string` | computed |  | Unique string identifying the resource's tags. Used as primary identifier, which ideally should be a string |

Supports update: no

Discovery: supported (parent resource required)
