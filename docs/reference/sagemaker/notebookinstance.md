# aws.notebookinstance

**CloudFormation type:** `AWS::SageMaker::NotebookInstance`

Resource Type definition for AWS::SageMaker::NotebookInstance

Region attribute: `region`

**Import ID:** `<region>/NotebookInstanceArn` (AWS::SageMaker::NotebookInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceleratorTypes` | accelerator_types | `list` | optional, computed, provider-chosen |  | A list of Amazon Elastic Inference (EI) instance types to associate with the notebook instance. Currently, only one instance type can be associated with a notebook instance. |
| `AdditionalCodeRepositories` | additional_code_repositories | `list` | optional, computed, provider-chosen |  | An array of up to three Git repositories associated with the notebook instance. These can be either the names of Git repositories stored as resources in your account, or the URL of Git repositories in AWS CodeCommit or in any other Git repository. These repositories are cloned at the same level as the default repository of your notebook instance. |
| `DefaultCodeRepository` | default_code_repository | `string` | optional, computed, provider-chosen |  | The Git repository associated with the notebook instance as its default code repository. This can be either the name of a Git repository stored as a resource in your account, or the URL of a Git repository in AWS CodeCommit or in any other Git repository. When you open a notebook instance, it opens in the directory that contains this repository. |
| `DirectInternetAccess` | direct_internet_access | `string` | optional, computed, provider-chosen, replaces on change |  | Sets whether SageMaker AI provides internet access to the notebook instance. If you set this to Disabled this notebook instance is able to access resources only in your VPC, and is not be able to connect to SageMaker AI training and endpoint services unless you configure a NAT Gateway in your VPC. You can set the value of this parameter to Disabled only if you set a value for the SubnetId parameter. |
| `InstanceMetadataServiceConfiguration` | instance_metadata_service_configuration | `map` | optional, computed, provider-chosen |  | Information on the IMDS configuration of the notebook instance |
| `InstanceType` | instance_type | `string` | required |  | The type of ML compute instance to launch for the notebook instance. Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of a AWS Key Management Service key that SageMaker AI uses to encrypt data on the storage volume attached to your notebook instance. The KMS key you provide must be enabled. |
| `LifecycleConfigName` | lifecycle_config_name | `string` | optional, computed, provider-chosen |  | The name of a lifecycle configuration to associate with the notebook instance. |
| `NotebookInstanceArn` | notebook_instance_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the notebook instance. |
| `NotebookInstanceName` | notebook_instance_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the new notebook instance. |
| `PlatformIdentifier` | platform_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The platform identifier of the notebook instance runtime environment. The default value is notebook-al2023-v1. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | When you send any requests to AWS resources from the notebook instance, SageMaker AI assumes this role to perform tasks on your behalf. You must grant this role necessary permissions so SageMaker AI can perform these tasks. The policy must allow the SageMaker AI service principal (sagemaker.amazonaws.com) permissions to assume this role. To be able to pass this role to SageMaker AI, the caller of this API must have the iam:PassRole permission. |
| `RootAccess` | root_access | `string` | optional, computed, provider-chosen |  | Whether root access is enabled or disabled for users of the notebook instance. The default value is Enabled. Lifecycle configurations need root access to be able to set up a notebook instance. Because of this, lifecycle configurations associated with a notebook instance always run with root access even if you disable root access for users. |
| `SecurityGroupIds` | security_group_ids | `list` | optional, computed, provider-chosen, replaces on change | aws.securitygroup.Id | The VPC security group IDs, in the form sg-xxxxxxxx. The security groups must be for the same VPC as specified in the subnet. |
| `SubnetId` | subnet_id | `string` | optional, computed, provider-chosen, replaces on change | aws.subnet.SubnetId | The ID of the subnet in a VPC to which you would like to have a connectivity from your ML compute instance. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs to apply to this resource. |
| `VolumeSizeInGB` | volume_size_in_gb | `integer` | optional, computed, provider-chosen |  | The size, in GB, of the ML storage volume to attach to the notebook instance. The default value is 5 GB. Expect some interruption of service if this parameter is changed as CloudFormation stops a notebook instance and starts it up again to update it. |

Supports update: yes

Discovery: supported
