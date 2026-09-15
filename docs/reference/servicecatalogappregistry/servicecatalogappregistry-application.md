# aws.servicecatalogappregistry.application

**CloudFormation type:** `AWS::ServiceCatalogAppRegistry::Application`

Resource Schema for AWS::ServiceCatalogAppRegistry::Application

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceCatalogAppRegistry::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationName` | application_name | `string` | computed |  | The name of the application. |
| `ApplicationTagKey` | application_tag_key | `string` | computed |  | The key of the AWS application tag, which is awsApplication. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value. |
| `ApplicationTagValue` | application_tag_value | `string` | computed |  | The value of the AWS application tag, which is the identifier of an associated resource. Applications created before 11/13/2023 or applications without the AWS application tag resource group return no value. |
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the application. |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | required |  | The name of the application. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
