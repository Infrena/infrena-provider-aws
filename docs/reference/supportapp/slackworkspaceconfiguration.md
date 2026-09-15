# aws.slackworkspaceconfiguration

**CloudFormation type:** `AWS::SupportApp::SlackWorkspaceConfiguration`

An AWS Support App resource that creates, updates, lists, and deletes Slack workspace configurations.

Region attribute: `region`

**Import ID:** `<region>/TeamId` (AWS::SupportApp::SlackWorkspaceConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `TeamId` | team_id | `string` | required, replaces on change |  | The team ID in Slack, which uniquely identifies a workspace. |
| `VersionId` | version_id | `string` | optional, computed, provider-chosen, write-only |  | An identifier used to update an existing Slack workspace configuration in AWS CloudFormation. |

Supports update: yes

Discovery: supported
