# aws.recommender

**CloudFormation type:** `AWS::CustomerProfiles::Recommender`

Resource Type definition for AWS::CustomerProfiles::Recommender

Region attribute: `region`

**Import ID:** `<region>/DomainName|RecommenderName` (AWS::CustomerProfiles::Recommender)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the recommender was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the recommender. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The name of the domain for which the recommender will be created |
| `FailureReason` | failure_reason | `string` | computed |  | The reason for recommender failure. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp of when the recommender was last updated. |
| `LatestRecommenderUpdate` | latest_recommender_update | `map` | computed |  | Information about the latest recommender update |
| `RecommenderArn` | recommender_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the recommender. |
| `RecommenderConfig` | recommender_config | `map` | optional, computed, provider-chosen |  | Configuration for the recommender |
| `RecommenderName` | recommender_name | `string` | required, replaces on change |  | The name of the recommender |
| `RecommenderRecipeName` | recommender_recipe_name | `string` | required, replaces on change |  | The name of the recommender recipe. |
| `Status` |  | `string` | computed |  | The status of the recommender |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags used to organize, track, or control access for this resource. |
| `TrainingMetrics` | training_metrics | `list` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
