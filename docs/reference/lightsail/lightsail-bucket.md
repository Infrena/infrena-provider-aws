# aws.lightsail.bucket

**CloudFormation type:** `AWS::Lightsail::Bucket`

Resource Type definition for AWS::Lightsail::Bucket

Region attribute: `region`

**Import ID:** `<region>/BucketName` (AWS::Lightsail::Bucket)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AbleToUpdateBundle` | able_to_update_bundle | `boolean` | computed |  | Indicates whether the bundle that is currently applied to a bucket can be changed to another bundle. You can update a bucket's bundle only one time within a monthly AWS billing cycle. |
| `AccessRules` | access_rules | `map` | optional, computed, provider-chosen |  | An object that sets the public accessibility of objects in the specified bucket. |
| `BucketArn` | bucket_arn | `string` | computed |  |  |
| `BucketName` | bucket_name | `string` | required, replaces on change |  | The name for the bucket. |
| `BundleId` | bundle_id | `string` | required |  | The ID of the bundle to use for the bucket. |
| `ObjectVersioning` | object_versioning | `boolean` | optional, computed, provider-chosen |  | Specifies whether to enable or disable versioning of objects in the bucket. |
| `ReadOnlyAccessAccounts` | read_only_access_accounts | `list` | optional, computed, provider-chosen |  | An array of strings to specify the AWS account IDs that can access the bucket. |
| `ResourcesReceivingAccess` | resources_receiving_access | `list` | optional, computed, provider-chosen |  | The names of the Lightsail resources for which to set bucket access. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Url` |  | `string` | computed |  | The URL of the bucket. |

Supports update: yes

Discovery: supported
