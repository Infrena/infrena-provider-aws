# aws.graphqlapi

**CloudFormation type:** `AWS::AppSync::GraphQLApi`

Resource Type definition for AWS::AppSync::GraphQLApi

Region attribute: `region`

**Import ID:** `<region>/ApiId` (AWS::AppSync::GraphQLApi)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalAuthenticationProviders` | additional_authentication_providers | `list` | optional, computed, provider-chosen |  | A list of additional authentication providers for the GraphqlApi API. |
| `ApiId` | api_id | `string` | computed |  | Unique AWS AppSync GraphQL API identifier. |
| `ApiType` | api_type | `string` | optional, computed, provider-chosen |  | The value that indicates whether the GraphQL API is a standard API (GRAPHQL) or merged API (MERGED). |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the API key |
| `AuthenticationType` | authentication_type | `string` | required |  | Security configuration for your GraphQL API |
| `EnhancedMetricsConfig` | enhanced_metrics_config | `map` | optional, computed, provider-chosen |  | Enables and controls the enhanced metrics feature. Enhanced metrics emit granular data on API usage and performance such as AppSync request and error counts, latency, and cache hits/misses. All enhanced metric data is sent to your CloudWatch account, and you can configure the types of data that will be sent. |
| `EnvironmentVariables` | environment_variables | `map` | optional, computed, provider-chosen |  | A map containing the list of resources with their properties and environment variables. |
| `GraphQLDns` | graph_ql_dns | `string` | computed |  | The fully qualified domain name (FQDN) of the endpoint URL of your GraphQL API. |
| `GraphQLEndpointArn` | graph_ql_endpoint_arn | `string` | computed |  | The GraphQL endpoint ARN. |
| `GraphQLUrl` | graph_ql_url | `string` | computed |  | The Endpoint URL of your GraphQL API. |
| `IntrospectionConfig` | introspection_config | `string` | optional, computed, provider-chosen |  | Sets the value of the GraphQL API to enable (ENABLED) or disable (DISABLED) introspection. If no value is provided, the introspection configuration will be set to ENABLED by default. This field will produce an error if the operation attempts to use the introspection feature while this field is disabled. |
| `LambdaAuthorizerConfig` | lambda_authorizer_config | `map` | optional, computed, provider-chosen |  | A LambdaAuthorizerConfig holds configuration on how to authorize AWS AppSync API access when using the AWS_LAMBDA authorizer mode. Be aware that an AWS AppSync API may have only one Lambda authorizer configured at a time. |
| `LogConfig` | log_config | `map` | optional, computed, provider-chosen |  | The Amazon CloudWatch Logs configuration. |
| `MergedApiExecutionRoleArn` | merged_api_execution_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The AWS Identity and Access Management service role ARN for a merged API. |
| `Name` |  | `string` | required |  | The API name |
| `OpenIDConnectConfig` | open_id_connect_config | `map` | optional, computed, provider-chosen |  | The OpenID Connect configuration. |
| `OwnerContact` | owner_contact | `string` | optional, computed, provider-chosen |  | The owner contact information for an API resource. |
| `QueryDepthLimit` | query_depth_limit | `integer` | optional, computed, provider-chosen |  | The maximum depth a query can have in a single request. Depth refers to the amount of nested levels allowed in the body of query. |
| `RealtimeDns` | realtime_dns | `string` | computed |  | The fully qualified domain name (FQDN) of the real-time endpoint URL of your GraphQL API. |
| `RealtimeUrl` | realtime_url | `string` | computed |  | The GraphQL API real-time endpoint URL. |
| `ResolverCountLimit` | resolver_count_limit | `integer` | optional, computed, provider-chosen |  | The maximum number of resolvers that can be invoked in a single request. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this GraphQL API. |
| `UserPoolConfig` | user_pool_config | `map` | optional, computed, provider-chosen |  | Optional authorization configuration for using Amazon Cognito user pools with your GraphQL endpoint. |
| `Visibility` |  | `string` | optional, computed, provider-chosen |  | Sets the scope of the GraphQL API to public (GLOBAL) or private (PRIVATE). By default, the scope is set to Global if no value is provided. |
| `XrayEnabled` | xray_enabled | `boolean` | optional, computed, provider-chosen |  | A flag indicating whether to use AWS X-Ray tracing for this GraphqlApi. |

Supports update: yes

Discovery: supported
