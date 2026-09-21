# aws.infrastructureconfiguration

**CloudFormation type:** `AWS::ImageBuilder::InfrastructureConfiguration`

Resource Type definition for AWS::ImageBuilder::InfrastructureConfiguration

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::InfrastructureConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the infrastructure configuration. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the infrastructure configuration. |
| `InstanceMetadataOptions` | instance_metadata_options | `map` | optional, computed, provider-chosen |  | The instance metadata option settings for the infrastructure configuration. |
| `InstanceProfileName` | instance_profile_name | `string` | required |  | The instance profile of the infrastructure configuration. |
| `InstanceTypes` | instance_types | `list` | optional, computed, provider-chosen |  | The instance types of the infrastructure configuration. |
| `KeyPair` | key_pair | `string` | optional, computed, provider-chosen |  | The EC2 key pair of the infrastructure configuration.. |
| `Logging` |  | `map` | optional, computed, provider-chosen |  | The logging configuration of the infrastructure configuration. |
| `Name` |  | `string` | required, replaces on change |  | The name of the infrastructure configuration. |
| `Placement` |  | `map` | optional, computed, provider-chosen |  | The placement options |
| `ResourceTags` | resource_tags | `map` | optional, computed, provider-chosen |  | The tags attached to the resource created by Image Builder. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen | aws.securitygroup.Id | The security group IDs of the infrastructure configuration. |
| `SnsTopicArn` | sns_topic_arn | `string` | optional, computed, provider-chosen |  | The SNS Topic Amazon Resource Name (ARN) of the infrastructure configuration. |
| `SubnetId` | subnet_id | `string` | optional, computed, provider-chosen | aws.subnet.SubnetId | The subnet ID of the infrastructure configuration. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags associated with the component. |
| `TerminateInstanceOnFailure` | terminate_instance_on_failure | `boolean` | optional, computed, provider-chosen |  | The terminate instance on failure configuration of the infrastructure configuration. |

Supports update: yes

Discovery: supported
