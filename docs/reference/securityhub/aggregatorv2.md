# aws.aggregatorv2

**CloudFormation type:** `AWS::SecurityHub::AggregatorV2`

The AWS::SecurityHub::AggregatorV2 resource represents the AWS Security Hub AggregatorV2 in your account. One aggregatorv2 resource is created for each account in non opt-in region in which you configure region linking mode.

Region attribute: `region`

**Import ID:** `<region>/AggregatorV2Arn` (AWS::SecurityHub::AggregatorV2)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AggregationRegion` | aggregation_region | `string` | computed |  | The aggregation Region of the AggregatorV2 |
| `AggregatorV2Arn` | aggregator_v2_arn | `string` | computed |  | The ARN of the AggregatorV2 being created and assigned as the unique identifier |
| `LinkedRegions` | linked_regions | `list` | required |  | The list of included Regions |
| `RegionLinkingMode` | region_linking_mode | `string` | required |  | Indicates to link a list of included Regions |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with the Security Hub V2 resource. |

Supports update: yes

Discovery: supported
