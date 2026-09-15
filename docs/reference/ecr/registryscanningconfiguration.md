# aws.registryscanningconfiguration

**CloudFormation type:** `AWS::ECR::RegistryScanningConfiguration`

The scanning configuration for a private registry.

Region attribute: `region`

**Import ID:** `<region>/RegistryId` (AWS::ECR::RegistryScanningConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RegistryId` | registry_id | `string` | computed |  | The registry id. |
| `Rules` |  | `list` | required |  | The scanning rules associated with the registry. A registry scanning configuration may contain a maximum of 2 rules. |
| `ScanType` | scan_type | `string` | required |  | The type of scanning configured for the registry. |

Supports update: yes

Discovery: supported
