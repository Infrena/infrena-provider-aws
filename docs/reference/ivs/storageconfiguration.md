# aws.storageconfiguration

**CloudFormation type:** `AWS::IVS::StorageConfiguration`

Resource Type definition for AWS::IVS::StorageConfiguration

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::StorageConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Storage Configuration ARN is automatically generated on creation and assigned as the unique identifier. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Storage Configuration Name. |
| `S3` |  | `map` | required, replaces on change |  | A complex type that describes an S3 location where recorded videos will be stored. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the asset model. |

Supports update: yes

Discovery: supported
