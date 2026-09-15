# aws.sagemaker.project

**CloudFormation type:** `AWS::SageMaker::Project`

Resource Type definition for AWS::SageMaker::Project

Region attribute: `region`

**Import ID:** `<region>/ProjectArn` (AWS::SageMaker::Project)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The time at which the project was created. |
| `ProjectArn` | project_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Project. |
| `ProjectDescription` | project_description | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the project. |
| `ProjectId` | project_id | `string` | computed |  | Project Id. |
| `ProjectName` | project_name | `string` | required, replaces on change |  | The name of the project. |
| `ProjectStatus` | project_status | `string` | computed |  | The status of a project. |
| `ServiceCatalogProvisionedProductDetails` | service_catalog_provisioned_product_details | `map` | optional, computed, provider-chosen |  | Provisioned ServiceCatalog  Details |
| `ServiceCatalogProvisioningDetails` | service_catalog_provisioning_details | `map` | optional, computed, provider-chosen, replaces on change |  | Input ServiceCatalog Provisioning Details |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | An array of key-value pairs to apply to this resource. |
| `TemplateProviderDetails` | template_provider_details | `list` | optional, computed, provider-chosen, replaces on change |  | An array of template providers associated with the project. |

Supports update: yes

Discovery: supported
