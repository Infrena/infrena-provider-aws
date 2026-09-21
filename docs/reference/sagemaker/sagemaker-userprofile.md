# aws.sagemaker.userprofile

**CloudFormation type:** `AWS::SageMaker::UserProfile`

Resource Type definition for AWS::SageMaker::UserProfile

Region attribute: `region`

**Import ID:** `<region>/UserProfileName|DomainId` (AWS::SageMaker::UserProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DomainId` | domain_id | `string` | required, replaces on change | aws.sagemaker.domain.DomainId | The ID of the associated Domain. |
| `SingleSignOnUserIdentifier` | single_sign_on_user_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is "UserName". If the Domain's AuthMode is SSO, this field is required. If the Domain's AuthMode is not SSO, this field cannot be specified. |
| `SingleSignOnUserValue` | single_sign_on_user_value | `string` | optional, computed, provider-chosen, replaces on change |  | The username of the associated AWS Single Sign-On User for this UserProfile. If the Domain's AuthMode is SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not SSO, this field cannot be specified. |
| `Tags` |  | `map` | tags map |  | A list of tags to apply to the user profile. |
| `UserProfileArn` | user_profile_arn | `string` | computed |  | The user profile Amazon Resource Name (ARN). |
| `UserProfileName` | user_profile_name | `string` | required, replaces on change |  | A name for the UserProfile. |
| `UserSettings` | user_settings | `map` | optional, computed, provider-chosen |  | A collection of settings that apply to users of Amazon SageMaker Studio. These settings are specified when the CreateUserProfile API is called, and as DefaultUserSettings when the CreateDomain API is called. |

Supports update: yes

Discovery: supported
