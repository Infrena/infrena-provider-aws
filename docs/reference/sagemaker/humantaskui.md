# aws.humantaskui

**CloudFormation type:** `AWS::SageMaker::HumanTaskUi`

Resource Type definition for AWS::SageMaker::HumanTaskUi

Region attribute: `region`

**Import ID:** `<region>/HumanTaskUiArn` (AWS::SageMaker::HumanTaskUi)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The timestamp when the human task user interface was created. |
| `HumanTaskUiArn` | human_task_ui_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the human task user interface. |
| `HumanTaskUiName` | human_task_ui_name | `string` | required, replaces on change |  | The name of the human task user interface. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs that contain metadata to help you categorize and organize a human review workflow user interface. |
| `UiTemplate` | ui_template | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | The Liquid template for the worker user interface. |

Supports update: yes

Discovery: supported
