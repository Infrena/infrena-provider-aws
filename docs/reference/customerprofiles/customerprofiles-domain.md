# aws.customerprofiles.domain

**CloudFormation type:** `AWS::CustomerProfiles::Domain`

A domain defined for 3rd party data source in Profile Service

Region attribute: `region`

**Import ID:** `<region>/DomainName` (AWS::CustomerProfiles::Domain)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time of this integration got created |
| `DataStore` | data_store | `map` | optional, computed, provider-chosen |  | Configuration and status of the data store for the domain. |
| `DeadLetterQueueUrl` | dead_letter_queue_url | `string` | optional, computed, provider-chosen |  | The URL of the SQS dead letter queue |
| `DefaultEncryptionKey` | default_encryption_key | `string` | optional, computed, provider-chosen |  | The default encryption key |
| `DefaultExpirationDays` | default_expiration_days | `integer` | required |  | The default number of days until the data within the domain expires. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The unique name of the domain. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The time of this integration got last updated at |
| `Matching` |  | `map` | optional, computed, provider-chosen |  | The process of matching duplicate profiles. If Matching = true, Amazon Connect Customer Profiles starts a weekly batch process called Identity Resolution Job. If you do not specify a date and time for Identity Resolution Job to run, by default it runs every Saturday at 12AM UTC to detect duplicate profiles in your domains. After the Identity Resolution Job completes, use the GetMatches API to return and review the results. Or, if you have configured ExportingConfig in the MatchingRequest, you can download the results from S3. |
| `RuleBasedMatching` | rule_based_matching | `map` | optional, computed, provider-chosen |  | The process of matching duplicate profiles using the Rule-Based matching. If RuleBasedMatching = true, Amazon Connect Customer Profiles will start to match and merge your profiles according to your configuration in the RuleBasedMatchingRequest. You can use the ListRuleBasedMatches and GetSimilarProfiles API to return and review the results. Also, if you have configured ExportingConfig in the RuleBasedMatchingRequest, you can download the results from S3. |
| `Stats` |  | `map` | computed |  | Usage-specific statistics about the domain. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the domain |

Supports update: yes

Discovery: supported
