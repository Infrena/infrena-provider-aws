# aws.packagingconfiguration

**CloudFormation type:** `AWS::MediaPackage::PackagingConfiguration`

Resource schema for AWS::MediaPackage::PackagingConfiguration

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaPackage::PackagingConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the PackagingConfiguration. |
| `CmafPackage` | cmaf_package | `map` | optional, computed, provider-chosen, replaces on change |  | A CMAF packaging configuration. |
| `DashPackage` | dash_package | `map` | optional, computed, provider-chosen, replaces on change |  | A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration. |
| `HlsPackage` | hls_package | `map` | optional, computed, provider-chosen, replaces on change |  | An HTTP Live Streaming (HLS) packaging configuration. |
| `Id` |  | `string` | required, replaces on change |  | The ID of the PackagingConfiguration. |
| `MssPackage` | mss_package | `map` | optional, computed, provider-chosen, replaces on change |  | A Microsoft Smooth Streaming (MSS) PackagingConfiguration. |
| `PackagingGroupId` | packaging_group_id | `string` | required, replaces on change | aws.packaginggroup.Id | The ID of a PackagingGroup. |
| `Tags` |  | `map` | replaces on change, tags map |  | A collection of tags associated with a resource |

Supports update: no

Discovery: supported
