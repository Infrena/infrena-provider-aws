# aws.lightsail.distribution

**CloudFormation type:** `AWS::Lightsail::Distribution`

Resource Type definition for AWS::Lightsail::Distribution

Region attribute: `region`

**Import ID:** `<region>/DistributionName` (AWS::Lightsail::Distribution)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AbleToUpdateBundle` | able_to_update_bundle | `boolean` | computed |  | Indicates whether the bundle that is currently applied to your distribution, specified using the distributionName parameter, can be changed to another bundle. |
| `BundleId` | bundle_id | `string` | required |  | The bundle ID to use for the distribution. |
| `CacheBehaviorSettings` | cache_behavior_settings | `map` | optional, computed, provider-chosen |  | Describes the cache settings of an Amazon Lightsail content delivery network (CDN) distribution. |
| `CacheBehaviors` | cache_behaviors | `list` | optional, computed, provider-chosen |  | An array of objects that describe the per-path cache behavior for the distribution. |
| `CertificateName` | certificate_name | `string` | optional, computed, provider-chosen |  | The certificate attached to the Distribution. |
| `DefaultCacheBehavior` | default_cache_behavior | `map` | required |  | Describes the default cache behavior of an Amazon Lightsail content delivery network (CDN) distribution. |
| `DistributionArn` | distribution_arn | `string` | computed |  |  |
| `DistributionName` | distribution_name | `string` | required, replaces on change |  | The name for the distribution. |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen, replaces on change |  | The IP address type for the distribution. |
| `IsEnabled` | is_enabled | `boolean` | optional, computed, provider-chosen |  | Indicates whether the distribution is enabled. |
| `Origin` |  | `map` | required |  | Describes the origin resource of an Amazon Lightsail content delivery network (CDN) distribution. |
| `Status` |  | `string` | computed |  | The status of the distribution. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
