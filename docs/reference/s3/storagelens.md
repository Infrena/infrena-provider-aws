# aws.storagelens

**CloudFormation type:** `AWS::S3::StorageLens`

The AWS::S3::StorageLens resource is an Amazon S3 resource type that you can use to create Storage Lens configurations.

Region attribute: `region`

**Import ID:** `<region>/StorageLensConfiguration/Id` (AWS::S3::StorageLens)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `StorageLensConfiguration` | storage_lens_configuration | `map` | required |  | Specifies the details of Amazon S3 Storage Lens configuration. |
| `Tags` |  | `map` | tags map |  | A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration. |

Supports update: yes

Discovery: supported
