# aws.assetmodel

**CloudFormation type:** `AWS::IoTSiteWise::AssetModel`

Resource schema for AWS::IoTSiteWise::AssetModel

Region attribute: `region`

**Import ID:** `<region>/AssetModelId` (AWS::IoTSiteWise::AssetModel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssetModelArn` | asset_model_arn | `string` | computed |  | The ARN of the asset model, which has the following format. |
| `AssetModelCompositeModels` | asset_model_composite_models | `list` | optional, computed, provider-chosen |  | The composite asset models that are part of this asset model. Composite asset models are asset models that contain specific properties. |
| `AssetModelDescription` | asset_model_description | `string` | optional, computed, provider-chosen |  | A description for the asset model. |
| `AssetModelExternalId` | asset_model_external_id | `string` | optional, computed, provider-chosen |  | The external ID of the asset model. |
| `AssetModelHierarchies` | asset_model_hierarchies | `list` | optional, computed, provider-chosen |  | The hierarchy definitions of the asset model. Each hierarchy specifies an asset model whose assets can be children of any other assets created from this asset model. You can specify up to 10 hierarchies per asset model. |
| `AssetModelId` | asset_model_id | `string` | computed |  | The ID of the asset model. |
| `AssetModelName` | asset_model_name | `string` | required |  | A unique, friendly name for the asset model. |
| `AssetModelProperties` | asset_model_properties | `list` | optional, computed, provider-chosen |  | The property definitions of the asset model. You can specify up to 200 properties per asset model. |
| `AssetModelType` | asset_model_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of the asset model (ASSET_MODEL OR COMPONENT_MODEL or INTERFACE) |
| `EnforcedAssetModelInterfaceRelationships` | enforced_asset_model_interface_relationships | `list` | optional, computed, provider-chosen |  | a list of asset model and interface relationships |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the asset model. |

Supports update: yes

Discovery: supported
