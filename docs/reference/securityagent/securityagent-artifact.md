# aws.securityagent.artifact

**CloudFormation type:** `AWS::SecurityAgent::Artifact`

Uploads an artifact to an agent space. Artifacts provide additional context for security testing, such as architecture diagrams, API specifications, or configuration files.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SecurityAgent::Artifact)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentSpaceId` | agent_space_id | `string` | required, replaces on change | aws.securityagent.agentspace.AgentSpaceId | The unique identifier of the agent space to add the artifact to. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the artifact. |
| `ArtifactContent` | artifact_content | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The binary content of the artifact to upload, encoded as a Base64 string. |
| `ArtifactId` | artifact_id | `string` | computed |  | The unique identifier assigned to the uploaded artifact. |
| `ArtifactType` | artifact_type | `string` | required, replaces on change |  | The file type of the artifact. |
| `FileName` | file_name | `string` | required, replaces on change |  | The file name of the artifact. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time the artifact was last updated, in ISO 8601 format. |

Supports update: no

Discovery: supported (parent resource required)
