# aws.xray.group

**CloudFormation type:** `AWS::XRay::Group`

This schema provides construct and validation rules for AWS-XRay Group resource parameters.

Region attribute: `region`

**Import ID:** `<region>/GroupARN` (AWS::XRay::Group)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FilterExpression` | filter_expression | `string` | optional, computed, provider-chosen |  | The filter expression defining criteria by which to group traces. |
| `GroupARN` | group_arn | `string` | computed |  | The ARN of the group that was generated on creation. |
| `GroupName` | group_name | `string` | required |  | The case-sensitive name of the new group. Names must be unique. |
| `InsightsConfiguration` | insights_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
