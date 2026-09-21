# aws.wafv2.rulegroup

**CloudFormation type:** `AWS::WAFv2::RuleGroup`

Contains the Rules that identify the requests that you want to allow, block, or count. In a RuleGroup, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a RuleGroup, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the RuleGroup with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a RuleGroup, a request needs to match only one of the specifications to be allowed, blocked, or counted.

Region attribute: `region`

**Import ID:** `<region>/Name|Id|Scope` (AWS::WAFv2::RuleGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | ARN of the WAF entity. |
| `AvailableLabels` | available_labels | `list` | optional, computed, provider-chosen |  | Collection of Available Labels. |
| `Capacity` |  | `integer` | required |  |  |
| `ConsumedLabels` | consumed_labels | `list` | optional, computed, provider-chosen |  | Collection of Consumed Labels. |
| `CustomResponseBodies` | custom_response_bodies | `map` | optional, computed, provider-chosen |  | Custom response key and body map. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the entity. |
| `Id` |  | `string` | computed |  | Id of the RuleGroup |
| `LabelNamespace` | label_namespace | `string` | computed |  | Name of the Label. |
| `MonetizationConfig` | monetization_config | `map` | optional, computed, provider-chosen |  | Configures monetization for the web ACL or rule group. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the RuleGroup. |
| `Rules` |  | `list` | optional, computed, provider-chosen |  | Collection of Rules. |
| `Scope` |  | `string` | required, replaces on change |  | Use CLOUDFRONT for CloudFront RuleGroup, use REGIONAL for Application Load Balancer and API Gateway. |
| `Tags` |  | `map` | tags map |  |  |
| `VisibilityConfig` | visibility_config | `map` | required |  | Visibility Metric of the RuleGroup. |

Supports update: yes

Discovery: supported (parent resource required)
