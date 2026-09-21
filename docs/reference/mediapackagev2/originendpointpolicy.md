# aws.originendpointpolicy

**CloudFormation type:** `AWS::MediaPackageV2::OriginEndpointPolicy`

<p>Represents a resource policy that allows or denies access to an origin endpoint.</p>

Region attribute: `region`

**Import ID:** `<region>/ChannelGroupName|ChannelName|OriginEndpointName` (AWS::MediaPackageV2::OriginEndpointPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CdnAuthConfiguration` | cdn_auth_configuration | `map` | optional, computed, provider-chosen |  | <p>The settings to enable CDN authorization headers in MediaPackage.</p> |
| `ChannelGroupName` | channel_group_name | `string` | required, replaces on change |  |  |
| `ChannelName` | channel_name | `string` | required, replaces on change |  |  |
| `OriginEndpointName` | origin_endpoint_name | `string` | required, replaces on change |  |  |
| `Policy` |  | `string` | required |  |  |

Supports update: yes

Discovery: not supported
