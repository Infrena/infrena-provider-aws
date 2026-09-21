# aws.delegatedadmin

**CloudFormation type:** `AWS::SecurityHub::DelegatedAdmin`

The ``AWS::SecurityHub::DelegatedAdmin`` resource designates the delegated ASHlong administrator account for an organization. You must enable the integration between ASH and AOlong before you can designate a delegated ASH administrator. Only the management account for an organization can designate the delegated ASH administrator account. For more information, see [Designating the delegated administrator](https://docs.aws.amazon.com/securityhub/latest/userguide/designate-orgs-admin-account.html#designate-admin-instructions) in the *User Guide*.

Region attribute: `region`

**Import ID:** `<region>/DelegatedAdminIdentifier` (AWS::SecurityHub::DelegatedAdmin)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdminAccountId` | admin_account_id | `string` | required, replaces on change |  | The AWS-account identifier of the account to designate as the Security Hub CSPM administrator account. |
| `DelegatedAdminIdentifier` | delegated_admin_identifier | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |

Supports update: no

Discovery: supported
