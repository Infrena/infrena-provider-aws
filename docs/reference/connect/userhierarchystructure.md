# aws.userhierarchystructure

**CloudFormation type:** `AWS::Connect::UserHierarchyStructure`

Resource Type definition for AWS::Connect::UserHierarchyStructure

Region attribute: `region`

**Import ID:** `<region>/UserHierarchyStructureArn` (AWS::Connect::UserHierarchyStructure)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `UserHierarchyStructure` | user_hierarchy_structure | `map` | optional, computed, provider-chosen |  | Information about the hierarchy structure. |
| `UserHierarchyStructureArn` | user_hierarchy_structure_arn | `string` | computed |  | The identifier of the User Hierarchy Structure. |

Supports update: yes

Discovery: not supported
