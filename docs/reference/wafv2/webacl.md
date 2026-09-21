# aws.webacl

**CloudFormation type:** `AWS::WAFv2::WebACL`

Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.

Region attribute: `region`

**Import ID:** `<region>/Name|Id|Scope` (AWS::WAFv2::WebACL)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationConfig` | application_config | `map` | optional, computed, provider-chosen |  | Configures the ability for the WAF; console to store and retrieve application attributes during the webacl; creation process. Application attributes help WAF; give recommendations for protection packs. |
| `Arn` |  | `string` | computed |  | ARN of the WAF entity. |
| `AssociationConfig` | association_config | `map` | optional, computed, provider-chosen |  | AssociationConfig for body inspection |
| `Capacity` |  | `integer` | computed |  |  |
| `CaptchaConfig` | captcha_config | `map` | optional, computed, provider-chosen |  |  |
| `ChallengeConfig` | challenge_config | `map` | optional, computed, provider-chosen |  |  |
| `CustomResponseBodies` | custom_response_bodies | `map` | optional, computed, provider-chosen |  | Custom response key and body map. |
| `DataProtectionConfig` | data_protection_config | `map` | optional, computed, provider-chosen |  | Collection of dataProtects. |
| `DefaultAction` | default_action | `map` | required |  | Default Action WebACL will take against ingress traffic when there is no matching Rule. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the entity. |
| `Id` |  | `string` | computed |  | Id of the WebACL |
| `LabelNamespace` | label_namespace | `string` | computed |  | Name of the Label. |
| `MonetizationConfig` | monetization_config | `map` | optional, computed, provider-chosen |  | Configures monetization for the web ACL or rule group. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the WebACL. |
| `OnSourceDDoSProtectionConfig` | on_source_d_do_s_protection_config | `map` | optional, computed, provider-chosen |  | Configures the options for on-source DDoS protection provided by supported resource type. |
| `Rules` |  | `list` | optional, computed, provider-chosen |  | Collection of Rules. |
| `Scope` |  | `string` | required, replaces on change |  | Use CLOUDFRONT for CloudFront WebACL, use REGIONAL for Application Load Balancer and API Gateway. |
| `Tags` |  | `map` | tags map |  |  |
| `TokenDomains` | token_domains | `list` | optional, computed, provider-chosen |  | List of domains to accept in web request tokens, in addition to the domain of the protected resource. |
| `VisibilityConfig` | visibility_config | `map` | required |  | Visibility Metric of the WebACL. |

Supports update: yes

Discovery: supported (parent resource required)
