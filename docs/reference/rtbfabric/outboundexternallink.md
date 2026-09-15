# aws.outboundexternallink

**CloudFormation type:** `AWS::RTBFabric::OutboundExternalLink`

Resource Type definition for AWS::RTBFabric::OutboundExternalLink Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RTBFabric::OutboundExternalLink)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  |  |
| `GatewayId` | gateway_id | `string` | required |  |  |
| `LinkAttributes` | link_attributes | `map` | optional, computed, provider-chosen |  |  |
| `LinkId` | link_id | `string` | computed |  |  |
| `LinkLogSettings` | link_log_settings | `map` | required |  |  |
| `LinkStatus` | link_status | `string` | computed |  |  |
| `PublicEndpoint` | public_endpoint | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the Link. |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
