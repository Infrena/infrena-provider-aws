# aws.qbusiness.application

**CloudFormation type:** `AWS::QBusiness::Application`

Definition of AWS::QBusiness::Application Resource Type

Region attribute: `region`

**Import ID:** `<region>/ApplicationId` (AWS::QBusiness::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationArn` | application_arn | `string` | computed |  |  |
| `ApplicationId` | application_id | `string` | computed |  |  |
| `AttachmentsConfiguration` | attachments_configuration | `map` | optional, computed, provider-chosen |  |  |
| `AutoSubscriptionConfiguration` | auto_subscription_configuration | `map` | optional, computed, provider-chosen |  |  |
| `ClientIdsForOIDC` | client_ids_for_oidc | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `EncryptionConfiguration` | encryption_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `IamIdentityProviderArn` | iam_identity_provider_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `IdentityCenterApplicationArn` | identity_center_application_arn | `string` | computed |  |  |
| `IdentityCenterInstanceArn` | identity_center_instance_arn | `string` | optional, computed, provider-chosen, write-only |  |  |
| `IdentityType` | identity_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PersonalizationConfiguration` | personalization_configuration | `map` | optional, computed, provider-chosen |  |  |
| `QAppsConfiguration` | q_apps_configuration | `map` | optional, computed, provider-chosen |  |  |
| `QuickSightConfiguration` | quick_sight_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported
