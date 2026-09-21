# aws.collectiongroup

**CloudFormation type:** `AWS::OpenSearchServerless::CollectionGroup`

Resource Type definition for AWS::OpenSearchServerless::CollectionGroup

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::OpenSearchServerless::CollectionGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the collection group. |
| `CapacityLimits` | capacity_limits | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the collection group. |
| `Generation` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The generation of Amazon OpenSearch Serverless for the collection group. Valid values are CLASSIC and NEXTGEN. |
| `Id` |  | `string` | computed |  | The unique identifier of the collection group. |
| `Name` |  | `string` | required, replaces on change |  | The name of the collection group. |
| `StandbyReplicas` | standby_replicas | `string` | required, replaces on change |  | Indicates whether standby replicas are used for the collection group. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
