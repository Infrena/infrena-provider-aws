# aws.athena.workgroup

**CloudFormation type:** `AWS::Athena::WorkGroup`

Resource schema for AWS::Athena::WorkGroup

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Athena::WorkGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The date and time the workgroup was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The workgroup description. |
| `Name` |  | `string` | required, replaces on change |  | The workGroup name. |
| `RecursiveDeleteOption` | recursive_delete_option | `boolean` | optional, computed, provider-chosen, write-only |  | The option to delete the workgroup and its contents even if the workgroup contains any named queries. |
| `State` |  | `string` | optional, computed, provider-chosen |  | The state of the workgroup: ENABLED or DISABLED. |
| `Tags` |  | `map` | tags map |  | One or more tags, separated by commas, that you want to attach to the workgroup as you create it |
| `WorkGroupConfiguration` | work_group_configuration | `map` | optional, computed, provider-chosen |  | The workgroup configuration |
| `WorkGroupConfigurationUpdates` | work_group_configuration_updates | `map` | optional, computed, provider-chosen, write-only |  | The configuration information that will be updated for this workgroup, which includes the location in Amazon S3 where query results are stored, the encryption option, if any, used for query results, whether the Amazon CloudWatch Metrics are enabled for the workgroup, whether the workgroup settings override the client-side settings, and the data usage limit for the amount of bytes scanned per query, if it is specified. |

Supports update: yes

Discovery: supported
