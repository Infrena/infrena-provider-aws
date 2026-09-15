# aws.opensearch.datasource

**CloudFormation type:** `AWS::OpenSearch::DataSource`

Creates a data source for an Amazon OpenSearch Service domain.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::OpenSearch::DataSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the data source. |
| `DataSourceType` | data_source_type | `map` | required |  | The type of data source. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the data source. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The name of the OpenSearch Service domain. |
| `Name` |  | `string` | required, replaces on change |  | The name of the data source. |
| `Status` |  | `string` | computed |  | The status of the data source. |

Supports update: yes

Discovery: supported (parent resource required)
