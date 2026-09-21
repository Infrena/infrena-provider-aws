# aws.appblock

**CloudFormation type:** `AWS::AppStream::AppBlock`

Resource Type definition for AWS::AppStream::AppBlock

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::AppStream::AppBlock)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `PackagingType` | packaging_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PostSetupScriptDetails` | post_setup_script_details | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SetupScriptDetails` | setup_script_details | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SourceS3Location` | source_s3_location | `map` | required, replaces on change |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: not supported
