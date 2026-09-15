# aws.vodsource

**CloudFormation type:** `AWS::MediaTailor::VodSource`

Definition of AWS::MediaTailor::VodSource Resource Type

Region attribute: `region`

**Import ID:** `<region>/SourceLocationName|VodSourceName` (AWS::MediaTailor::VodSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The ARN of the VOD source.</p> |
| `HttpPackageConfigurations` | http_package_configurations | `list` | required |  | <p>A list of HTTP package configuration parameters for this VOD source.</p> |
| `SourceLocationName` | source_location_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to assign to the VOD source. |
| `VodSourceName` | vod_source_name | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported (parent resource required)
