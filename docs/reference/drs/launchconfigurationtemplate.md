# aws.launchconfigurationtemplate

**CloudFormation type:** `AWS::DRS::LaunchConfigurationTemplate`

Account level Launch Configuration Template for AWS Elastic Disaster Recovery.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::DRS::LaunchConfigurationTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | ARN of the Launch Configuration Template. |
| `CopyPrivateIp` | copy_private_ip | `boolean` | optional, computed, provider-chosen |  | Copy private IP. |
| `CopyTags` | copy_tags | `boolean` | optional, computed, provider-chosen |  | Copy tags. |
| `ExportBucketArn` | export_bucket_arn | `string` | optional, computed, provider-chosen |  | S3 bucket ARN to export Source Network templates. |
| `LaunchConfigurationTemplateID` | launch_configuration_template_id | `string` | computed |  | ID of the Launch Configuration Template. |
| `LaunchDisposition` | launch_disposition | `string` | optional, computed, provider-chosen |  | Launch disposition. |
| `LaunchIntoSourceInstance` | launch_into_source_instance | `boolean` | optional, computed, provider-chosen |  | DRS will set the 'launch into instance ID' of any source server when performing a drill, recovery or failback to the previous region or availability zone, using the instance ID of the source instance. |
| `Licensing` |  | `map` | optional, computed, provider-chosen |  | Configuration of a machine's license. |
| `PostLaunchEnabled` | post_launch_enabled | `boolean` | optional, computed, provider-chosen |  | Whether we want to activate post-launch actions. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | A set of tags associated with the Launch Configuration Template. |
| `TargetInstanceTypeRightSizingMethod` | target_instance_type_right_sizing_method | `string` | optional, computed, provider-chosen |  | Target instance type right-sizing method. |

Supports update: yes

Discovery: supported
