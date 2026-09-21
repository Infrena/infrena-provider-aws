# aws.channelgroup

**CloudFormation type:** `AWS::MediaPackageV2::ChannelGroup`

<p>Represents a channel group that facilitates the grouping of multiple channels.</p>

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MediaPackageV2::ChannelGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) associated with the resource.</p> |
| `ChannelGroupName` | channel_group_name | `string` | required, replaces on change |  |  |
| `CreatedAt` | created_at | `string` | computed |  | <p>The date and time the channel group was created.</p> |
| `Description` |  | `string` | optional, computed, provider-chosen |  | <p>Enter any descriptive text that helps you to identify the channel group.</p> |
| `EgressDomain` | egress_domain | `string` | computed |  | <p>The output domain where the source stream should be sent. Integrate the domain with a downstream CDN (such as Amazon CloudFront) or playback device.</p> |
| `ModifiedAt` | modified_at | `string` | computed |  | <p>The date and time the channel group was modified.</p> |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
