# aws.consumableresource

**CloudFormation type:** `AWS::Batch::ConsumableResource`

Resource Type definition for AWS::Batch::ConsumableResource

Region attribute: `region`

**Import ID:** `<region>/ConsumableResourceArn` (AWS::Batch::ConsumableResource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AvailableQuantity` | available_quantity | `integer` | computed |  | Available Quantity of ConsumableResource. |
| `ConsumableResourceArn` | consumable_resource_arn | `string` | computed |  | ARN of the Consumable Resource. |
| `ConsumableResourceName` | consumable_resource_name | `string` | optional, computed, provider-chosen, replaces on change |  | Name of ConsumableResource. |
| `CreatedAt` | created_at | `integer` | computed |  |  |
| `InUseQuantity` | in_use_quantity | `integer` | computed |  | In-use Quantity of ConsumableResource. |
| `ResourceType` | resource_type | `string` | required, replaces on change |  | Type of Consumable Resource. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |
| `TotalQuantity` | total_quantity | `integer` | required |  | Total Quantity of ConsumableResource. |

Supports update: yes

Discovery: supported
