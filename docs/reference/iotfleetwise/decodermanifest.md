# aws.decodermanifest

**CloudFormation type:** `AWS::IoTFleetWise::DecoderManifest`

Definition of AWS::IoTFleetWise::DecoderManifest Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoTFleetWise::DecoderManifest)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `DefaultForUnmappedSignals` | default_for_unmapped_signals | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `LastModificationTime` | last_modification_time | `string` | computed |  |  |
| `ModelManifestArn` | model_manifest_arn | `string` | required, replaces on change | aws.modelmanifest.Arn |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `NetworkInterfaces` | network_interfaces | `list` | optional, computed, provider-chosen |  |  |
| `SignalDecoders` | signal_decoders | `list` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
