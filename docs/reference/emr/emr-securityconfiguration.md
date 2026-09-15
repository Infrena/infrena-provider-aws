# aws.emr.securityconfiguration

**CloudFormation type:** `AWS::EMR::SecurityConfiguration`

Use a SecurityConfiguration resource to configure data encryption, Kerberos authentication, and Amazon S3 authorization for EMRFS.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::EMR::SecurityConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the security configuration. |
| `SecurityConfiguration` | security_configuration | `string` | required, replaces on change |  | The security configuration details in JSON format. |

Supports update: no

Discovery: supported
