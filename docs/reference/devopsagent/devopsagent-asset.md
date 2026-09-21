# aws.devopsagent.asset

**CloudFormation type:** `AWS::DevOpsAgent::Asset`

Resource Type definition for AWS::DevOpsAgent::Asset. An asset attached to an existing AWS DevOps Agent Space. Customer-creatable types include skill, agents_md, and attachment; call ListAssetTypes for the current authoritative set.

Region attribute: `region`

**Import ID:** `<region>/AgentSpaceId|AssetId` (AWS::DevOpsAgent::Asset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentSpaceId` | agent_space_id | `string` | required, replaces on change | aws.devopsagent.agentspace.AgentSpaceId | The unique identifier of the parent Agent Space. The asset is created as a child of this agent space. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the asset. Nested under the parent Agent Space: arn:<partition>:aidevops:<region>:<account-id>:agentspace/<agentspace-id>/asset/<asset-id>. |
| `AssetId` | asset_id | `string` | computed |  | The unique identifier of the asset (assigned by the service on Create). |
| `AssetType` | asset_type | `string` | required, replaces on change |  | The type of asset. The Asset API treats this as an open string; call ListAssetTypes for the current authoritative set of supported types. As of launch, customer-creatable types include skill, agents_md, and attachment. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the asset was created. |
| `Files` |  | `list` | optional, computed, provider-chosen, write-only |  | Inline file list. Mutually exclusive with Zip; enforced by the handler at Create/Update time. Write-only: not repopulated by Read. |
| `Metadata` |  | `map` | optional, computed, provider-chosen |  | Asset metadata document. Required and optional keys depend on AssetType. Values may be strings, numbers, booleans, or lists of any of those - validated server-side; see the public Asset API docs for the per-type metadata schema. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the asset was last updated. |
| `Version` |  | `integer` | computed |  | The current asset version. Server-managed; bumps on every successful Update (including no-op updates). This is the drift signal for change detection. |
| `Zip` |  | `string` | optional, computed, provider-chosen, write-only |  | Base64-encoded zip bundle containing all files for the asset. Mutually exclusive with Files; enforced by the handler at Create/Update time. Write-only: not repopulated by Read. Server treats a zip as 'replace all files' (max 6 MiB). |

Supports update: yes

Discovery: supported (parent resource required)
