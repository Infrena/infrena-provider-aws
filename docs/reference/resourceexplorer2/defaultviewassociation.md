# aws.defaultviewassociation

**CloudFormation type:** `AWS::ResourceExplorer2::DefaultViewAssociation`

Definition of AWS::ResourceExplorer2::DefaultViewAssociation Resource Type

Region attribute: `region`

**Import ID:** `<region>/AssociatedAwsPrincipal` (AWS::ResourceExplorer2::DefaultViewAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociatedAwsPrincipal` | associated_aws_principal | `string` | computed |  | The AWS principal that the default view is associated with, used as the unique identifier for this resource. |
| `ViewArn` | view_arn | `string` | required | aws.resourceexplorer2.view.ViewArn |  |

Supports update: yes

Discovery: not supported
