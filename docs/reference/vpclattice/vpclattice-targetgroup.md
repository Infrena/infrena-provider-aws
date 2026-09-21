# aws.vpclattice.targetgroup

**CloudFormation type:** `AWS::VpcLattice::TargetGroup`

A target group is a collection of targets, or compute resources, that run your application or service. A target group can only be used by a single service.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::VpcLattice::TargetGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Config` |  | `map` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Targets` |  | `list` | optional, computed, provider-chosen |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: supported
