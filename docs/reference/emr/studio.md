# aws.studio

**CloudFormation type:** `AWS::EMR::Studio`

Resource schema for AWS::EMR::Studio

Region attribute: `region`

**Import ID:** `<region>/StudioId` (AWS::EMR::Studio)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the EMR Studio. |
| `AuthMode` | auth_mode | `string` | required, replaces on change |  | Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication. |
| `DefaultS3Location` | default_s3_location | `string` | required |  | The default Amazon S3 location to back up EMR Studio Workspaces and notebook files. A Studio user can select an alternative Amazon S3 location when creating a Workspace. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A detailed description of the Studio. |
| `EncryptionKeyArn` | encryption_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS KMS key identifier (ARN) used to encrypt AWS EMR Studio workspace and notebook files when backed up to AWS S3. |
| `EngineSecurityGroupId` | engine_security_group_id | `string` | required, replaces on change | aws.securitygroup.Id | The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by VpcId. |
| `IdcInstanceArn` | idc_instance_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the IAM Identity Center instance to create the Studio application. |
| `IdcUserAssignment` | idc_user_assignment | `string` | optional, computed, provider-chosen, replaces on change |  | Specifies whether IAM Identity Center user assignment is REQUIRED or OPTIONAL. If the value is set to REQUIRED, users must be explicitly assigned to the Studio application to access the Studio. |
| `IdpAuthUrl` | idp_auth_url | `string` | optional, computed, provider-chosen |  | Your identity provider's authentication endpoint. Amazon EMR Studio redirects federated users to this endpoint for authentication when logging in to a Studio with the Studio URL. |
| `IdpRelayStateParameterName` | idp_relay_state_parameter_name | `string` | optional, computed, provider-chosen |  | The name of relay state parameter for external Identity Provider. |
| `Name` |  | `string` | required |  | A descriptive name for the Amazon EMR Studio. |
| `ServiceRole` | service_role | `string` | required, replaces on change |  | The IAM role that will be assumed by the Amazon EMR Studio. The service role provides a way for Amazon EMR Studio to interoperate with other AWS services. |
| `StudioId` | studio_id | `string` | computed |  | The ID of the EMR Studio. |
| `SubnetIds` | subnet_ids | `list` | required | aws.subnet.SubnetId | A list of up to 5 subnet IDs to associate with the Studio. The subnets must belong to the VPC specified by VpcId. Studio users can create a Workspace in any of the specified subnets. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to associate with the Studio. Tags are user-defined key-value pairs that consist of a required key string with a maximum of 128 characters, and an optional value string with a maximum of 256 characters. |
| `TrustedIdentityPropagationEnabled` | trusted_identity_propagation_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  | A Boolean indicating whether to enable Trusted identity propagation for the Studio. The default value is false. |
| `Url` |  | `string` | computed |  | The unique Studio access URL. |
| `UserRole` | user_role | `string` | optional, computed, provider-chosen, replaces on change |  | The IAM user role that will be assumed by users and groups logged in to a Studio. The permissions attached to this IAM role can be scoped down for each user or group using session policies. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio. |
| `WorkspaceSecurityGroupId` | workspace_security_group_id | `string` | required, replaces on change | aws.securitygroup.Id | The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by VpcId. |

Supports update: yes

Discovery: supported
