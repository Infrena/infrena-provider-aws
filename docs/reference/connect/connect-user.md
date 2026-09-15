# aws.connect.user

**CloudFormation type:** `AWS::Connect::User`

Resource Type definition for AWS::Connect::User

Region attribute: `region`

**Import ID:** `<region>/UserArn` (AWS::Connect::User)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AfterContactWorkConfigs` | after_contact_work_configs | `list` | optional, computed, provider-chosen |  | After Contact Work configurations of a user. |
| `AutoAcceptConfigs` | auto_accept_configs | `list` | optional, computed, provider-chosen |  | Auto-accept configurations of a user. |
| `DirectoryUserId` | directory_user_id | `string` | optional, computed, provider-chosen |  | The identifier of the user account in the directory used for identity management. |
| `HierarchyGroupArn` | hierarchy_group_arn | `string` | optional, computed, provider-chosen |  | The identifier of the hierarchy group for the user. |
| `IdentityInfo` | identity_info | `map` | optional, computed, provider-chosen |  | Contains information about the identity of a user. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Password` |  | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The password for the user account. A password is required if you are using Amazon Connect for identity management. Otherwise, it is an error to include a password. |
| `PersistentConnectionConfigs` | persistent_connection_configs | `list` | optional, computed, provider-chosen |  | Persistent Connection configurations of a user. |
| `PhoneConfig` | phone_config | `map` | optional, computed, provider-chosen |  | Contains information about the phone configuration settings for a user. |
| `PhoneNumberConfigs` | phone_number_configs | `list` | optional, computed, provider-chosen |  | Phone Number configurations of a user. |
| `RoutingProfileArn` | routing_profile_arn | `string` | required | aws.routingprofile.RoutingProfileArn | The identifier of the routing profile for the user. |
| `SecurityProfileArns` | security_profile_arns | `list` | required | aws.connect.securityprofile.SecurityProfileArn | One or more security profile arns for the user |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |
| `UserArn` | user_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the user. |
| `UserProficiencies` | user_proficiencies | `list` | optional, computed, provider-chosen |  | One or more predefined attributes assigned to a user, with a level that indicates how skilled they are. |
| `Username` |  | `string` | required |  | The user name for the account. |
| `VoiceEnhancementConfigs` | voice_enhancement_configs | `list` | optional, computed, provider-chosen |  | Voice Enhancement configurations of a user. |

Supports update: yes

Discovery: supported (parent resource required)
