# aws.adapter

**CloudFormation type:** `AWS::Textract::Adapter`

The AWS::Textract::Adapter resource creates an Amazon Textract adapter, which can be fine-tuned for enhanced performance on user-provided documents.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Textract::Adapter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdapterId` | adapter_id | `string` | computed |  | A unique identifier for the adapter resource. |
| `AdapterName` | adapter_name | `string` | required |  | The name to be assigned to the adapter being created. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the adapter. |
| `AutoUpdate` | auto_update | `string` | optional, computed, provider-chosen |  | Controls whether or not the adapter should automatically update. |
| `CreationTime` | creation_time | `string` | computed |  | The date and time that the adapter was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description to be assigned to the adapter being created. |
| `FeatureTypes` | feature_types | `list` | required, replaces on change |  | The type of feature that the adapter is being trained on. Currently, supported feature types are: QUERIES |
| `Tags` |  | `map` | tags map |  | A list of tags to be added to the adapter. |

Supports update: yes

Discovery: supported
