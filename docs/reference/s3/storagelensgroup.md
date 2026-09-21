# aws.storagelensgroup

**CloudFormation type:** `AWS::S3::StorageLensGroup`

The AWS::S3::StorageLensGroup resource is an Amazon S3 resource type that you can use to create Storage Lens Group.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::S3::StorageLensGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Filter` |  | `map` | required |  | Sets the Storage Lens Group filter. |
| `Name` |  | `string` | required, replaces on change |  | The name that identifies the Amazon S3 Storage Lens Group. |
| `StorageLensGroupArn` | storage_lens_group_arn | `string` | computed |  | The ARN for the Amazon S3 Storage Lens Group. |
| `Tags` |  | `map` | tags map |  | A set of tags (key-value pairs) for this Amazon S3 Storage Lens Group. |

Supports update: yes

Discovery: supported
