# aws.farm

**CloudFormation type:** `AWS::Deadline::Farm`

Resource Type definition for AWS::Deadline::Farm

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Deadline::Farm)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CostScaleFactor` | cost_scale_factor | `float` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `DisplayName` | display_name | `string` | required |  |  |
| `FarmId` | farm_id | `string` | computed |  |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
