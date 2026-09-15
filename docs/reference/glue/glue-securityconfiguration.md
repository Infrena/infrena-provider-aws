# aws.glue.securityconfiguration

**CloudFormation type:** `AWS::Glue::SecurityConfiguration`

Resource Type definition for AWS::Glue::SecurityConfiguration

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::SecurityConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `EncryptionConfiguration` | encryption_configuration | `map` | required, replaces on change |  | The encryption configuration for the security configuration. |
| `Name` |  | `string` | required, replaces on change |  | The name for the security configuration. |

Supports update: no

Discovery: supported
