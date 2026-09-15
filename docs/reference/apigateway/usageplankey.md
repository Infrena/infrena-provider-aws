# aws.usageplankey

**CloudFormation type:** `AWS::ApiGateway::UsagePlanKey`

The ``AWS::ApiGateway::UsagePlanKey`` resource associates an API key with a usage plan. This association determines which users the usage plan is applied to.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ApiGateway::UsagePlanKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `KeyId` | key_id | `string` | required, replaces on change |  | The Id of the UsagePlanKey resource. |
| `KeyType` | key_type | `string` | required, replaces on change |  |  |
| `UsagePlanId` | usage_plan_id | `string` | required, replaces on change | aws.usageplan.Id | The Id of the UsagePlan resource representing the usage plan containing the UsagePlanKey resource representing a plan customer. |

Supports update: no

Discovery: supported (parent resource required)
