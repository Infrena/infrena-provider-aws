# aws.organizationadmin

**CloudFormation type:** `AWS::Detective::OrganizationAdmin`

Resource schema for AWS::Detective::OrganizationAdmin

Region attribute: `region`

**Import ID:** `<region>/AccountId` (AWS::Detective::OrganizationAdmin)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | required, replaces on change |  | The account ID of the account that should be registered as your Organization's delegated administrator for Detective |
| `GraphArn` | graph_arn | `string` | computed |  | The Detective graph ARN |

Supports update: yes

Discovery: supported
