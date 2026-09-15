# aws.apigateway.stage

**CloudFormation type:** `AWS::ApiGateway::Stage`

The ``AWS::ApiGateway::Stage`` resource creates a stage for a deployment.

Region attribute: `region`

**Import ID:** `<region>/RestApiId|StageName` (AWS::ApiGateway::Stage)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessLogSetting` | access_log_setting | `map` | optional, computed, provider-chosen |  | The ``AccessLogSetting`` property type specifies settings for logging access in this stage. |
| `CacheClusterEnabled` | cache_cluster_enabled | `boolean` | optional, computed, provider-chosen |  |  |
| `CacheClusterSize` | cache_cluster_size | `string` | optional, computed, provider-chosen |  |  |
| `CanarySetting` | canary_setting | `map` | optional, computed, provider-chosen |  |  |
| `ClientCertificateId` | client_certificate_id | `string` | optional, computed, provider-chosen | aws.clientcertificate.ClientCertificateId |  |
| `DeploymentId` | deployment_id | `string` | optional, computed, provider-chosen | aws.apigateway.deployment.DeploymentId |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DocumentationVersion` | documentation_version | `string` | optional, computed, provider-chosen |  |  |
| `MethodSettings` | method_settings | `list` | optional, computed, provider-chosen |  |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |
| `StageName` | stage_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TracingEnabled` | tracing_enabled | `boolean` | optional, computed, provider-chosen |  |  |
| `Variables` |  | `map` | optional, computed, provider-chosen |  | A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: ``[A-Za-z0-9-._~:/?#&=,]+``. |

Supports update: yes

Discovery: supported (parent resource required)
