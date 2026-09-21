# aws.integrationresourceproperty

**CloudFormation type:** `AWS::Glue::IntegrationResourceProperty`

Resource Type definition for AWS::Glue::IntegrationResourceProperty

Region attribute: `region`

**Import ID:** `<region>/ResourceArn|ResourcePropertyArn` (AWS::Glue::IntegrationResourceProperty)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The connection ARN of the source, or the database ARN of the target. |
| `ResourcePropertyArn` | resource_property_arn | `string` | computed |  | The integration resource property ARN. |
| `SourceProcessingProperties` | source_processing_properties | `map` | optional, computed, provider-chosen |  | The resource properties associated with the integration source. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetProcessingProperties` | target_processing_properties | `map` | optional, computed, provider-chosen |  | The resource properties associated with the integration target. |

Supports update: yes

Discovery: supported
