# aws.sagemaker.hub

**CloudFormation type:** `AWS::SageMaker::Hub`

Resource type definition for AWS::SageMaker::Hub

Region attribute: `region`

**Import ID:** `<region>/HubArn` (AWS::SageMaker::Hub)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The date and time that the hub was created. |
| `HubArn` | hub_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the hub. |
| `HubDescription` | hub_description | `string` | required |  | A description of the hub. |
| `HubDisplayName` | hub_display_name | `string` | optional, computed, provider-chosen |  | The display name of the hub. |
| `HubName` | hub_name | `string` | required, replaces on change |  | The name of the hub. |
| `HubSearchKeywords` | hub_search_keywords | `list` | optional, computed, provider-chosen |  | The searchable keywords for the hub. |
| `HubStatus` | hub_status | `string` | computed |  | The status of the hub. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The date and time that the hub was last modified. |
| `S3StorageConfig` | s3_storage_config | `map` | optional, computed, provider-chosen, replaces on change |  | The Amazon S3 storage configuration for the hub. |
| `Tags` |  | `map` | tags map |  | Tags to associate with the hub. |

Supports update: yes

Discovery: supported
