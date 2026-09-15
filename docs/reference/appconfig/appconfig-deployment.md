# aws.appconfig.deployment

**CloudFormation type:** `AWS::AppConfig::Deployment`

Resource Type definition for AWS::AppConfig::Deployment

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|EnvironmentId|DeploymentNumber` (AWS::AppConfig::Deployment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.appconfig.application.ApplicationId | The application ID. |
| `ConfigurationProfileId` | configuration_profile_id | `string` | required, replaces on change | aws.configurationprofile.ConfigurationProfileId | The configuration profile ID. |
| `ConfigurationVersion` | configuration_version | `string` | required, replaces on change |  | The configuration version to deploy. If deploying an AWS AppConfig hosted configuration version, you can specify either the version number or version label. For all other configurations, you must specify the version number. |
| `DeploymentNumber` | deployment_number | `string` | computed |  | The sequence number of the deployment. |
| `DeploymentStrategyId` | deployment_strategy_id | `string` | required, replaces on change | aws.deploymentstrategy.Id | The deployment strategy ID. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description of the deployment. |
| `DynamicExtensionParameters` | dynamic_extension_parameters | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `EnvironmentId` | environment_id | `string` | required, replaces on change | aws.appconfig.environment.EnvironmentId | The environment ID. |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS Key Management Service key identifier (key ID, key alias, or key ARN) provided when the resource was created or updated. |
| `State` |  | `string` | computed |  | The state of the deployment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: no

Discovery: supported (parent resource required)
