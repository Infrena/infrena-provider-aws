# aws.databrew.project

**CloudFormation type:** `AWS::DataBrew::Project`

Resource schema for AWS::DataBrew::Project.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::DataBrew::Project)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DatasetName` | dataset_name | `string` | required |  | Dataset name |
| `Name` |  | `string` | required, replaces on change |  | Project name |
| `RecipeName` | recipe_name | `string` | required |  | Recipe name |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | Role arn |
| `Sample` |  | `map` | optional, computed, provider-chosen |  | Sample |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
