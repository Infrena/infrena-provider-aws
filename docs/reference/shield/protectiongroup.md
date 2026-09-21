# aws.protectiongroup

**CloudFormation type:** `AWS::Shield::ProtectionGroup`

A grouping of protected resources so they can be handled as a collective. This resource grouping improves the accuracy of detection and reduces false positives.

Region attribute: `region`

**Import ID:** `<region>/ProtectionGroupArn` (AWS::Shield::ProtectionGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Aggregation` |  | `string` | required |  | Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events. |
| `Members` |  | `list` | optional, computed, provider-chosen |  | The Amazon Resource Names (ARNs) of the resources to include in the protection group. You must set this when you set `Pattern` to `ARBITRARY` and you must not set it for any other `Pattern` setting. |
| `Pattern` |  | `string` | required |  | The criteria to use to choose the protected resources for inclusion in the group. You can include all resources that have protections, provide a list of resource Amazon Resource Names (ARNs), or include all resources of a specified resource type. |
| `ProtectionGroupArn` | protection_group_arn | `string` | computed |  | The ARN (Amazon Resource Name) of the protection group. |
| `ProtectionGroupId` | protection_group_id | `string` | required, replaces on change | aws.protectiongroup.ProtectionGroupId | The name of the protection group. You use this to identify the protection group in lists and to manage the protection group, for example to update, delete, or describe it. |
| `ResourceType` | resource_type | `string` | optional, computed, provider-chosen |  | The resource type to include in the protection group. All protected resources of this type are included in the protection group. Newly protected resources of this type are automatically added to the group. You must set this when you set `Pattern` to `BY_RESOURCE_TYPE` and you must not set it for any other `Pattern` setting. |
| `Tags` |  | `map` | tags map |  | One or more tag key-value pairs for the Protection object. |

Supports update: yes

Discovery: supported
