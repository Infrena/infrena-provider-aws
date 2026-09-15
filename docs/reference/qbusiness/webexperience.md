# aws.webexperience

**CloudFormation type:** `AWS::QBusiness::WebExperience`

Definition of AWS::QBusiness::WebExperience Resource Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|WebExperienceId` (AWS::QBusiness::WebExperience)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.qbusiness.application.ApplicationId |  |
| `BrowserExtensionConfiguration` | browser_extension_configuration | `map` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `CustomizationConfiguration` | customization_configuration | `map` | optional, computed, provider-chosen |  |  |
| `DefaultEndpoint` | default_endpoint | `string` | computed |  |  |
| `IdentityProviderConfiguration` | identity_provider_configuration | `string` | optional, computed, provider-chosen |  |  |
| `Origins` |  | `list` | optional, computed, provider-chosen |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `SamplePromptsControlMode` | sample_prompts_control_mode | `string` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | computed |  |  |
| `Subtitle` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Title` |  | `string` | optional, computed, provider-chosen |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |
| `WebExperienceArn` | web_experience_arn | `string` | computed |  |  |
| `WebExperienceId` | web_experience_id | `string` | computed |  |  |
| `WelcomeMessage` | welcome_message | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
