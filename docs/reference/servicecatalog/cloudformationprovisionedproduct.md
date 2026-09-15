# aws.cloudformationprovisionedproduct

**CloudFormation type:** `AWS::ServiceCatalog::CloudFormationProvisionedProduct`

Resource Schema for AWS::ServiceCatalog::CloudFormationProvisionedProduct

Region attribute: `region`

**Import ID:** `<region>/ProvisionedProductId` (AWS::ServiceCatalog::CloudFormationProvisionedProduct)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen |  |  |
| `CloudformationStackArn` | cloudformation_stack_arn | `string` | computed |  |  |
| `NotificationArns` | notification_arns | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `Outputs` |  | `map` | computed |  | List of key-value pair outputs. |
| `PathId` | path_id | `string` | optional, computed, provider-chosen |  |  |
| `PathName` | path_name | `string` | optional, computed, provider-chosen |  |  |
| `ProductId` | product_id | `string` | optional, computed, provider-chosen |  |  |
| `ProductName` | product_name | `string` | optional, computed, provider-chosen |  |  |
| `ProvisionedProductId` | provisioned_product_id | `string` | computed |  |  |
| `ProvisionedProductName` | provisioned_product_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ProvisioningArtifactId` | provisioning_artifact_id | `string` | optional, computed, provider-chosen |  |  |
| `ProvisioningArtifactName` | provisioning_artifact_name | `string` | optional, computed, provider-chosen |  |  |
| `ProvisioningParameters` | provisioning_parameters | `list` | optional, computed, provider-chosen |  |  |
| `ProvisioningPreferences` | provisioning_preferences | `map` | optional, computed, provider-chosen |  |  |
| `RecordId` | record_id | `string` | computed |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: not supported
