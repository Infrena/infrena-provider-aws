# aws.s3express.accesspoint

**CloudFormation type:** `AWS::S3Express::AccessPoint`

The AWS::S3Express::AccessPoint resource is an Amazon S3 resource type that you can use to access buckets.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::S3Express::AccessPoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | the Amazon Resource Name (ARN) of the specified accesspoint. |
| `Bucket` |  | `string` | required, replaces on change |  | The name of the bucket that you want to associate this Access Point with. |
| `BucketAccountId` | bucket_account_id | `string` | optional, computed, provider-chosen, replaces on change |  | The AWS account ID associated with the S3 bucket associated with this access point. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name you want to assign to this Access Point. If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the access point name. For directory buckets, the access point name must consist of a base name that you provide and suﬃx that includes the ZoneID (AWS Availability Zone or Local Zone) of your bucket location, followed by --xa-s3. |
| `NetworkOrigin` | network_origin | `string` | computed |  | Indicates whether this Access Point allows access from the public Internet. If VpcConfiguration is specified for this Access Point, then NetworkOrigin is VPC, and the Access Point doesn't allow access from the public Internet. Otherwise, NetworkOrigin is Internet, and the Access Point allows access from the public Internet, subject to the Access Point and bucket access policies. |
| `Policy` |  | `map` | optional, computed, provider-chosen |  | The Access Point Policy you want to apply to this access point. |
| `PublicAccessBlockConfiguration` | public_access_block_configuration | `map` | optional, computed, provider-chosen |  | The PublicAccessBlock configuration that you want to apply to this Access Point. |
| `Scope` |  | `map` | optional, computed, provider-chosen |  | For directory buckets, you can ﬁlter access control to speciﬁc preﬁxes, API operations, or a combination of both. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VpcConfiguration` | vpc_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The Virtual Private Cloud (VPC) configuration for a bucket access point. |

Supports update: yes

Discovery: supported
