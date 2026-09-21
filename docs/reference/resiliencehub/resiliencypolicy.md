# aws.resiliencypolicy

**CloudFormation type:** `AWS::ResilienceHub::ResiliencyPolicy`

Resource Type Definition for Resiliency Policy.

Region attribute: `region`

**Import ID:** `<region>/PolicyArn` (AWS::ResilienceHub::ResiliencyPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataLocationConstraint` | data_location_constraint | `string` | optional, computed, provider-chosen |  | Data Location Constraint of the Policy. |
| `Policy` |  | `map` | required |  |  |
| `PolicyArn` | policy_arn | `string` | computed |  | Amazon Resource Name (ARN) of the Resiliency Policy. |
| `PolicyDescription` | policy_description | `string` | optional, computed, provider-chosen |  | Description of Resiliency Policy. |
| `PolicyName` | policy_name | `string` | required |  | Name of Resiliency Policy. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `Tier` |  | `string` | required |  | Resiliency Policy Tier. |

Supports update: yes

Discovery: supported
