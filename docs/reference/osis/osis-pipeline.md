# aws.osis.pipeline

**CloudFormation type:** `AWS::OSIS::Pipeline`

An OpenSearch Ingestion Service Data Prepper pipeline running Data Prepper.

Region attribute: `region`

**Import ID:** `<region>/PipelineArn` (AWS::OSIS::Pipeline)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BufferOptions` | buffer_options | `map` | optional, computed, provider-chosen |  | Key-value pairs to configure buffering. |
| `EncryptionAtRestOptions` | encryption_at_rest_options | `map` | optional, computed, provider-chosen |  | Key-value pairs to configure encryption at rest. |
| `IngestEndpointUrls` | ingest_endpoint_urls | `list` | computed |  | A list of endpoints that can be used for ingesting data into a pipeline |
| `LogPublishingOptions` | log_publishing_options | `map` | optional, computed, provider-chosen |  | Key-value pairs to configure log publishing. |
| `MaxUnits` | max_units | `integer` | required |  | The maximum pipeline capacity, in Ingestion OpenSearch Compute Units (OCUs). |
| `MinUnits` | min_units | `integer` | required |  | The minimum pipeline capacity, in Ingestion OpenSearch Compute Units (OCUs). |
| `PipelineArn` | pipeline_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the pipeline. |
| `PipelineConfigurationBody` | pipeline_configuration_body | `string` | required |  | The Data Prepper pipeline configuration. |
| `PipelineName` | pipeline_name | `string` | required, replaces on change |  | Name of the OpenSearch Ingestion Service pipeline to create. Pipeline names are unique across the pipelines owned by an account within an AWS Region. |
| `PipelineRoleArn` | pipeline_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Pipeline Role (ARN) for the pipeline. |
| `ResourcePolicy` | resource_policy | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `VpcEndpointService` | vpc_endpoint_service | `string` | computed |  | The VPC endpoint service name for the pipeline. |
| `VpcEndpoints` | vpc_endpoints | `list` | computed |  | The VPC interface endpoints that have access to the pipeline. |
| `VpcOptions` | vpc_options | `map` | optional, computed, provider-chosen, write-only |  | Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion Service creates the pipeline with a public endpoint. |

Supports update: yes

Discovery: supported
