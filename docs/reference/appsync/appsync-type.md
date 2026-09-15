# aws.appsync.type

**CloudFormation type:** `AWS::AppSync::Type`

Represents a GraphQL type in an AWS AppSync GraphQL API.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AppSync::Type)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.appsync.api.ApiId | The API ID. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the type. |
| `Definition` |  | `string` | required |  | The type definition, in GraphQL Schema Definition Language (SDL) format. |
| `Format` |  | `string` | required |  | The type format: SDL or JSON. |
| `Name` |  | `string` | computed |  | The type name. |

Supports update: yes

Discovery: supported (parent resource required)
