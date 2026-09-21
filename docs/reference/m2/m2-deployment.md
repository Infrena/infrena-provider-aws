# aws.m2.deployment

**CloudFormation type:** `AWS::M2::Deployment`

Represents a deployment resource of an AWS Mainframe Modernization (M2) application to a specified environment

Region attribute: `region`

**Import ID:** `<region>/ApplicationId` (AWS::M2::Deployment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.m2.application.ApplicationId | The application ID. |
| `ApplicationVersion` | application_version | `integer` | required |  | The version number of the application to deploy |
| `DeploymentId` | deployment_id | `string` | computed |  | The deployment ID. |
| `EnvironmentId` | environment_id | `string` | required, replaces on change | aws.m2.environment.EnvironmentId | The environment ID. |
| `Status` |  | `string` | computed |  | The status of the deployment. |

Supports update: yes

Discovery: supported (parent resource required)
