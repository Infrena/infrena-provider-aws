# aws.method

**CloudFormation type:** `AWS::ApiGateway::Method`

The ``AWS::ApiGateway::Method`` resource creates API Gateway methods that define the parameters and body that clients must send in their requests.

Region attribute: `region`

**Import ID:** `<region>/RestApiId|ResourceId|HttpMethod` (AWS::ApiGateway::Method)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiKeyRequired` | api_key_required | `boolean` | optional, computed, provider-chosen |  |  |
| `AuthorizationScopes` | authorization_scopes | `list` | optional, computed, provider-chosen |  |  |
| `AuthorizationType` | authorization_type | `string` | optional, computed, provider-chosen |  | The method's authorization type. This parameter is required. For valid values, see [Method](https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html) in the *API Gateway API Reference*. |
| `AuthorizerId` | authorizer_id | `string` | optional, computed, provider-chosen | aws.apigateway.authorizer.AuthorizerId |  |
| `HttpMethod` | http_method | `string` | required, replaces on change |  |  |
| `Integration` |  | `map` | optional, computed, provider-chosen |  | ``Integration`` is a property of the [AWS::ApiGateway::Method](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-method.html) resource that specifies information about the target backend that a method calls. |
| `MethodResponses` | method_responses | `list` | optional, computed, provider-chosen |  |  |
| `OperationName` | operation_name | `string` | optional, computed, provider-chosen |  |  |
| `RequestModels` | request_models | `map` | optional, computed, provider-chosen |  |  |
| `RequestParameters` | request_parameters | `map` | optional, computed, provider-chosen |  |  |
| `RequestValidatorId` | request_validator_id | `string` | optional, computed, provider-chosen | aws.requestvalidator.RequestValidatorId |  |
| `ResourceId` | resource_id | `string` | required, replaces on change | aws.resource.ResourceId |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |

Supports update: yes

Discovery: not supported
