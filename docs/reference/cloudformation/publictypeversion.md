# aws.publictypeversion

**CloudFormation type:** `AWS::CloudFormation::PublicTypeVersion`

Test and Publish a resource that has been registered in the CloudFormation Registry.

Region attribute: `region`

**Import ID:** `<region>/PublicTypeArn` (AWS::CloudFormation::PublicTypeVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The Amazon Resource Number (ARN) of the extension. |
| `LogDeliveryBucket` | log_delivery_bucket | `string` | optional, computed, provider-chosen, replaces on change |  | A url to the S3 bucket where logs for the testType run will be available |
| `PublicTypeArn` | public_type_arn | `string` | computed |  | The Amazon Resource Number (ARN) assigned to the public extension upon publication |
| `PublicVersionNumber` | public_version_number | `string` | optional, computed, provider-chosen, replaces on change |  | The version number of a public third-party extension |
| `PublisherId` | publisher_id | `string` | computed |  | The reserved publisher id for this type, or the publisher id assigned by CloudFormation for publishing in this region. |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  | The kind of extension |
| `TypeName` | type_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the type being registered. |
| `TypeVersionArn` | type_version_arn | `string` | computed |  | The Amazon Resource Number (ARN) of the extension with the versionId. |

Supports update: no

Discovery: supported
