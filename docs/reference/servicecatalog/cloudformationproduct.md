# aws.cloudformationproduct

**CloudFormation type:** `AWS::ServiceCatalog::CloudFormationProduct`

Resource type definition for AWS::ServiceCatalog::CloudFormationProduct

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalog::CloudFormationProduct)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AcceptLanguage` | accept_language | `string` | optional, computed, provider-chosen, write-only |  | The language code. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the product. |
| `Distributor` |  | `string` | optional, computed, provider-chosen |  | The distributor of the product. |
| `Id` |  | `string` | computed |  | The ID of the product, such as prod-tsjbmal34qvek |
| `Name` |  | `string` | required |  | The name of the product. |
| `Owner` |  | `string` | required |  | The owner of the product. |
| `ProductName` | product_name | `string` | computed |  | The name of the product. |
| `ProductType` | product_type | `string` | optional, computed, provider-chosen |  | The type of product. |
| `ProvisioningArtifactIds` | provisioning_artifact_ids | `string` | computed |  | The IDs of the provisioning artifacts |
| `ProvisioningArtifactNames` | provisioning_artifact_names | `string` | computed |  | The names of the provisioning artifacts |
| `ProvisioningArtifactParameters` | provisioning_artifact_parameters | `list` | optional, computed, provider-chosen |  | The configuration of the provisioning artifact (also known as a version). |
| `ReplaceProvisioningArtifacts` | replace_provisioning_artifacts | `boolean` | optional, computed, provider-chosen, write-only |  | This property is turned off by default. If turned off, you can update provisioning artifacts or product attributes (such as description, distributor, name, owner, and more) and the associated provisioning artifacts will retain the same unique identifier. Provisioning artifacts are matched within the CloudFormationProduct resource, and only those that have been updated will be changed. Provisioning artifacts are matched by a combinaton of provisioning artifact template URL and name. |
| `SourceConnection` | source_connection | `map` | optional, computed, provider-chosen |  | A top level ProductViewDetail response containing details about the product's connection. AWS Service Catalog returns this field for the CreateProduct, UpdateProduct, DescribeProductAsAdmin, and SearchProductAsAdmin APIs. This response contains the same fields as the ConnectionParameters request, with the addition of the LastSync response. |
| `SupportDescription` | support_description | `string` | optional, computed, provider-chosen |  | The support information about the product. |
| `SupportEmail` | support_email | `string` | optional, computed, provider-chosen |  | The contact email for product support. |
| `SupportUrl` | support_url | `string` | optional, computed, provider-chosen |  | The contact URL for product support. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |

Supports update: yes

Discovery: supported
