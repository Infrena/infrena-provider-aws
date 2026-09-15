# aws.signingconfiguration

**CloudFormation type:** `AWS::ECR::SigningConfiguration`

The AWS::ECR::SigningConfiguration resource creates or updates the signing configuration for an Amazon ECR registry.

Region attribute: `region`

**Import ID:** `<region>/RegistryId` (AWS::ECR::SigningConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RegistryId` | registry_id | `string` | computed |  | 12-digit AWS account ID of the ECR registry. |
| `Rules` |  | `list` | required |  | Array of signing rules that define which repositories should be signed and with which signing profiles. |

Supports update: yes

Discovery: supported
