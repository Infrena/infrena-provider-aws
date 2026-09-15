# aws.apigatewayv2.api

**CloudFormation type:** `AWS::ApiGatewayV2::Api`

The ``AWS::ApiGatewayV2::Api`` resource creates an API. WebSocket APIs and HTTP APIs are supported. For more information about WebSocket APIs, see [About WebSocket APIs in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-overview.html) in the *API Gateway Developer Guide*. For more information about HTTP APIs, see [HTTP APIs](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api.html) in the *API Gateway Developer Guide.*

Region attribute: `region`

**Import ID:** `<region>/ApiId` (AWS::ApiGatewayV2::Api)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiEndpoint` | api_endpoint | `string` | computed |  |  |
| `ApiId` | api_id | `string` | computed |  |  |
| `ApiKeySelectionExpression` | api_key_selection_expression | `string` | optional, computed, provider-chosen |  | An API key selection expression. Supported only for WebSocket APIs. See [API Key Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-apikey-selection-expressions). |
| `BasePath` | base_path | `string` | optional, computed, provider-chosen, write-only |  | Specifies how to interpret the base path of the API during import. Valid values are ``ignore``, ``prepend``, and ``split``. The default value is ``ignore``. To learn more, see [Set the OpenAPI basePath Property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-import-api-basePath.html). Supported only for HTTP APIs. |
| `Body` |  | `map` | optional, computed, provider-chosen, write-only |  | The OpenAPI definition. Supported only for HTTP APIs. To import an HTTP API, you must specify a ``Body`` or ``BodyS3Location``. If you specify a ``Body`` or ``BodyS3Location``, don't specify CloudFormation resources such as ``AWS::ApiGatewayV2::Authorizer`` or ``AWS::ApiGatewayV2::Route``. API Gateway doesn't support the combination of OpenAPI and CloudFormation resources. |
| `BodyS3Location` | body_s3_location | `map` | optional, computed, provider-chosen, write-only |  | The ``BodyS3Location`` property specifies an S3 location from which to import an OpenAPI definition. Supported only for HTTP APIs. |
| `CorsConfiguration` | cors_configuration | `map` | optional, computed, provider-chosen |  | The ``Cors`` property specifies a CORS configuration for an API. Supported only for HTTP APIs. See [Configuring CORS](https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-cors.html) for more information. |
| `CredentialsArn` | credentials_arn | `string` | optional, computed, provider-chosen, write-only |  | This property is part of quick create. It specifies the credentials required for the integration, if any. For a Lambda integration, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify ``arn:aws:iam::*:user/*``. To use resource-based permissions on supported AWS services, specify ``null``. Currently, this property is not used for HTTP integrations. Supported only for HTTP APIs. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the API. |
| `DisableExecuteApiEndpoint` | disable_execute_api_endpoint | `boolean` | optional, computed, provider-chosen |  | Specifies whether clients can invoke your API by using the default ``execute-api`` endpoint. By default, clients can invoke your API with the default https://{api_id}.execute-api.{region}.amazonaws.com endpoint. To require that clients use a custom domain name to invoke your API, disable the default endpoint. |
| `DisableSchemaValidation` | disable_schema_validation | `boolean` | optional, computed, provider-chosen, write-only |  | Avoid validating models when creating a deployment. Supported only for WebSocket APIs. |
| `ExecuteApiArn` | execute_api_arn | `string` | computed |  |  |
| `FailOnWarnings` | fail_on_warnings | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether to rollback the API creation when a warning is encountered. By default, API creation continues if a warning is encountered. |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  | The IP address types that can invoke the API. Use ``ipv4`` to allow only IPv4 addresses to invoke your API, or use ``dualstack`` to allow both IPv4 and IPv6 addresses to invoke your API. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the API. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``. |
| `ProtocolType` | protocol_type | `string` | optional, computed, provider-chosen, replaces on change |  | The API protocol. Valid values are ``WEBSOCKET`` or ``HTTP``. Required unless you specify an OpenAPI definition for ``Body`` or ``S3BodyLocation``. |
| `RouteKey` | route_key | `string` | optional, computed, provider-chosen, write-only |  | This property is part of quick create. If you don't specify a ``routeKey``, a default route of ``$default`` is created. The ``$default`` route acts as a catch-all for any request made to your API, for a particular stage. The ``$default`` route key can't be modified. You can add routes after creating the API, and you can update the route keys of additional routes. Supported only for HTTP APIs. |
| `RouteSelectionExpression` | route_selection_expression | `string` | optional, computed, provider-chosen |  | The route selection expression for the API. For HTTP APIs, the ``routeSelectionExpression`` must be ``${request.method} ${request.path}``. If not provided, this will be the default for HTTP APIs. This property is required for WebSocket APIs. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The collection of tags. Each tag element is associated with a given resource. |
| `Target` |  | `string` | optional, computed, provider-chosen, write-only |  | This property is part of quick create. Quick create produces an API with an integration, a default catch-all route, and a default stage which is configured to automatically deploy changes. For HTTP integrations, specify a fully qualified URL. For Lambda integrations, specify a function ARN. The type of the integration will be HTTP_PROXY or AWS_PROXY, respectively. Supported only for HTTP APIs. |
| `Version` |  | `string` | optional, computed, provider-chosen |  | A version identifier for the API. |

Supports update: yes

Discovery: supported
