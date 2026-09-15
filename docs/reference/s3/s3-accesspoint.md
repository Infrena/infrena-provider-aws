# aws.s3.accesspoint

**CloudFormation type:** `AWS::S3::AccessPoint`

The AWS::S3::AccessPoint resource is an Amazon S3 resource type that you can use to access buckets.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::S3::AccessPoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | computed |  | The alias of this Access Point. This alias can be used for compatibility purposes with other AWS services and third-party applications. |
| `Arn` |  | `string` | computed |  | the Amazon Resource Name (ARN) of the specified accesspoint. |
| `Bucket` |  | `string` | required, replaces on change |  | The name of the bucket that you want to associate this Access Point with. |
| `BucketAccountId` | bucket_account_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS account ID associated with the S3 bucket associated with this access point. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name. |
| `NetworkOrigin` | network_origin | `string` | computed |  | Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies. |
| `Policy` |  | `map` | optional, computed, provider-chosen |  | The Access Point Policy you want to apply to this access point. |
| `PublicAccessBlockConfiguration` | public_access_block_configuration | `map` | optional, computed, provider-chosen |  | The PublicAccessBlock configuration that you want to apply to this Access Point. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status 'The Meaning of Public' in the Amazon Simple Storage Service Developer Guide. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this S3 Access Point. |
| `VpcConfiguration` | vpc_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The Virtual Private Cloud (VPC) configuration for a bucket access point. |

Supports update: yes

Discovery: supported
