# aws.hookversion

**CloudFormation type:** `AWS::CloudFormation::HookVersion`

Publishes new or first hook version to AWS CloudFormation Registry.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CloudFormation::HookVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the type, here the HookVersion. This is used to uniquely identify a HookVersion resource |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials. |
| `IsDefaultVersion` | is_default_version | `boolean` | computed |  | Indicates if this type version is the current default version |
| `LoggingConfig` | logging_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies logging configuration information for a type. |
| `SchemaHandlerPackage` | schema_handler_package | `string` | required, replaces on change, write-only |  | A url to the S3 bucket containing the schema handler package that contains the schema, event handlers, and associated files for the type you want to register. |
| `TypeArn` | type_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the type without the versionID. |
| `TypeName` | type_name | `string` | required, replaces on change |  | The name of the type being registered. |
| `VersionId` | version_id | `string` | computed |  | The ID of the version of the type represented by this hook instance. |
| `Visibility` |  | `string` | computed |  | The scope at which the type is visible and usable in CloudFormation operations. |

Supports update: no

Discovery: supported
