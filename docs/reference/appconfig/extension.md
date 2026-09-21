# aws.extension

**CloudFormation type:** `AWS::AppConfig::Extension`

Resource Type definition for AWS::AppConfig::Extension

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::AppConfig::Extension)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `map` | required |  |  |
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the extension. |
| `Id` |  | `string` | computed |  |  |
| `LatestVersionNumber` | latest_version_number | `integer` | optional, computed, provider-chosen, write-only |  |  |
| `Name` |  | `string` | required, replaces on change |  | Name of the extension. |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  | An array of key-value tags to apply to this resource. |
| `VersionNumber` | version_number | `integer` | computed |  |  |

Supports update: yes

Discovery: supported
