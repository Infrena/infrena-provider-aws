# aws.mediapackage.channel

**CloudFormation type:** `AWS::MediaPackage::Channel`

Resource schema for AWS::MediaPackage::Channel

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaPackage::Channel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) assigned to the Channel. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A short text description of the Channel. |
| `EgressAccessLogs` | egress_access_logs | `map` | optional, computed, provider-chosen |  | The configuration parameters for egress access logging. |
| `HlsIngest` | hls_ingest | `map` | optional, computed, provider-chosen |  | An HTTP Live Streaming (HLS) ingest resource configuration. |
| `Id` |  | `string` | required, replaces on change |  | The ID of the Channel. |
| `IngressAccessLogs` | ingress_access_logs | `map` | optional, computed, provider-chosen |  | The configuration parameters for egress access logging. |
| `Tags` |  | `map` | replaces on change, tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
