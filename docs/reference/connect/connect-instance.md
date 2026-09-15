# aws.connect.instance

**CloudFormation type:** `AWS::Connect::Instance`

Resource Type definition for AWS::Connect::Instance

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Connect::Instance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | An instanceArn is automatically generated on creation based on instanceId. |
| `Attributes` |  | `map` | required |  | The attributes for the instance. |
| `CreatedTime` | created_time | `string` | computed |  | Timestamp of instance creation logged as part of instance creation. |
| `DirectoryId` | directory_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Existing directoryId user wants to map to the new Connect instance. |
| `Id` |  | `string` | computed |  | An instanceId is automatically generated on creation and assigned as the unique identifier. |
| `IdentityManagementType` | identity_management_type | `string` | required, replaces on change |  | Specifies the type of directory integration for new instance. |
| `InstanceAlias` | instance_alias | `string` | optional, computed, provider-chosen, replaces on change |  | Alias of the new directory created as part of new instance creation. |
| `InstanceStatus` | instance_status | `string` | computed |  | Specifies the creation status of new instance. |
| `ServiceRole` | service_role | `string` | computed |  | Service linked role created as part of instance creation. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
