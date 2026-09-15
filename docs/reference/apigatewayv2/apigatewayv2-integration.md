# aws.apigatewayv2.integration

**CloudFormation type:** `AWS::ApiGatewayV2::Integration`

Resource Type definition for AWS::ApiGatewayV2::Integration

Region attribute: `region`

**Import ID:** `<region>/ApiId|IntegrationId` (AWS::ApiGatewayV2::Integration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.apigatewayv2.api.ApiId | The API identifier. |
| `ConnectionId` | connection_id | `string` | optional, computed, provider-chosen |  | The ID of the VPC link for a private integration. Supported only for HTTP APIs. |
| `ConnectionType` | connection_type | `string` | optional, computed, provider-chosen |  | The type of the network connection to the integration endpoint. Specify INTERNET for connections through the public routable internet or VPC_LINK for private connections between API Gateway and resources in a VPC. The default value is INTERNET. |
| `ContentHandlingStrategy` | content_handling_strategy | `string` | optional, computed, provider-chosen |  | Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are CONVERT_TO_BINARY and CONVERT_TO_TEXT. |
| `CredentialsArn` | credentials_arn | `string` | optional, computed, provider-chosen |  | Specifies the credentials required for the integration, if any. For AWS integrations, three options are available. To specify an IAM Role for API Gateway to assume, use the role's Amazon Resource Name (ARN). To require that the caller's identity be passed through from the request, specify the string arn:aws:iam::*:user/*. To use resource-based permissions on supported AWS services, don't specify this parameter. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the integration. |
| `IntegrationId` | integration_id | `string` | computed |  | The integration ID. |
| `IntegrationMethod` | integration_method | `string` | optional, computed, provider-chosen |  | Specifies the integration's HTTP method type. |
| `IntegrationSubtype` | integration_subtype | `string` | optional, computed, provider-chosen |  | Supported only for HTTP API AWS_PROXY integrations. Specifies the AWS service action to invoke. |
| `IntegrationType` | integration_type | `string` | required |  | The integration type of an integration. |
| `IntegrationUri` | integration_uri | `string` | optional, computed, provider-chosen |  | For a Lambda integration, specify the URI of a Lambda function. For an HTTP integration, specify a fully-qualified URL. For an HTTP API private integration, specify the ARN of an Application Load Balancer listener, Network Load Balancer listener, or AWS Cloud Map service. |
| `PassthroughBehavior` | passthrough_behavior | `string` | optional, computed, provider-chosen |  | Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource. There are three valid values: WHEN_NO_MATCH, WHEN_NO_TEMPLATES, and NEVER. Supported only for WebSocket APIs. |
| `PayloadFormatVersion` | payload_format_version | `string` | optional, computed, provider-chosen |  | Specifies the format of the payload sent to an integration. Required for HTTP APIs. For HTTP APIs, supported values for Lambda proxy integrations are 1.0 and 2.0 For all other integrations, 1.0 is the only supported value. |
| `RequestParameters` | request_parameters | `map` | optional, computed, provider-chosen |  | A key-value map specifying parameters. |
| `RequestTemplates` | request_templates | `map` | optional, computed, provider-chosen |  | A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client. |
| `ResponseParameters` | response_parameters | `map` | optional, computed, provider-chosen |  | Parameters that transform the HTTP response from a backend integration before returning the response to clients. Supported only for HTTP APIs. |
| `TemplateSelectionExpression` | template_selection_expression | `string` | optional, computed, provider-chosen |  | The template selection expression for the integration. Supported only for WebSocket APIs. |
| `TimeoutInMillis` | timeout_in_millis | `integer` | optional, computed, provider-chosen |  | Custom timeout between 50 and 29000 milliseconds for WebSocket APIs and between 50 and 30000 milliseconds for HTTP APIs. The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs. |
| `TlsConfig` | tls_config | `map` | optional, computed, provider-chosen |  | The TlsConfig property specifies the TLS configuration for a private integration. Supported only for HTTP APIs. |

Supports update: yes

Discovery: supported (parent resource required)
