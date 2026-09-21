# aws.iotsitewise.project

**CloudFormation type:** `AWS::IoTSiteWise::Project`

Resource schema for AWS::IoTSiteWise::Project

Region attribute: `region`

**Import ID:** `<region>/ProjectId` (AWS::IoTSiteWise::Project)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssetIds` | asset_ids | `list` | optional, computed, provider-chosen | aws.iotsitewise.asset.AssetId | The IDs of the assets to be associated to the project. |
| `PortalId` | portal_id | `string` | required, replaces on change | aws.iotsitewise.portal.PortalId | The ID of the portal in which to create the project. |
| `ProjectArn` | project_arn | `string` | computed |  | The ARN of the project. |
| `ProjectDescription` | project_description | `string` | optional, computed, provider-chosen |  | A description for the project. |
| `ProjectId` | project_id | `string` | computed |  | The ID of the project. |
| `ProjectName` | project_name | `string` | required |  | A friendly name for the project. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the project. |

Supports update: yes

Discovery: supported
