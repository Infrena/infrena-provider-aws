# aws.organizations.account

**CloudFormation type:** `AWS::Organizations::Account`

You can use AWS::Organizations::Account to manage accounts in organization.

Global type (no region attribute)

**Import ID:** `global/AccountId` (AWS::Organizations::Account)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  | If the account was created successfully, the unique identifier (ID) of the new account. |
| `AccountName` | account_name | `string` | required |  | The friendly name of the member account. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the account. |
| `Email` |  | `string` | required |  | The email address of the owner to assign to the new member account. |
| `JoinedMethod` | joined_method | `string` | computed |  | The method by which the account joined the organization. |
| `JoinedTimestamp` | joined_timestamp | `string` | computed |  | The date the account became a part of the organization. |
| `ParentIds` | parent_ids | `list` | optional, computed, provider-chosen |  | List of parent nodes for the member account. Currently only one parent at a time is supported. Default is root. |
| `Paths` |  | `list` | computed |  | The paths in the organization where the account exists. |
| `RoleName` | role_name | `string` | optional, computed, provider-chosen, write-only |  | The name of an IAM role that AWS Organizations automatically preconfigures in the new member account. Default name is OrganizationAccountAccessRole if not specified. |
| `State` |  | `string` | computed |  | The state of the account in the organization. |
| `Status` |  | `string` | computed |  | The status of the account in the organization. |
| `Tags` |  | `map` | tags map |  | A list of tags that you want to attach to the newly created account. For each tag in the list, you must specify both a tag key and a value. |

Supports update: yes

Discovery: supported
