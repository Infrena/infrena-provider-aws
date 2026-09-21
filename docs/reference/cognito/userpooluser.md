# aws.userpooluser

**CloudFormation type:** `AWS::Cognito::UserPoolUser`

Resource Type definition for AWS::Cognito::UserPoolUser

Region attribute: `region`

**Import ID:** `<region>/UserPoolId|Username` (AWS::Cognito::UserPoolUser)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ClientMetadata` | client_metadata | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `DesiredDeliveryMediums` | desired_delivery_mediums | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ForceAliasCreation` | force_alias_creation | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `MessageAction` | message_action | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `UserAttributes` | user_attributes | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `UserPoolId` | user_pool_id | `string` | required, replaces on change | aws.userpool.UserPoolId |  |
| `Username` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ValidationData` | validation_data | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |

Supports update: no

Discovery: supported (parent resource required)
