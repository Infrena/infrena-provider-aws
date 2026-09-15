# aws.opensearchserverless.accesspolicy

**CloudFormation type:** `AWS::OpenSearchServerless::AccessPolicy`

Resource Type definition for AWS::OpenSearchServerless::AccessPolicy

Region attribute: `region`

**Import ID:** `<region>/Type|Name` (AWS::OpenSearchServerless::AccessPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the policy |
| `Name` |  | `string` | required, replaces on change |  | The name of the policy |
| `Policy` |  | `string` | required |  | The JSON policy document that is the content for the policy |
| `Type` | type_value | `string` | required, replaces on change |  | The possible types for the access policy |

Supports update: yes

Discovery: supported (parent resource required)
