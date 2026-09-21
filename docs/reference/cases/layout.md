# aws.layout

**CloudFormation type:** `AWS::Cases::Layout`

A layout in the Cases domain. Layouts define the following configuration in the top section and More Info tab of the Cases user interface: Fields to display to the users and Field ordering.

Region attribute: `region`

**Import ID:** `<region>/LayoutArn` (AWS::Cases::Layout)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Content` |  | `string` | required |  | Defines the layout structure and field organization for the case interface. Specifies which fields appear in the top panel and More Info tab, and their display order. |
| `CreatedTime` | created_time | `string` | computed |  | The time at which the layout was created. |
| `DomainId` | domain_id | `string` | optional, computed, provider-chosen, replaces on change | aws.cases.domain.DomainId | The unique identifier of the Cases domain. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The time at which the layout was created or last modified. |
| `LayoutArn` | layout_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the layout. |
| `LayoutId` | layout_id | `string` | computed |  | The unique identifier of the layout. |
| `Name` |  | `string` | required |  | A descriptive name for the layout. Must be unique within the Cases domain and should clearly indicate the layout's purpose and field organization. |
| `Tags` |  | `map` | tags map |  | The tags that you attach to this layout. |

Supports update: yes

Discovery: supported (parent resource required)
