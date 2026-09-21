# aws.livesource

**CloudFormation type:** `AWS::MediaTailor::LiveSource`

Definition of AWS::MediaTailor::LiveSource Resource Type

Region attribute: `region`

**Import ID:** `<region>/LiveSourceName|SourceLocationName` (AWS::MediaTailor::LiveSource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The ARN of the live source.</p> |
| `HttpPackageConfigurations` | http_package_configurations | `list` | required |  | <p>A list of HTTP package configuration parameters for this live source.</p> |
| `LiveSourceName` | live_source_name | `string` | required, replaces on change |  |  |
| `SourceLocationName` | source_location_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  | The tags to assign to the live source. |

Supports update: yes

Discovery: supported (parent resource required)
