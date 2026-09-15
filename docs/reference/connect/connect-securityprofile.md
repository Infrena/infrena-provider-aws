# aws.connect.securityprofile

**CloudFormation type:** `AWS::Connect::SecurityProfile`

Resource Type definition for AWS::Connect::SecurityProfile

Region attribute: `region`

**Import ID:** `<region>/SecurityProfileArn` (AWS::Connect::SecurityProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedAccessControlHierarchyGroupId` | allowed_access_control_hierarchy_group_id | `string` | optional, computed, provider-chosen |  | The identifier of the hierarchy group that a security profile uses to restrict access to resources in Amazon Connect. |
| `AllowedAccessControlTags` | allowed_access_control_tags | `list` | optional, computed, provider-chosen |  | The list of tags that a security profile uses to restrict access to resources in Amazon Connect. |
| `AllowedFlowModules` | allowed_flow_modules | `list` | optional, computed, provider-chosen |  | The list of flow-module resources to be linked to a security profile in Amazon Connect. |
| `Applications` |  | `list` | optional, computed, provider-chosen |  | A list of third-party applications that the security profile will give access to. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the security profile. |
| `GranularAccessControlConfiguration` | granular_access_control_configuration | `map` | optional, computed, provider-chosen |  |  |
| `HierarchyRestrictedResources` | hierarchy_restricted_resources | `list` | optional, computed, provider-chosen |  | The list of resources that a security profile applies hierarchy restrictions to in Amazon Connect. |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `LastModifiedRegion` | last_modified_region | `string` | computed |  | The AWS Region where this resource was last modified. |
| `LastModifiedTime` | last_modified_time | `float` | computed |  | The timestamp when this resource was last modified. |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  | Permissions assigned to the security profile. |
| `SecurityProfileArn` | security_profile_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the security profile. |
| `SecurityProfileName` | security_profile_name | `string` | required, replaces on change |  | The name of the security profile. |
| `TagRestrictedResources` | tag_restricted_resources | `list` | optional, computed, provider-chosen |  | The list of resources that a security profile applies tag restrictions to in Amazon Connect. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags used to organize, track, or control access for this resource. |

Supports update: yes

Discovery: supported (parent resource required)
