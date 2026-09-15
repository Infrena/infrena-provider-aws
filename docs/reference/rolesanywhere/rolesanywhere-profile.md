# aws.rolesanywhere.profile

**CloudFormation type:** `AWS::RolesAnywhere::Profile`

Definition of AWS::RolesAnywhere::Profile Resource Type

Region attribute: `region`

**Import ID:** `<region>/ProfileId` (AWS::RolesAnywhere::Profile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptRoleSessionName` | accept_role_session_name | `boolean` | optional, computed, provider-chosen |  |  |
| `AttributeMappings` | attribute_mappings | `list` | optional, computed, provider-chosen |  |  |
| `DurationSeconds` | duration_seconds | `float` | optional, computed, provider-chosen |  |  |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  |  |
| `ManagedPolicyArns` | managed_policy_arns | `list` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `ProfileArn` | profile_arn | `string` | computed |  |  |
| `ProfileId` | profile_id | `string` | computed |  |  |
| `RequireInstanceProperties` | require_instance_properties | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `RoleArns` | role_arns | `list` | required | aws.role.Arn |  |
| `SessionPolicy` | session_policy | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
