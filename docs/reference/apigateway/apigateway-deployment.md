# aws.apigateway.deployment

**CloudFormation type:** `AWS::ApiGateway::Deployment`

The ``AWS::ApiGateway::Deployment`` resource deploys an API Gateway ``RestApi`` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.

Region attribute: `region`

**Import ID:** `<region>/DeploymentId|RestApiId` (AWS::ApiGateway::Deployment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeploymentCanarySettings` | deployment_canary_settings | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | The ``DeploymentCanarySettings`` property type specifies settings for the canary deployment. |
| `DeploymentId` | deployment_id | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |
| `StageDescription` | stage_description | `map` | optional, computed, provider-chosen, write-only |  | ``StageDescription`` is a property of the [AWS::ApiGateway::Deployment](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html) resource that configures a deployment stage. |
| `StageName` | stage_name | `string` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported (parent resource required)
