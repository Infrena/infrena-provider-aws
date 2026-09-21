# aws.appsync.datasource

**CloudFormation type:** `AWS::AppSync::DataSource`

Resource Type definition for AWS::AppSync::DataSource

Region attribute: `region`

**Import ID:** `<region>/DataSourceArn` (AWS::AppSync::DataSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.appsync.api.ApiId | Unique AWS AppSync GraphQL API identifier where this data source will be created. |
| `DataSourceArn` | data_source_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the API key, such as arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/datasources/datasourcename. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the data source. |
| `DynamoDBConfig` | dynamo_db_config | `map` | optional, computed, provider-chosen |  | AWS Region and TableName for an Amazon DynamoDB table in your account. |
| `ElasticsearchConfig` | elasticsearch_config | `map` | optional, computed, provider-chosen |  | AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account. |
| `EventBridgeConfig` | event_bridge_config | `map` | optional, computed, provider-chosen |  | ARN for the EventBridge bus. |
| `HttpConfig` | http_config | `map` | optional, computed, provider-chosen |  | Endpoints for an HTTP data source. |
| `LambdaConfig` | lambda_config | `map` | optional, computed, provider-chosen |  | An ARN of a Lambda function in valid ARN format. This can be the ARN of a Lambda function that exists in the current account or in another account. |
| `MetricsConfig` | metrics_config | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  | Friendly name for you to identify your AppSync data source after creation. |
| `OpenSearchServiceConfig` | open_search_service_config | `map` | optional, computed, provider-chosen |  | AWS Region and Endpoints for an Amazon OpenSearch Service domain in your account. |
| `RelationalDatabaseConfig` | relational_database_config | `map` | optional, computed, provider-chosen |  | Relational Database configuration of the relational database data source. |
| `ServiceRoleArn` | service_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The AWS Identity and Access Management service role ARN for the data source. The system assumes this role when accessing the data source. |
| `Type` | type_value | `string` | required |  | The type of the data source. |

Supports update: yes

Discovery: supported (parent resource required)
