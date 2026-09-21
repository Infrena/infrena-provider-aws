# aws.sourcelocation

**CloudFormation type:** `AWS::MediaTailor::SourceLocation`

Definition of AWS::MediaTailor::SourceLocation Resource Type

Region attribute: `region`

**Import ID:** `<region>/SourceLocationName` (AWS::MediaTailor::SourceLocation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessConfiguration` | access_configuration | `map` | optional, computed, provider-chosen |  | <p>Access configuration parameters.</p> |
| `Arn` |  | `string` | computed |  | <p>The ARN of the source location.</p> |
| `DefaultSegmentDeliveryConfiguration` | default_segment_delivery_configuration | `map` | optional, computed, provider-chosen |  | <p>The optional configuration for a server that serves segments. Use this if you want the segment delivery server to be different from the source location server. For example, you can configure your source location server to be an origination server, such as MediaPackage, and the segment delivery server to be a content delivery network (CDN), such as CloudFront. If you don't specify a segment delivery server, then the source location server is used.</p> |
| `HttpConfiguration` | http_configuration | `map` | required |  | <p>The HTTP configuration for the source location.</p> |
| `SegmentDeliveryConfigurations` | segment_delivery_configurations | `list` | optional, computed, provider-chosen |  | <p>A list of the segment delivery configurations associated with this resource.</p> |
| `SourceLocationName` | source_location_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  | The tags to assign to the source location. |

Supports update: yes

Discovery: supported
