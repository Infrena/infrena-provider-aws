# aws.datazone.userprofile

**CloudFormation type:** `AWS::DataZone::UserProfile`

A user profile represents Amazon DataZone users. Amazon DataZone supports both IAM roles and SSO identities to interact with the Amazon DataZone Management Console and the data portal for different purposes. Domain administrators use IAM roles to perform the initial administrative domain-related work in the Amazon DataZone Management Console, including creating new Amazon DataZone domains, configuring metadata form types, and implementing policies. Data workers use their SSO corporate identities via Identity Center to log into the Amazon DataZone Data Portal and access projects where they have memberships.

Region attribute: `region`

**Import ID:** `<region>/DomainId|Id` (AWS::DataZone::UserProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Details` |  | `string` | computed |  |  |
| `DomainId` | domain_id | `string` | computed |  | The identifier of the Amazon DataZone domain in which the user profile is created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  | The identifier of the Amazon DataZone domain in which the user profile would be created. |
| `Id` |  | `string` | computed |  | The ID of the Amazon DataZone user profile. |
| `SessionName` | session_name | `string` | optional, computed, provider-chosen, write-only |  | The session name of the user profile. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the user profile. |
| `Type` | type_value | `string` | computed |  | The type of the user profile. |
| `UserIdentifier` | user_identifier | `string` | required, replaces on change, write-only |  | The ID of the user. |
| `UserType` | user_type | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The type of the user. |

Supports update: yes

Discovery: supported (parent resource required)
