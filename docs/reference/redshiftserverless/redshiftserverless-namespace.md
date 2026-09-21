# aws.redshiftserverless.namespace

**CloudFormation type:** `AWS::RedshiftServerless::Namespace`

Definition of AWS::RedshiftServerless::Namespace Resource Type

Region attribute: `region`

**Import ID:** `<region>/NamespaceName` (AWS::RedshiftServerless::Namespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdminPasswordSecretKmsKeyId` | admin_password_secret_kms_key_id | `string` | optional, computed, provider-chosen |  | The ID of the AWS Key Management Service (KMS) key used to encrypt and store the namespace's admin credentials secret. You can only use this parameter if manageAdminPassword is true. |
| `AdminUserPassword` | admin_user_password | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The password associated with the admin user for the namespace that is being created. Password must be at least 8 characters in length, should be any printable ASCII character. Must contain at least one lowercase letter, one uppercase letter and one decimal digit. You can't use adminUserPassword if manageAdminPassword is true. |
| `AdminUsername` | admin_username | `string` | optional, computed, provider-chosen |  | The user name associated with the admin user for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet. |
| `DbName` | db_name | `string` | optional, computed, provider-chosen |  | The database name associated for the namespace that is being created. Only alphanumeric characters and underscores are allowed. It should start with an alphabet. |
| `DefaultIamRoleArn` | default_iam_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The default IAM role ARN for the namespace that is being created. |
| `FinalSnapshotName` | final_snapshot_name | `string` | optional, computed, provider-chosen, write-only |  | The name of the namespace the source snapshot was created from. Please specify the name if needed before deleting namespace |
| `FinalSnapshotRetentionPeriod` | final_snapshot_retention_period | `integer` | optional, computed, provider-chosen, write-only |  | The number of days to retain automated snapshot in the destination region after they are copied from the source region. If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3,653. |
| `IamRoles` | iam_roles | `list` | optional, computed, provider-chosen |  | A list of AWS Identity and Access Management (IAM) roles that can be used by the namespace to access other AWS services. You must supply the IAM roles in their Amazon Resource Name (ARN) format. The Default role limit for each request is 10. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | The AWS Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the namespace. |
| `LogExports` | log_exports | `list` | optional, computed, provider-chosen |  | The collection of log types to be exported provided by the customer. Should only be one of the three supported log types: userlog, useractivitylog and connectionlog |
| `ManageAdminPassword` | manage_admin_password | `boolean` | optional, computed, provider-chosen, write-only |  | If true, Amazon Redshift uses AWS Secrets Manager to manage the namespace's admin credentials. You can't use adminUserPassword if manageAdminPassword is true. If manageAdminPassword is false or not set, Amazon Redshift uses adminUserPassword for the admin user account's password. |
| `Namespace` |  | `map` | computed |  | Definition of Namespace resource. |
| `NamespaceName` | namespace_name | `string` | required, replaces on change |  | A unique identifier for the namespace. You use this identifier to refer to the namespace for any subsequent namespace operations such as deleting or modifying. All alphabetical characters must be lower case. Namespace name should be unique for all namespaces within an AWS account. |
| `NamespaceResourcePolicy` | namespace_resource_policy | `map` | optional, computed, provider-chosen |  | The resource policy document that will be attached to the namespace. |
| `RedshiftIdcApplicationArn` | redshift_idc_application_arn | `string` | optional, computed, provider-chosen, write-only |  | The ARN for the Redshift application that integrates with IAM Identity Center. |
| `SnapshotCopyConfigurations` | snapshot_copy_configurations | `list` | optional, computed, provider-chosen |  | The snapshot copy configurations for the namespace. |
| `Tags` |  | `map` | tags map |  | The list of tags for the namespace. |

Supports update: yes

Discovery: supported
