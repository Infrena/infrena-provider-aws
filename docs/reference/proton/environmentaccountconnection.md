# aws.environmentaccountconnection

**CloudFormation type:** `AWS::Proton::EnvironmentAccountConnection`

Resource Schema describing various properties for AWS Proton Environment Account Connections resources.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Proton::EnvironmentAccountConnection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the environment account connection. |
| `CodebuildRoleArn` | codebuild_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Amazon Resource Name (ARN) of an IAM service role in the environment account. AWS Proton uses this role to provision infrastructure resources using CodeBuild-based provisioning in the associated environment account. |
| `ComponentRoleArn` | component_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM service role that AWS Proton uses when provisioning directly defined components in the associated environment account. It determines the scope of infrastructure that a component can provision in the account. |
| `EnvironmentAccountId` | environment_account_id | `string` | optional, computed, provider-chosen |  | The environment account that's connected to the environment account connection. |
| `EnvironmentName` | environment_name | `string` | optional, computed, provider-chosen |  | The name of the AWS Proton environment that's created in the associated management account. |
| `Id` |  | `string` | computed |  | The ID of the environment account connection. |
| `ManagementAccountId` | management_account_id | `string` | optional, computed, provider-chosen |  | The ID of the management account that accepts or rejects the environment account connection. You create an manage the AWS Proton environment in this account. If the management account accepts the environment account connection, AWS Proton can use the associated IAM role to provision environment infrastructure resources in the associated environment account. |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM service role that's created in the environment account. AWS Proton uses this role to provision infrastructure resources in the associated environment account. |
| `Status` |  | `string` | computed |  | The status of the environment account connection. |
| `Tags` |  | `map` | tags map |  | <p>An optional list of metadata items that you can associate with the Proton environment account connection. A tag is a key-value pair.</p> |

Supports update: yes

Discovery: supported
