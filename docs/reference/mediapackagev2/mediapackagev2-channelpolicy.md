# aws.mediapackagev2.channelpolicy

**CloudFormation type:** `AWS::MediaPackageV2::ChannelPolicy`

<p>Represents a resource-based policy that allows or denies access to a channel.</p>

Region attribute: `region`

**Import ID:** `<region>/ChannelGroupName|ChannelName` (AWS::MediaPackageV2::ChannelPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChannelGroupName` | channel_group_name | `string` | required, replaces on change |  |  |
| `ChannelName` | channel_name | `string` | required, replaces on change |  |  |
| `Policy` |  | `string` | required |  |  |

Supports update: yes

Discovery: not supported
