# aws.flowalias

**CloudFormation type:** `AWS::Bedrock::FlowAlias`

Definition of AWS::Bedrock::FlowAlias Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn|FlowArn` (AWS::Bedrock::FlowAlias)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn of the Flow Alias |
| `ConcurrencyConfiguration` | concurrency_configuration | `map` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the Resource. |
| `FlowArn` | flow_arn | `string` | required, replaces on change | aws.bedrock.flow.Arn | Arn representation of the Flow |
| `FlowId` | flow_id | `string` | computed |  | Identifier for a flow resource. |
| `Id` |  | `string` | computed |  | Id for a Flow Alias generated at the server side. |
| `Name` |  | `string` | required |  | Name for a resource. |
| `RoutingConfiguration` | routing_configuration | `list` | required |  | Routing configuration for a Flow alias. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp. |

Supports update: yes

Discovery: supported (parent resource required)
