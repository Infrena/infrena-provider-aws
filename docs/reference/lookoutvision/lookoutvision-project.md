# aws.lookoutvision.project

**CloudFormation type:** `AWS::LookoutVision::Project`

The AWS::LookoutVision::Project type creates an Amazon Lookout for Vision project. A project is a grouping of the resources needed to create and manage a Lookout for Vision model.

Region attribute: `region`

**Import ID:** `<region>/ProjectName` (AWS::LookoutVision::Project)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ProjectName` | project_name | `string` | required, replaces on change |  | The name of the Amazon Lookout for Vision project |

Supports update: yes

Discovery: supported
