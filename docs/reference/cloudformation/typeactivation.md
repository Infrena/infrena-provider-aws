# aws.typeactivation

**CloudFormation type:** `AWS::CloudFormation::TypeActivation`

Enable a resource that has been published in the CloudFormation Registry.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CloudFormation::TypeActivation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the extension. |
| `AutoUpdate` | auto_update | `boolean` | optional, computed, provider-chosen, write-only |  | Whether to automatically update the extension in this account and region when a new minor version is published by the extension publisher. Major versions released by the publisher must be manually updated. |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, write-only | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM execution role to use to register the type. If your resource type calls AWS APIs in any of its handlers, you must create an IAM execution role that includes the necessary permissions to call those AWS APIs, and provision that execution role in your account. CloudFormation then assumes that execution role to provide your resource type with the appropriate credentials. |
| `LoggingConfig` | logging_config | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies logging configuration information for a type. |
| `MajorVersion` | major_version | `string` | optional, computed, provider-chosen, write-only |  | The Major Version of the type you want to enable |
| `PublicTypeArn` | public_type_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Number (ARN) assigned to the public extension upon publication |
| `PublisherId` | publisher_id | `string` | optional, computed, provider-chosen | aws.publisher.PublisherId | The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region. |
| `Type` | type_value | `string` | optional, computed, provider-chosen, write-only |  | The kind of extension |
| `TypeName` | type_name | `string` | optional, computed, provider-chosen |  | The name of the type being registered. |
| `TypeNameAlias` | type_name_alias | `string` | optional, computed, provider-chosen |  | An alias to assign to the public extension in this account and region. If you specify an alias for the extension, you must then use the alias to refer to the extension in your templates. |
| `VersionBump` | version_bump | `string` | optional, computed, provider-chosen, write-only |  | Manually updates a previously-enabled type to a new major or minor version, if available. You can also use this parameter to update the value of AutoUpdateEnabled |

Supports update: yes

Discovery: supported
