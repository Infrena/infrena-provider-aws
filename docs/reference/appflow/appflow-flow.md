# aws.appflow.flow

**CloudFormation type:** `AWS::AppFlow::Flow`

Resource schema for AWS::AppFlow::Flow.

Region attribute: `region`

**Import ID:** `<region>/FlowName` (AWS::AppFlow::Flow)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the flow. |
| `DestinationFlowConfigList` | destination_flow_config_list | `list` | required |  | List of Destination connectors of the flow. |
| `FlowArn` | flow_arn | `string` | computed |  | ARN identifier of the flow. |
| `FlowName` | flow_name | `string` | required, replaces on change |  | Name of the flow. |
| `FlowStatus` | flow_status | `string` | optional, computed, provider-chosen |  | Flow activation status for Scheduled- and Event-triggered flows |
| `KMSArn` | kms_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key. |
| `MetadataCatalogConfig` | metadata_catalog_config | `map` | optional, computed, provider-chosen |  | Configurations of metadata catalog of the flow. |
| `SourceFlowConfig` | source_flow_config | `map` | required |  | Configurations of Source connector of the flow. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | List of Tags. |
| `Tasks` |  | `list` | required |  | List of tasks for the flow. |
| `TriggerConfig` | trigger_config | `map` | required |  | Trigger settings of the flow. |

Supports update: yes

Discovery: supported
