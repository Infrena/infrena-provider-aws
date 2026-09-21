# aws.resourceexplorer2.view

**CloudFormation type:** `AWS::ResourceExplorer2::View`

Definition of AWS::ResourceExplorer2::View Resource Type

Region attribute: `region`

**Import ID:** `<region>/ViewArn` (AWS::ResourceExplorer2::View)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Filters` |  | `map` | optional, computed, provider-chosen |  |  |
| `IncludedProperties` | included_properties | `list` | optional, computed, provider-chosen |  |  |
| `Scope` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `ViewArn` | view_arn | `string` | computed |  |  |
| `ViewName` | view_name | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
