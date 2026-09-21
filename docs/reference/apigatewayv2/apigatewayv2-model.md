# aws.apigatewayv2.model

**CloudFormation type:** `AWS::ApiGatewayV2::Model`

The ``AWS::ApiGatewayV2::Model`` resource updates data model for a WebSocket API. For more information, see [Model Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/ApiId|ModelId` (AWS::ApiGatewayV2::Model)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.apigatewayv2.api.ApiId | The API identifier. |
| `ContentType` | content_type | `string` | optional, computed, provider-chosen |  | The content-type for the model, for example, "application/json". |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the model. |
| `ModelId` | model_id | `string` | computed |  |  |
| `Name` |  | `string` | required |  | The name of the model. |
| `Schema` |  | `map` | required |  | The schema for the model. For application/json models, this should be JSON schema draft 4 model. |

Supports update: yes

Discovery: supported (parent resource required)
