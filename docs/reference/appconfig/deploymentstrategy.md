# aws.deploymentstrategy

**CloudFormation type:** `AWS::AppConfig::DeploymentStrategy`

Resource Type definition for AWS::AppConfig::DeploymentStrategy

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::AppConfig::DeploymentStrategy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeploymentDurationInMinutes` | deployment_duration_in_minutes | `float` | required |  | Total amount of time for a deployment to last. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the deployment strategy. |
| `FinalBakeTimeInMinutes` | final_bake_time_in_minutes | `float` | optional, computed, provider-chosen |  | Specifies the amount of time AWS AppConfig monitors for Amazon CloudWatch alarms after the configuration has been deployed to 100% of its targets, before considering the deployment to be complete. If an alarm is triggered during this time, AWS AppConfig rolls back the deployment. You must configure permissions for AWS AppConfig to roll back based on CloudWatch alarms. For more information, see Configuring permissions for rollback based on Amazon CloudWatch alarms in the AWS AppConfig User Guide. |
| `GrowthFactor` | growth_factor | `float` | required |  | The percentage of targets to receive a deployed configuration during each interval. |
| `GrowthType` | growth_type | `string` | optional, computed, provider-chosen |  | The algorithm used to define how percentage grows over time. AWS AppConfig supports the following growth types: |
| `Id` |  | `string` | computed |  | The deployment strategy ID. |
| `Name` |  | `string` | required, replaces on change |  | A name for the deployment strategy. |
| `ReplicateTo` | replicate_to | `string` | required, replaces on change |  | Save the deployment strategy to a Systems Manager (SSM) document. |
| `Tags` |  | `map` | tags map |  | Assigns metadata to an AWS AppConfig resource. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. You can specify a maximum of 50 tags for a resource. |

Supports update: yes

Discovery: supported
