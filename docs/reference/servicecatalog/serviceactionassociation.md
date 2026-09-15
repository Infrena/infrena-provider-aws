# aws.serviceactionassociation

**CloudFormation type:** `AWS::ServiceCatalog::ServiceActionAssociation`

Resource Schema for AWS::ServiceCatalog::ServiceActionAssociation

Region attribute: `region`

**Import ID:** `<region>/ProductId|ProvisioningArtifactId|ServiceActionId` (AWS::ServiceCatalog::ServiceActionAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ProductId` | product_id | `string` | required, replaces on change |  |  |
| `ProvisioningArtifactId` | provisioning_artifact_id | `string` | required, replaces on change |  |  |
| `ServiceActionId` | service_action_id | `string` | required, replaces on change | aws.serviceaction.Id |  |

Supports update: no

Discovery: supported (parent resource required)
