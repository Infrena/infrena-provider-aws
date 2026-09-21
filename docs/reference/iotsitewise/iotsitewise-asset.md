# aws.iotsitewise.asset

**CloudFormation type:** `AWS::IoTSiteWise::Asset`

Resource schema for AWS::IoTSiteWise::Asset

Region attribute: `region`

**Import ID:** `<region>/AssetId` (AWS::IoTSiteWise::Asset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssetArn` | asset_arn | `string` | computed |  | The ARN of the asset |
| `AssetDescription` | asset_description | `string` | optional, computed, provider-chosen |  | A description for the asset |
| `AssetExternalId` | asset_external_id | `string` | optional, computed, provider-chosen |  | The External ID of the asset |
| `AssetHierarchies` | asset_hierarchies | `list` | optional, computed, provider-chosen |  |  |
| `AssetId` | asset_id | `string` | computed |  | The ID of the asset |
| `AssetModelId` | asset_model_id | `string` | required | aws.assetmodel.AssetModelId | The ID of the asset model from which to create the asset. |
| `AssetName` | asset_name | `string` | required |  | A unique, friendly name for the asset. |
| `AssetProperties` | asset_properties | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the asset. |

Supports update: yes

Discovery: supported
