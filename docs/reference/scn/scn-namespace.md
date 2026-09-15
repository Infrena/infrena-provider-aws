# aws.scn.namespace

**CloudFormation type:** `AWS::SCN::Namespace`

Definition of AWS::SCN::Namespace Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SCN::Namespace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the namespace. |
| `CreatedTime` | created_time | `string` | computed |  | The creation time of the namespace. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the namespace. |
| `InstanceId` | instance_id | `string` | required, replaces on change |  | The Amazon Web Services Supply Chain instance identifier. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The last modified time of the namespace. |
| `Name` |  | `string` | required, replaces on change |  | The name of the namespace. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags for the namespace. |

Supports update: yes

Discovery: supported (parent resource required)
