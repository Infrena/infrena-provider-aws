# aws.iotsitewise.gateway

**CloudFormation type:** `AWS::IoTSiteWise::Gateway`

Resource schema for AWS::IoTSiteWise::Gateway

Region attribute: `region`

**Import ID:** `<region>/GatewayId` (AWS::IoTSiteWise::Gateway)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `GatewayCapabilitySummaries` | gateway_capability_summaries | `list` | optional, computed, provider-chosen |  | A list of gateway capability summaries that each contain a namespace and status. |
| `GatewayId` | gateway_id | `string` | computed |  | The ID of the gateway device. |
| `GatewayName` | gateway_name | `string` | required |  | A unique, friendly name for the gateway. |
| `GatewayPlatform` | gateway_platform | `map` | required, replaces on change |  | Contains a gateway's platform information. |
| `GatewayVersion` | gateway_version | `string` | optional, computed, provider-chosen, replaces on change |  | The version of the gateway you want to create. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the gateway. |

Supports update: yes

Discovery: supported
