# aws.rekognition.project

**CloudFormation type:** `AWS::Rekognition::Project`

The AWS::Rekognition::Project type creates an Amazon Rekognition CustomLabels Project. A project is a grouping of the resources needed to create and manage Dataset and ProjectVersions.

Region attribute: `region`

**Import ID:** `<region>/ProjectName` (AWS::Rekognition::Project)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ProjectName` | project_name | `string` | required, replaces on change |  | The name of the project |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
