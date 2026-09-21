# aws.resource

**CloudFormation type:** `AWS::ApiGateway::Resource`

The ``AWS::ApiGateway::Resource`` resource creates a resource in an API.

Region attribute: `region`

**Import ID:** `<region>/RestApiId|ResourceId` (AWS::ApiGateway::Resource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ParentId` | parent_id | `string` | required, replaces on change |  |  |
| `PathPart` | path_part | `string` | required, replaces on change |  |  |
| `ResourceId` | resource_id | `string` | computed |  |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |

Supports update: yes

Discovery: supported (parent resource required)
