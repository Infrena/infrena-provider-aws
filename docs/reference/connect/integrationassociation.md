# aws.integrationassociation

**CloudFormation type:** `AWS::Connect::IntegrationAssociation`

Resource Type definition for AWS::Connect::IntegrationAssociation

Region attribute: `region`

**Import ID:** `<region>/InstanceId|IntegrationType|IntegrationArn` (AWS::Connect::IntegrationAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `InstanceId` | instance_id | `string` | required, replaces on change | aws.connect.instance.Id | Amazon Connect instance identifier |
| `IntegrationArn` | integration_arn | `string` | required, replaces on change |  | ARN of Integration being associated with the instance |
| `IntegrationAssociationId` | integration_association_id | `string` | computed |  | Identifier of the association with Connect Instance |
| `IntegrationType` | integration_type | `string` | required, replaces on change |  | Specifies the integration type to be associated with the instance |
| `Tags` |  | `map` | tags map |  | The tags used to organize, track, or control access for this resource. |

Supports update: yes

Discovery: supported (parent resource required)
