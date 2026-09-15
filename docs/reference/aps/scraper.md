# aws.scraper

**CloudFormation type:** `AWS::APS::Scraper`

Resource Type definition for AWS::APS::Scraper

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::APS::Scraper)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | optional, computed, provider-chosen |  | Scraper alias. |
| `Arn` |  | `string` | computed |  | Scraper ARN. |
| `Destination` |  | `map` | required |  | Scraper metrics destination |
| `RoleArn` | role_arn | `string` | computed |  | IAM role ARN for the scraper. |
| `RoleConfiguration` | role_configuration | `map` | optional, computed, provider-chosen |  | Role configuration |
| `ScrapeConfiguration` | scrape_configuration | `map` | required |  | Scraper configuration |
| `ScraperId` | scraper_id | `string` | computed |  | Required to identify a specific scraper. |
| `ScraperLoggingConfiguration` | scraper_logging_configuration | `map` | optional, computed, provider-chosen |  | Configuration for scraper logging |
| `Source` |  | `map` | required, replaces on change |  | Scraper metrics source |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
