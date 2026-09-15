# aws.cell

**CloudFormation type:** `AWS::Route53RecoveryReadiness::Cell`

The API Schema for AWS Route53 Recovery Readiness Cells.

Region attribute: `region`

**Import ID:** `<region>/CellName` (AWS::Route53RecoveryReadiness::Cell)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CellArn` | cell_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the cell. |
| `CellName` | cell_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the cell to create. |
| `Cells` |  | `list` | optional, computed, provider-chosen |  | A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions. |
| `ParentReadinessScopes` | parent_readiness_scopes | `list` | computed |  | The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN) or a recovery group ARN. This is a list but currently can have only one element. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A collection of tags associated with a resource |

Supports update: yes

Discovery: supported
