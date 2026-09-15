# aws.logs.resourcepolicy

**CloudFormation type:** `AWS::Logs::ResourcePolicy`

The resource schema for AWSLogs ResourcePolicy

Region attribute: `region`

**Import ID:** `<region>/PolicyName` (AWS::Logs::ResourcePolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PolicyDocument` | policy_document | `string` | required |  | The policy document |
| `PolicyName` | policy_name | `string` | required, replaces on change |  | A name for resource policy |

Supports update: yes

Discovery: supported
