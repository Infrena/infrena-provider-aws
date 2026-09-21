# aws.routeresponse

**CloudFormation type:** `AWS::ApiGatewayV2::RouteResponse`

The ``AWS::ApiGatewayV2::RouteResponse`` resource creates a route response for a WebSocket API. For more information, see [Set up Route Responses for a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-route-response.html) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/ApiId|RouteId|RouteResponseId` (AWS::ApiGatewayV2::RouteResponse)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.apigatewayv2.api.ApiId | The API identifier. |
| `ModelSelectionExpression` | model_selection_expression | `string` | optional, computed, provider-chosen |  | The model selection expression for the route response. Supported only for WebSocket APIs. |
| `ResponseModels` | response_models | `map` | optional, computed, provider-chosen |  | The response models for the route response. |
| `ResponseParameters` | response_parameters | `map` | optional, computed, provider-chosen |  | The route response parameters. |
| `RouteId` | route_id | `string` | required, replaces on change | aws.apigatewayv2.route.RouteId | The route ID. |
| `RouteResponseId` | route_response_id | `string` | computed |  |  |
| `RouteResponseKey` | route_response_key | `string` | required |  | The route response key. |

Supports update: yes

Discovery: supported
