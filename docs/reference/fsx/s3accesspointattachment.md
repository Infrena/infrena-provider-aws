# aws.s3accesspointattachment

**CloudFormation type:** `AWS::FSx::S3AccessPointAttachment`

Resource type definition for AWS::FSx::S3AccessPointAttachment

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::FSx::S3AccessPointAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Lifecycle` | lifecycle_value | `string` | computed |  | The lifecycle status of the S3 access point attachment. |
| `Name` |  | `string` | required, replaces on change |  | The name of the S3 access point attachment; also used for the name of the S3 access point. |
| `OntapConfiguration` | ontap_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The OntapConfiguration of the S3 access point attachment. |
| `OpenZFSConfiguration` | open_zfs_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The OpenZFSConfiguration of the S3 access point attachment. |
| `S3AccessPoint` | s3_access_point | `map` | optional, computed, provider-chosen, replaces on change |  | The S3 access point configuration of the S3 access point attachment. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of Amazon FSx volume that the S3 access point is attached to. |

Supports update: no

Discovery: supported
