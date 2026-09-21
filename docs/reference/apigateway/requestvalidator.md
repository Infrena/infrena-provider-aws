# aws.requestvalidator

**CloudFormation type:** `AWS::ApiGateway::RequestValidator`

The ``AWS::ApiGateway::RequestValidator`` resource sets up basic validation rules for incoming requests to your API. For more information, see [Enable Basic Request Validation for an API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-method-request-validation.html) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/RestApiId|RequestValidatorId` (AWS::ApiGateway::RequestValidator)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `RequestValidatorId` | request_validator_id | `string` | computed |  |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |
| `ValidateRequestBody` | validate_request_body | `boolean` | optional, computed, provider-chosen |  |  |
| `ValidateRequestParameters` | validate_request_parameters | `boolean` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
