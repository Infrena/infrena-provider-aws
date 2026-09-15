# aws.scn.dataset

**CloudFormation type:** `AWS::SCN::Dataset`

Represents an AWS Supply Chain data lake dataset.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SCN::Dataset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the dataset. |
| `CreatedTime` | created_time | `string` | computed |  | The creation time of the dataset. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the dataset. |
| `InstanceId` | instance_id | `string` | required, replaces on change |  | The Amazon Web Services Supply Chain instance identifier. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The last modified time of the dataset. |
| `Name` |  | `string` | required, replaces on change |  | The name of the dataset. |
| `Namespace` |  | `string` | required, replaces on change |  | The namespace of the dataset. |
| `PartitionSpec` | partition_spec | `map` | optional, computed, provider-chosen, replaces on change |  | The partition specification of the dataset. |
| `Schema` |  | `map` | optional, computed, provider-chosen, replaces on change |  | The schema of the dataset. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the dataset. |

Supports update: yes

Discovery: supported (parent resource required)
