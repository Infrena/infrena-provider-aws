# aws.inboundexternallink

**CloudFormation type:** `AWS::RTBFabric::InboundExternalLink`

Resource Type definition for AWS::RTBFabric::InboundExternalLink Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RTBFabric::InboundExternalLink)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  |  |
| `DomainName` | domain_name | `string` | computed |  |  |
| `GatewayId` | gateway_id | `string` | required |  |  |
| `LinkAttributes` | link_attributes | `map` | optional, computed, provider-chosen |  |  |
| `LinkId` | link_id | `string` | computed |  |  |
| `LinkLogSettings` | link_log_settings | `map` | required |  |  |
| `LinkStatus` | link_status | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the Link. |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
