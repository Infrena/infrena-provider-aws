# aws.sso.instance

**CloudFormation type:** `AWS::SSO::Instance`

Resource Type definition for Identity Center (SSO) Instance

Region attribute: `region`

**Import ID:** `<region>/InstanceArn` (AWS::SSO::Instance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `IdentityStoreId` | identity_store_id | `string` | computed |  | The ID of the identity store associated with the created Identity Center (SSO) Instance |
| `InstanceArn` | instance_arn | `string` | computed |  | The SSO Instance ARN that is returned upon creation of the Identity Center (SSO) Instance |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name you want to assign to this Identity Center (SSO) Instance |
| `OwnerAccountId` | owner_account_id | `string` | computed |  | The AWS accountId of the owner of the Identity Center (SSO) Instance |
| `Status` |  | `string` | computed |  | The status of the Identity Center (SSO) Instance, create_in_progress/delete_in_progress/active |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
