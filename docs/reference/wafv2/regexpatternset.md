# aws.regexpatternset

**CloudFormation type:** `AWS::WAFv2::RegexPatternSet`

Contains a list of Regular expressions based on the provided inputs. RegexPatternSet can be used with other WAF entities with RegexPatternSetReferenceStatement to perform other actions .

Region attribute: `region`

**Import ID:** `<region>/Name|Id|Scope` (AWS::WAFv2::RegexPatternSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | ARN of the WAF entity. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the entity. |
| `Id` |  | `string` | computed |  | Id of the RegexPatternSet |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the RegexPatternSet. |
| `RegularExpressionList` | regular_expression_list | `list` | required |  |  |
| `Scope` |  | `string` | required, replaces on change |  | Use CLOUDFRONT for CloudFront RegexPatternSet, use REGIONAL for Application Load Balancer and API Gateway. |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported (parent resource required)
