# aws.integrationresponse

**CloudFormation type:** `AWS::ApiGatewayV2::IntegrationResponse`

The ``AWS::ApiGatewayV2::IntegrationResponse`` resource updates an integration response for an WebSocket API. For more information, see [Set up WebSocket API Integration Responses in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-integration-responses.html) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/ApiId|IntegrationId|IntegrationResponseId` (AWS::ApiGatewayV2::IntegrationResponse)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.apigatewayv2.api.ApiId | The API identifier. |
| `ContentHandlingStrategy` | content_handling_strategy | `string` | optional, computed, provider-chosen |  | Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are ``CONVERT_TO_BINARY`` and ``CONVERT_TO_TEXT``, with the following behaviors: |
| `IntegrationId` | integration_id | `string` | required, replaces on change | aws.apigatewayv2.integration.IntegrationId | The integration ID. |
| `IntegrationResponseId` | integration_response_id | `string` | computed |  |  |
| `IntegrationResponseKey` | integration_response_key | `string` | required |  | The integration response key. |
| `ResponseParameters` | response_parameters | `map` | optional, computed, provider-chosen |  | A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of ``method.response.header.{name}``, where name is a valid and unique header name. The mapped non-static value must match the pattern of ``integration.response.header.{name}`` or ``integration.response.body.{JSON-expression}``, where ``{name}`` is a valid and unique response header name and ``{JSON-expression}`` is a valid JSON expression without the ``$`` prefix. |
| `ResponseTemplates` | response_templates | `map` | optional, computed, provider-chosen |  | The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value. |
| `TemplateSelectionExpression` | template_selection_expression | `string` | optional, computed, provider-chosen |  | The template selection expression for the integration response. Supported only for WebSocket APIs. |

Supports update: yes

Discovery: supported
