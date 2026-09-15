# aws.cloudwatchalarmtemplategroup

**CloudFormation type:** `AWS::MediaLive::CloudWatchAlarmTemplateGroup`

Definition of AWS::MediaLive::CloudWatchAlarmTemplateGroup Resource Type

Region attribute: `region`

**Import ID:** `<region>/Identifier` (AWS::MediaLive::CloudWatchAlarmTemplateGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | A cloudwatch alarm template group's ARN (Amazon Resource Name) |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A resource's optional description. |
| `Id` |  | `string` | computed |  | A cloudwatch alarm template group's id. AWS provided template groups have ids that start with `aws-` |
| `Identifier` |  | `string` | computed |  |  |
| `ModifiedAt` | modified_at | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  | A resource's name. Names must be unique within the scope of a resource type in a specific region. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Represents the tags associated with a resource. |

Supports update: yes

Discovery: supported
