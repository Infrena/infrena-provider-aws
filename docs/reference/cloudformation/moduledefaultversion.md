# aws.moduledefaultversion

**CloudFormation type:** `AWS::CloudFormation::ModuleDefaultVersion`

A module that has been registered in the CloudFormation registry as the default version

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CloudFormation::ModuleDefaultVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the module version to set as the default version. |
| `ModuleName` | module_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of a module existing in the registry. |
| `VersionId` | version_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of an existing version of the named module to set as the default. |

Supports update: no

Discovery: supported
