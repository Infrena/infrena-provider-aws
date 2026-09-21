# aws.gatewayresponse

**CloudFormation type:** `AWS::ApiGateway::GatewayResponse`

The ``AWS::ApiGateway::GatewayResponse`` resource creates a gateway response for your API. When you delete a stack containing this resource, your custom gateway responses are reset. For more information, see [API Gateway Responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/customize-gateway-responses.html#api-gateway-gatewayResponse-definition) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ApiGateway::GatewayResponse)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `ResponseParameters` | response_parameters | `map` | optional, computed, provider-chosen |  |  |
| `ResponseTemplates` | response_templates | `map` | optional, computed, provider-chosen |  |  |
| `ResponseType` | response_type | `string` | required, replaces on change |  |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |
| `StatusCode` | status_code | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
