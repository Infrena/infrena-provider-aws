# aws.apigatewayv2.deployment

**CloudFormation type:** `AWS::ApiGatewayV2::Deployment`

The ``AWS::ApiGatewayV2::Deployment`` resource creates a deployment for an API.

Region attribute: `region`

**Import ID:** `<region>/ApiId|DeploymentId` (AWS::ApiGatewayV2::Deployment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.apigatewayv2.api.ApiId | The API identifier. |
| `DeploymentId` | deployment_id | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description for the deployment resource. |
| `StageName` | stage_name | `string` | optional, computed, provider-chosen, write-only |  | The name of an existing stage to associate with the deployment. |

Supports update: yes

Discovery: supported (parent resource required)
