# aws.deploymentgroup

**CloudFormation type:** `AWS::CodeDeploy::DeploymentGroup`

Resource type definition for AWS::CodeDeploy::DeploymentGroup

Region attribute: `region`

**Import ID:** `<region>/ApplicationName|DeploymentGroupName` (AWS::CodeDeploy::DeploymentGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AlarmConfiguration` | alarm_configuration | `map` | optional, computed, provider-chosen |  | Configures CloudWatch alarms for an AWS CodeDeploy deployment group. |
| `ApplicationName` | application_name | `string` | required, replaces on change |  | The name of an existing CodeDeploy application to associate this deployment group with. |
| `AutoRollbackConfiguration` | auto_rollback_configuration | `map` | optional, computed, provider-chosen |  | Configures automatic rollback for an AWS CodeDeploy deployment group when a deployment is not completed successfully. |
| `AutoScalingGroups` | auto_scaling_groups | `list` | optional, computed, provider-chosen |  | A list of associated Auto Scaling groups that CodeDeploy automatically deploys revisions to when new instances are created. Duplicates are not allowed. |
| `BlueGreenDeploymentConfiguration` | blue_green_deployment_configuration | `map` | optional, computed, provider-chosen |  | Information about blue/green deployment options for a deployment group. |
| `Deployment` |  | `map` | optional, computed, provider-chosen |  | Specifies an AWS CodeDeploy application revision to be deployed to instances in the deployment group. If you specify an application revision, your target revision is deployed as soon as the provisioning process is complete. |
| `DeploymentConfigName` | deployment_config_name | `string` | optional, computed, provider-chosen |  | A deployment configuration name or a predefined configuration name. With predefined configurations, you can deploy application revisions to one instance at a time (CodeDeployDefault.OneAtATime), half of the instances at a time (CodeDeployDefault.HalfAtATime), or all the instances at once (CodeDeployDefault.AllAtOnce). |
| `DeploymentGroupName` | deployment_group_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the deployment group. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment group name. |
| `DeploymentStyle` | deployment_style | `map` | optional, computed, provider-chosen |  | Information about the type of deployment, either in-place or blue/green, you want to run and whether to route deployment traffic behind a load balancer. |
| `ECSServices` | ecs_services | `list` | optional, computed, provider-chosen |  | The target Amazon ECS services in the deployment group. This applies only to deployment groups that use the Amazon ECS compute platform. A target Amazon ECS service is specified as an Amazon ECS cluster and service name pair using the format <clustername>:<servicename>. |
| `Ec2TagFilters` | ec2_tag_filters | `list` | optional, computed, provider-chosen |  | The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group. CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed. You can specify EC2TagFilters or Ec2TagSet, but not both. |
| `Ec2TagSet` | ec2_tag_set | `map` | optional, computed, provider-chosen |  | Specifies information about groups of tags applied to Amazon EC2 instances. The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same template as EC2TagFilters. |
| `LoadBalancerInfo` | load_balancer_info | `map` | optional, computed, provider-chosen |  | Specifies information about the load balancer or target group used for an AWS CodeDeploy deployment group. For AWS CloudFormation to use the properties specified in LoadBalancerInfo, the DeploymentStyle.DeploymentOption property must be set to WITH_TRAFFIC_CONTROL. |
| `OnPremisesInstanceTagFilters` | on_premises_instance_tag_filters | `list` | optional, computed, provider-chosen |  | The on-premises instance tags already applied to on-premises instances that you want to include in the deployment group. CodeDeploy includes all on-premises instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed. You can specify OnPremisesInstanceTagFilters or OnPremisesInstanceTagSet, but not both. |
| `OnPremisesTagSet` | on_premises_tag_set | `map` | optional, computed, provider-chosen |  | Specifies a list containing other lists of on-premises instance tag groups. In order for an instance to be included in the deployment group, it must be identified by all the tag groups in the list. |
| `OutdatedInstancesStrategy` | outdated_instances_strategy | `string` | optional, computed, provider-chosen |  | Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision. If this option is set to UPDATE or is unspecified, CodeDeploy initiates one or more 'auto-update outdated instances' deployments to apply the deployed application revision to the new Amazon EC2 instances. If this option is set to IGNORE, CodeDeploy does not initiate a deployment to update the new Amazon EC2 instances. This may result in instances having different revisions. |
| `ServiceRoleArn` | service_role_arn | `string` | required | aws.role.Arn | A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to AWS services on your behalf. For more information, see 'Create a Service Role for AWS CodeDeploy' in the AWS CodeDeploy User Guide. |
| `Tags` |  | `map` | tags map |  | The metadata that you apply to CodeDeploy deployment groups to help you organize and categorize them. Each tag consists of a key and an optional value, both of which you define. |
| `TerminationHookEnabled` | termination_hook_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether the deployment group was configured to have CodeDeploy install a termination hook into an Auto Scaling group. |
| `TriggerConfigurations` | trigger_configurations | `list` | optional, computed, provider-chosen |  | Information about triggers associated with the deployment group. Duplicates are not allowed. |

Supports update: yes

Discovery: supported (parent resource required)
