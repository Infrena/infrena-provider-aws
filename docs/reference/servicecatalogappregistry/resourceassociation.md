# aws.resourceassociation

**CloudFormation type:** `AWS::ServiceCatalogAppRegistry::ResourceAssociation`

Resource Schema for AWS::ServiceCatalogAppRegistry::ResourceAssociation

Region attribute: `region`

**Import ID:** `<region>/ApplicationArn|ResourceArn|ResourceType` (AWS::ServiceCatalogAppRegistry::ResourceAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Application` |  | `string` | required, replaces on change |  | The name or the Id of the Application. |
| `ApplicationArn` | application_arn | `string` | computed |  |  |
| `Resource` |  | `string` | required, replaces on change |  | The name or the Id of the Resource. |
| `ResourceArn` | resource_arn | `string` | computed |  |  |
| `ResourceType` | resource_type | `string` | required, replaces on change |  | The type of the CFN Resource for now it's enum CFN_STACK. |

Supports update: no

Discovery: supported (parent resource required)
