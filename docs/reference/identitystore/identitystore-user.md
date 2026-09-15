# aws.identitystore.user

**CloudFormation type:** `AWS::IdentityStore::User`

Creates a user within the specified identity store.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IdentityStore::User)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Addresses` |  | `list` | optional, computed, provider-chosen, replaces on change |  | A list of addresses associated with the user. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the user. |
| `Birthdate` |  | `string` | optional, computed, provider-chosen |  | The user's birthdate in YYYY-MM-DD format. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time the user was created. |
| `CreatedBy` | created_by | `string` | computed |  | The identifier of the user or system that created the user. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | A string containing the name of the user for display. |
| `Emails` |  | `list` | optional, computed, provider-chosen, replaces on change |  | A list of email addresses associated with the user. |
| `IdentityStoreId` | identity_store_id | `string` | required, replaces on change |  | The globally unique identifier for the identity store. |
| `Locale` |  | `string` | optional, computed, provider-chosen |  | The geographical region or location of the user. |
| `Name` |  | `map` | optional, computed, provider-chosen |  | The name of the user. |
| `NickName` | nick_name | `string` | optional, computed, provider-chosen |  | An alternate name for the user. |
| `PhoneNumbers` | phone_numbers | `list` | optional, computed, provider-chosen, replaces on change |  | A list of phone numbers associated with the user. |
| `Photos` |  | `list` | optional, computed, provider-chosen, replaces on change |  | A list of photos associated with the user. |
| `PreferredLanguage` | preferred_language | `string` | optional, computed, provider-chosen |  | The preferred language of the user. |
| `ProfileUrl` | profile_url | `string` | optional, computed, provider-chosen |  | A URL associated with the user. |
| `Roles` |  | `list` | optional, computed, provider-chosen, replaces on change |  | A list of roles associated with the user. |
| `Timezone` |  | `string` | optional, computed, provider-chosen |  | The time zone for the user. |
| `Title` |  | `string` | optional, computed, provider-chosen |  | The title of the user. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time the user was last updated. |
| `UpdatedBy` | updated_by | `string` | computed |  | The identifier of the user or system that last updated the user. |
| `UserId` | user_id | `string` | computed |  | The identifier for a user in the identity store. |
| `UserName` | user_name | `string` | optional, computed, provider-chosen, replaces on change |  | A unique string used to identify the user. |
| `UserStatus` | user_status | `string` | computed |  | The current status of the user account. |
| `UserType` | user_type | `string` | optional, computed, provider-chosen |  | A string indicating the type of user. |
| `Website` |  | `string` | optional, computed, provider-chosen |  | The user's personal website or blog URL. |

Supports update: yes

Discovery: supported (parent resource required)
