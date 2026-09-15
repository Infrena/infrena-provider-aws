# aws.billscenario

**CloudFormation type:** `AWS::BcmPricingCalculator::BillScenario`

Resource Type definition for AWS::BcmPricingCalculator::BillScenario

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::BcmPricingCalculator::BillScenario)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the bill scenario. |
| `BillInterval` | bill_interval | `map` | computed |  | The time period covered by the bill scenario |
| `CostCategoryGroupSharingPreferenceArn` | cost_category_group_sharing_preference_arn | `string` | optional, computed, provider-chosen |  | The ARN of the cost category group sharing preference |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the bill scenario was created |
| `ExpiresAt` | expires_at | `string` | optional, computed, provider-chosen |  | The timestamp when the bill scenario expires |
| `FailureMessage` | failure_message | `string` | computed |  | The failure message if the bill scenario failed |
| `GroupSharingPreference` | group_sharing_preference | `string` | optional, computed, provider-chosen |  | The group sharing preference for the bill scenario |
| `Id` |  | `string` | computed |  | The unique identifier of the bill scenario |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name of the bill scenario |
| `Status` |  | `string` | computed |  | The status of the bill scenario |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource |

Supports update: yes

Discovery: supported
