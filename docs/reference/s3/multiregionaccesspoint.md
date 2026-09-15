# aws.multiregionaccesspoint

**CloudFormation type:** `AWS::S3::MultiRegionAccessPoint`

AWS::S3::MultiRegionAccessPoint is an Amazon S3 resource type that dynamically routes S3 requests to easily satisfy geographic compliance requirements based on customer-defined routing policies.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::S3::MultiRegionAccessPoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | computed |  | The alias is a unique identifier to, and is part of the public DNS name for this Multi Region Access Point |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of the when the Multi Region Access Point is created |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name you want to assign to this Multi Region Access Point. |
| `PublicAccessBlockConfiguration` | public_access_block_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The PublicAccessBlock configuration that you want to apply to this Multi Region Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide. |
| `Regions` |  | `list` | required, replaces on change |  | The list of buckets that you want to associate this Multi Region Access Point with. |

Supports update: no

Discovery: supported
