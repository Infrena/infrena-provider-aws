# aws.layerversion

**CloudFormation type:** `AWS::Lambda::LayerVersion`

Resource Type definition for AWS::Lambda::LayerVersion

Region attribute: `region`

**Import ID:** `<region>/LayerVersionArn` (AWS::Lambda::LayerVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CompatibleArchitectures` | compatible_architectures | `list` | optional, computed, provider-chosen, replaces on change |  | A list of compatible instruction set architectures. |
| `CompatibleRuntimes` | compatible_runtimes | `list` | optional, computed, provider-chosen, replaces on change |  | A list of compatible function runtimes. Used for filtering with ListLayers and ListLayerVersions. |
| `Content` |  | `map` | required, replaces on change, write-only |  | The function layer archive. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the version. |
| `LayerName` | layer_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name or Amazon Resource Name (ARN) of the layer. |
| `LayerVersionArn` | layer_version_arn | `string` | computed |  |  |
| `LicenseInfo` | license_info | `string` | optional, computed, provider-chosen, replaces on change |  | The layer's software license. |

Supports update: no

Discovery: supported (parent resource required)
