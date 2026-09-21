# aws.findingaggregator

**CloudFormation type:** `AWS::SecurityHub::FindingAggregator`

The ``AWS::SecurityHub::FindingAggregator`` resource enables cross-Region aggregation. When cross-Region aggregation is enabled, you can aggregate findings, finding updates, insights, control compliance statuses, and security scores from one or more linked Regions to a single aggregation Region. You can then view and manage all of this data from the aggregation Region. For more details about cross-Region aggregation, see [Cross-Region aggregation](https://docs.aws.amazon.com/securityhub/latest/userguide/finding-aggregation.html) in the *User Guide*

Region attribute: `region`

**Import ID:** `<region>/FindingAggregatorArn` (AWS::SecurityHub::FindingAggregator)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `FindingAggregationRegion` | finding_aggregation_region | `string` | computed |  |  |
| `FindingAggregatorArn` | finding_aggregator_arn | `string` | computed |  |  |
| `RegionLinkingMode` | region_linking_mode | `string` | required |  | Indicates whether to aggregate findings from all of the available Regions in the current partition. Also determines whether to automatically aggregate findings from new Regions as Security Hub supports them and you opt into them. |
| `Regions` |  | `list` | optional, computed, provider-chosen |  | If ``RegionLinkingMode`` is ``ALL_REGIONS_EXCEPT_SPECIFIED``, then this is a space-separated list of Regions that do not aggregate findings to the aggregation Region. |

Supports update: yes

Discovery: supported
