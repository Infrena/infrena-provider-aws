# aws.crawler

**CloudFormation type:** `AWS::Glue::Crawler`

Resource Type definition for AWS::Glue::Crawler

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::Crawler)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Classifiers` |  | `list` | optional, computed, provider-chosen |  | A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler. |
| `Configuration` |  | `string` | optional, computed, provider-chosen |  | Crawler configuration information. This versioned JSON string allows users to specify aspects of a crawler's behavior. |
| `CrawlerSecurityConfiguration` | crawler_security_configuration | `string` | optional, computed, provider-chosen |  | The name of the SecurityConfiguration structure to be used by this crawler. |
| `DatabaseName` | database_name | `string` | optional, computed, provider-chosen |  | The name of the database in which the crawler's output is stored. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the crawler. |
| `LakeFormationConfiguration` | lake_formation_configuration | `map` | optional, computed, provider-chosen |  | Specifies AWS Lake Formation configuration settings for the crawler |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the crawler. |
| `RecrawlPolicy` | recrawl_policy | `map` | optional, computed, provider-chosen |  | When crawling an Amazon S3 data source after the first crawl is complete, specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run. For more information, see Incremental Crawls in AWS Glue in the developer guide. |
| `Role` |  | `string` | required |  | The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data. |
| `Schedule` |  | `map` | optional, computed, provider-chosen |  | A scheduling object using a cron statement to schedule an event. |
| `SchemaChangePolicy` | schema_change_policy | `map` | optional, computed, provider-chosen |  | The policy that specifies update and delete behaviors for the crawler. The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The SchemaChangePolicy does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the SchemaChangePolicy on a crawler. The SchemaChangePolicy consists of two components, UpdateBehavior and DeleteBehavior. |
| `TablePrefix` | table_prefix | `string` | optional, computed, provider-chosen |  | The prefix added to the names of tables that are created. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags to use with this crawler. |
| `Targets` |  | `map` | required |  | Specifies data stores to crawl. |

Supports update: yes

Discovery: supported
