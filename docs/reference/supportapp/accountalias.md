# aws.accountalias

**CloudFormation type:** `AWS::SupportApp::AccountAlias`

An AWS Support App resource that creates, updates, reads, and deletes a customer's account alias.

Region attribute: `region`

**Import ID:** `<region>/AccountAliasResourceId` (AWS::SupportApp::AccountAlias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountAlias` | account_alias | `string` | required |  | An account alias associated with a customer's account. |
| `AccountAliasResourceId` | account_alias_resource_id | `string` | computed |  | Unique identifier representing an alias tied to an account |

Supports update: yes

Discovery: supported
