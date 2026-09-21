# aws.configurationaggregator

**CloudFormation type:** `AWS::Config::ConfigurationAggregator`

Resource Type definition for AWS::Config::ConfigurationAggregator

Region attribute: `region`

**Import ID:** `<region>/ConfigurationAggregatorName` (AWS::Config::ConfigurationAggregator)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountAggregationSources` | account_aggregation_sources | `list` | optional, computed, provider-chosen |  |  |
| `ConfigurationAggregatorArn` | configuration_aggregator_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the aggregator. |
| `ConfigurationAggregatorName` | configuration_aggregator_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the aggregator. |
| `OrganizationAggregationSource` | organization_aggregation_source | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  | The tags for the configuration aggregator. |

Supports update: yes

Discovery: supported
