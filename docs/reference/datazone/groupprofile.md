# aws.groupprofile

**CloudFormation type:** `AWS::DataZone::GroupProfile`

Group profiles represent groups of Amazon DataZone users. Groups can be manually created, or mapped to Active Directory groups of enterprise customers. In Amazon DataZone, groups serve two purposes. First, a group can map to a team of users in the organizational chart, and thus reduce the administrative work of a Amazon DataZone project owner when there are new employees joining or leaving a team. Second, corporate administrators use Active Directory groups to manage and update user statuses and so Amazon DataZone domain administrators can use these group memberships to implement Amazon DataZone domain policies.

Region attribute: `region`

**Import ID:** `<region>/DomainId|Id` (AWS::DataZone::GroupProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DomainId` | domain_id | `string` | computed |  | The identifier of the Amazon DataZone domain in which the group profile is created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  | The identifier of the Amazon DataZone domain in which the group profile would be created. |
| `GroupIdentifier` | group_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of the group. |
| `GroupName` | group_name | `string` | computed |  | The group-name of the Group Profile. |
| `GroupType` | group_type | `string` | optional, computed, provider-chosen, write-only |  | The type of the group. |
| `Id` |  | `string` | computed |  | The ID of the Amazon DataZone group profile. |
| `RolePrincipalArn` | role_principal_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ARN of the role principal for the group profile. |
| `RolePrincipalId` | role_principal_id | `string` | computed |  | The ID of the role principal for the group profile. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the group profile. |

Supports update: yes

Discovery: supported (parent resource required)
