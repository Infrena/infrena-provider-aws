# aws.apigatewayv2.authorizer

**CloudFormation type:** `AWS::ApiGatewayV2::Authorizer`

The ``AWS::ApiGatewayV2::Authorizer`` resource creates an authorizer for a WebSocket API or an HTTP API. To learn more, see [Controlling and managing access to a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-control-access.html) and [Controlling and managing access to an HTTP API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-access-control.html) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/AuthorizerId|ApiId` (AWS::ApiGatewayV2::Authorizer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.apigatewayv2.api.ApiId | The API identifier. |
| `AuthorizerCredentialsArn` | authorizer_credentials_arn | `string` | optional, computed, provider-chosen, write-only |  | Specifies the required credentials as an IAM role for API Gateway to invoke the authorizer. To specify an IAM role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To use resource-based permissions on the Lambda function, specify null. Supported only for ``REQUEST`` authorizers. |
| `AuthorizerId` | authorizer_id | `string` | computed |  |  |
| `AuthorizerPayloadFormatVersion` | authorizer_payload_format_version | `string` | optional, computed, provider-chosen, write-only |  | Specifies the format of the payload sent to an HTTP API Lambda authorizer. Required for HTTP API Lambda authorizers. Supported values are ``1.0`` and ``2.0``. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html). |
| `AuthorizerResultTtlInSeconds` | authorizer_result_ttl_in_seconds | `integer` | optional, computed, provider-chosen, write-only |  | The time to live (TTL) for cached authorizer results, in seconds. If it equals 0, authorization caching is disabled. If it is greater than 0, API Gateway caches authorizer responses. The maximum value is 3600, or 1 hour. Supported only for HTTP API Lambda authorizers. |
| `AuthorizerType` | authorizer_type | `string` | required |  | The authorizer type. Specify ``REQUEST`` for a Lambda function using incoming request parameters. Specify ``JWT`` to use JSON Web Tokens (supported only for HTTP APIs). |
| `AuthorizerUri` | authorizer_uri | `string` | optional, computed, provider-chosen, write-only |  | The authorizer's Uniform Resource Identifier (URI). For ``REQUEST`` authorizers, this must be a well-formed Lambda function URI, for example, ``arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations``. In general, the URI has this form: ``arn:aws:apigateway:{region}:lambda:path/{service_api}``, where *{region}* is the same as the region hosting the Lambda function, path indicates that the remaining substring in the URI should be treated as the path to the resource, including the initial ``/``. For Lambda functions, this is usually of the form ``/2015-03-31/functions/[FunctionARN]/invocations``. |
| `EnableSimpleResponses` | enable_simple_responses | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether a Lambda authorizer returns a response in a simple format. By default, a Lambda authorizer must return an IAM policy. If enabled, the Lambda authorizer can return a boolean value instead of an IAM policy. Supported only for HTTP APIs. To learn more, see [Working with Lambda authorizers for HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html). |
| `IdentitySource` | identity_source | `list` | optional, computed, provider-chosen |  | The identity source for which authorization is requested. |
| `IdentityValidationExpression` | identity_validation_expression | `string` | optional, computed, provider-chosen |  | This parameter is not used. |
| `JwtConfiguration` | jwt_configuration | `map` | optional, computed, provider-chosen |  | The ``JWTConfiguration`` property specifies the configuration of a JWT authorizer. Required for the ``JWT`` authorizer type. Supported only for HTTP APIs. |
| `Name` |  | `string` | required |  | The name of the authorizer. |

Supports update: yes

Discovery: supported (parent resource required)
