# aws.rtbfabric.link

**CloudFormation type:** `AWS::RTBFabric::Link`

Resource Type definition for AWS::RTBFabric::Link Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RTBFabric::Link)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  |  |
| `GatewayId` | gateway_id | `string` | required |  |  |
| `HttpResponderAllowed` | http_responder_allowed | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `LinkAttributes` | link_attributes | `map` | optional, computed, provider-chosen |  |  |
| `LinkDirection` | link_direction | `string` | computed |  |  |
| `LinkId` | link_id | `string` | computed |  |  |
| `LinkLogSettings` | link_log_settings | `map` | required |  |  |
| `LinkStatus` | link_status | `string` | computed |  |  |
| `ModuleConfigurationList` | module_configuration_list | `list` | optional, computed, provider-chosen, write-only |  |  |
| `PeerGatewayId` | peer_gateway_id | `string` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the Link. |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
