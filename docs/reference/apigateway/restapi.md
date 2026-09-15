# aws.restapi

**CloudFormation type:** `AWS::ApiGateway::RestApi`

The ``AWS::ApiGateway::RestApi`` resource creates a REST API. For more information, see [restapi:create](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateRestApi.html) in the *Amazon API Gateway REST API Reference*.

Region attribute: `region`

**Import ID:** `<region>/RestApiId` (AWS::ApiGateway::RestApi)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiKeySourceType` | api_key_source_type | `string` | optional, computed, provider-chosen |  |  |
| `BinaryMediaTypes` | binary_media_types | `list` | optional, computed, provider-chosen |  |  |
| `Body` |  | `string` | optional, computed, provider-chosen, write-only |  | An OpenAPI specification that defines a set of RESTful APIs in JSON format. For YAML templates, you can also provide the specification in YAML format. |
| `BodyS3Location` | body_s3_location | `map` | optional, computed, provider-chosen, write-only |  | ``S3Location`` is a property of the [AWS::ApiGateway::RestApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html) resource that specifies the Amazon S3 location of a OpenAPI (formerly Swagger) file that defines a set of RESTful APIs in JSON or YAML. |
| `CloneFrom` | clone_from | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisableExecuteApiEndpoint` | disable_execute_api_endpoint | `boolean` | optional, computed, provider-chosen |  |  |
| `EndpointAccessMode` | endpoint_access_mode | `string` | optional, computed, provider-chosen |  |  |
| `EndpointConfiguration` | endpoint_configuration | `map` | optional, computed, provider-chosen |  | The ``EndpointConfiguration`` property type specifies the endpoint types and IP address types of a REST API. |
| `FailOnWarnings` | fail_on_warnings | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `MinimumCompressionSize` | minimum_compression_size | `integer` | optional, computed, provider-chosen |  |  |
| `Mode` |  | `string` | optional, computed, provider-chosen, write-only |  | This property applies only when you use OpenAPI to define your REST API. The ``Mode`` determines how API Gateway handles resource updates. |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the RestApi. A name is required if the REST API is not based on an OpenAPI specification. |
| `Parameters` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Policy` |  | `string` | optional, computed, provider-chosen |  | A policy document that contains the permissions for the ``RestApi`` resource. To set the ARN for the policy, use the ``!Join`` intrinsic function with ``""`` as delimiter and values of ``"execute-api:/"`` and ``"*"``. |
| `RestApiId` | rest_api_id | `string` | computed |  |  |
| `RootResourceId` | root_resource_id | `string` | computed |  |  |
| `SecurityPolicy` | security_policy | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Version` |  | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
