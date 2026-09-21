# aws.sagemaker.domain

**CloudFormation type:** `AWS::SageMaker::Domain`

Resource Type definition for AWS::SageMaker::Domain

Region attribute: `region`

**Import ID:** `<region>/DomainId` (AWS::SageMaker::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppNetworkAccessType` | app_network_access_type | `string` | optional, computed, provider-chosen |  | Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly. |
| `AppSecurityGroupManagement` | app_security_group_management | `string` | optional, computed, provider-chosen |  | The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided. |
| `AuthMode` | auth_mode | `string` | required, replaces on change |  | The mode of authentication that members use to access the domain. |
| `DefaultSpaceSettings` | default_space_settings | `map` | optional, computed, provider-chosen |  | A collection of settings that apply to spaces of Amazon SageMaker Studio. These settings are specified when the Create/Update Domain API is called. |
| `DefaultUserSettings` | default_user_settings | `map` | required |  | A collection of settings that apply to users of Amazon SageMaker Studio. These settings are specified when the CreateUserProfile API is called, and as DefaultUserSettings when the CreateDomain API is called. |
| `DomainArn` | domain_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the created domain. |
| `DomainId` | domain_id | `string` | computed |  | The domain name. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | A name for the domain. |
| `DomainSettings` | domain_settings | `map` | optional, computed, provider-chosen |  | A collection of Domain settings. |
| `HomeEfsFileSystemCreation` | home_efs_file_system_creation | `string` | optional, computed, provider-chosen |  | Indicates whether a home EFS file system is created for the domain. Set to Disabled to skip EFS creation and reduce domain creation time. |
| `HomeEfsFileSystemId` | home_efs_file_system_id | `string` | computed |  | The ID of the Amazon Elastic File System (EFS) managed by this Domain. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | SageMaker uses AWS KMS to encrypt the EFS volume attached to the domain with an AWS managed customer master key (CMK) by default. |
| `SecurityGroupIdForDomainBoundary` | security_group_id_for_domain_boundary | `string` | computed |  | The ID of the security group that authorizes traffic between the RSessionGateway apps and the RStudioServerPro app. |
| `SingleSignOnApplicationArn` | single_sign_on_application_arn | `string` | computed |  | The ARN of the application managed by SageMaker in IAM Identity Center. This value is only returned for domains created after October 1, 2023. |
| `SingleSignOnManagedApplicationInstanceId` | single_sign_on_managed_application_instance_id | `string` | computed |  | The SSO managed application instance ID. |
| `SubnetIds` | subnet_ids | `list` | optional, computed, provider-chosen | aws.subnet.SubnetId | The VPC subnets that Studio uses for communication. |
| `TagPropagation` | tag_propagation | `string` | optional, computed, provider-chosen |  | Indicates whether the tags added to Domain, User Profile and Space entity is propagated to all SageMaker resources. |
| `Tags` |  | `map` | tags map |  | A list of tags to apply to the user profile. |
| `Url` |  | `string` | computed |  | The URL to the created domain. |
| `VpcId` | vpc_id | `string` | optional, computed, provider-chosen | aws.vpc.VpcId | The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication. |

Supports update: yes

Discovery: supported
