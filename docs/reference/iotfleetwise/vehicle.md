# aws.vehicle

**CloudFormation type:** `AWS::IoTFleetWise::Vehicle`

Definition of AWS::IoTFleetWise::Vehicle Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::IoTFleetWise::Vehicle)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AssociationBehavior` | association_behavior | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Attributes` |  | `map` | optional, computed, provider-chosen |  |  |
| `CreationTime` | creation_time | `string` | computed |  |  |
| `DecoderManifestArn` | decoder_manifest_arn | `string` | required | aws.decodermanifest.Arn |  |
| `LastModificationTime` | last_modification_time | `string` | computed |  |  |
| `ModelManifestArn` | model_manifest_arn | `string` | required | aws.modelmanifest.Arn |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `StateTemplates` | state_templates | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
