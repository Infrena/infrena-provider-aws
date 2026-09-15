# aws.resolver

**CloudFormation type:** `AWS::AppSync::Resolver`

The ``AWS::AppSync::Resolver`` resource defines the logical GraphQL resolver that you attach to fields in a schema. Request and response templates for resolvers are written in Apache Velocity Template Language (VTL) format. For more information about resolvers, see [Resolver Mapping Template Reference](https://docs.aws.amazon.com/appsync/latest/devguide/resolver-mapping-template-reference.html).

Region attribute: `region`

**Import ID:** `<region>/ResolverArn` (AWS::AppSync::Resolver)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.appsync.api.ApiId | The APSYlong GraphQL API to which you want to attach this resolver. |
| `CachingConfig` | caching_config | `map` | optional, computed, provider-chosen |  | The caching configuration for a resolver that has caching activated. |
| `Code` |  | `string` | optional, computed, provider-chosen |  | The ``resolver`` code that contains the request and response functions. When code is used, the ``runtime`` is required. The runtime value must be ``APPSYNC_JS``. |
| `CodeS3Location` | code_s3_location | `string` | optional, computed, provider-chosen, write-only |  | The Amazon S3 endpoint. |
| `DataSourceName` | data_source_name | `string` | optional, computed, provider-chosen |  | The resolver data source name. |
| `FieldName` | field_name | `string` | required, replaces on change |  | The GraphQL field on a type that invokes the resolver. |
| `Kind` |  | `string` | optional, computed, provider-chosen |  | The resolver type. |
| `MaxBatchSize` | max_batch_size | `integer` | optional, computed, provider-chosen |  | The maximum number of resolver request inputs that will be sent to a single LAMlong function in a ``BatchInvoke`` operation. |
| `MetricsConfig` | metrics_config | `string` | optional, computed, provider-chosen |  | Enables or disables enhanced resolver metrics for specified resolvers. Note that ``MetricsConfig`` won't be used unless the ``resolverLevelMetricsBehavior`` value is set to ``PER_RESOLVER_METRICS``. If the ``resolverLevelMetricsBehavior`` is set to ``FULL_REQUEST_RESOLVER_METRICS`` instead, ``MetricsConfig`` will be ignored. However, you can still set its value. |
| `PipelineConfig` | pipeline_config | `map` | optional, computed, provider-chosen |  | Use the ``PipelineConfig`` property type to specify ``PipelineConfig`` for an APSYlong resolver. |
| `RequestMappingTemplate` | request_mapping_template | `string` | optional, computed, provider-chosen |  | The request mapping template. |
| `RequestMappingTemplateS3Location` | request_mapping_template_s3_location | `string` | optional, computed, provider-chosen, write-only |  | The location of a request mapping template in an S3 bucket. Use this if you want to provision with a template file in S3 rather than embedding it in your CFNshort template. |
| `ResolverArn` | resolver_arn | `string` | computed |  |  |
| `ResponseMappingTemplate` | response_mapping_template | `string` | optional, computed, provider-chosen |  | The response mapping template. |
| `ResponseMappingTemplateS3Location` | response_mapping_template_s3_location | `string` | optional, computed, provider-chosen, write-only |  | The location of a response mapping template in an S3 bucket. Use this if you want to provision with a template file in S3 rather than embedding it in your CFNshort template. |
| `Runtime` |  | `map` | optional, computed, provider-chosen |  | Describes a runtime used by an APSYlong resolver or APSYlong function. Specifies the name and version of the runtime to use. Note that if a runtime is specified, code must also be specified. |
| `SyncConfig` | sync_config | `map` | optional, computed, provider-chosen |  | Describes a Sync configuration for a resolver. |
| `TypeName` | type_name | `string` | required, replaces on change |  | The GraphQL type that invokes this resolver. |

Supports update: yes

Discovery: supported (parent resource required)
