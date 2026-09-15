# aws.documentationversion

**CloudFormation type:** `AWS::ApiGateway::DocumentationVersion`

The ``AWS::ApiGateway::DocumentationVersion`` resource creates a snapshot of the documentation for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/DocumentationVersion|RestApiId` (AWS::ApiGateway::DocumentationVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DocumentationVersion` | documentation_version | `string` | required, replaces on change |  |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |

Supports update: yes

Discovery: supported (parent resource required)
