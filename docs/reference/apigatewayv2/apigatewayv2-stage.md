# aws.apigatewayv2.stage

**CloudFormation type:** `AWS::ApiGatewayV2::Stage`

Resource Type definition for AWS::ApiGatewayV2::Stage

Region attribute: `region`

**Import ID:** `<region>/ApiId|StageName` (AWS::ApiGatewayV2::Stage)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessLogSettings` | access_log_settings | `map` | optional, computed, provider-chosen |  | Settings for logging access in this stage. |
| `ApiId` | api_id | `string` | required, replaces on change | aws.apigatewayv2.api.ApiId | The API identifier. |
| `AutoDeploy` | auto_deploy | `boolean` | optional, computed, provider-chosen |  | Specifies whether updates to an API automatically trigger a new deployment. The default value is false. |
| `ClientCertificateId` | client_certificate_id | `string` | optional, computed, provider-chosen |  | The identifier of a client certificate for a Stage. Supported only for WebSocket APIs. |
| `DefaultRouteSettings` | default_route_settings | `map` | optional, computed, provider-chosen |  | The default route settings for the stage. |
| `DeploymentId` | deployment_id | `string` | optional, computed, provider-chosen | aws.apigatewayv2.deployment.DeploymentId | The deployment identifier for the API stage. Can't be updated if autoDeploy is enabled. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description for the API stage. |
| `RouteSettings` | route_settings | `map` | optional, computed, provider-chosen |  | Route settings for the stage. |
| `StageName` | stage_name | `string` | required, replaces on change |  | The stage name. Stage names can contain only alphanumeric characters, hyphens, and underscores, or be $default. Maximum length is 128 characters. |
| `StageVariables` | stage_variables | `map` | optional, computed, provider-chosen |  | A map that defines the stage variables for a Stage. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The collection of tags. Each tag element is associated with a given resource. |

Supports update: yes

Discovery: supported (parent resource required)
