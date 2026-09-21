# aws.paymentmanager

**CloudFormation type:** `AWS::BedrockAgentCore::PaymentManager`

Resource Type definition for AWS::BedrockAgentCore::PaymentManager

Region attribute: `region`

**Import ID:** `<region>/PaymentManagerArn` (AWS::BedrockAgentCore::PaymentManager)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AuthorizerConfiguration` | authorizer_configuration | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `AuthorizerType` | authorizer_type | `string` | required, replaces on change |  |  |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the payment manager was created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the payment manager |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | The timestamp when the payment manager was last updated |
| `Name` |  | `string` | required, replaces on change |  | The name of the payment manager |
| `PaymentManagerArn` | payment_manager_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the payment manager |
| `PaymentManagerId` | payment_manager_id | `string` | computed |  | The unique identifier for the payment manager |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The ARN of the IAM role for the payment manager |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | Tags to assign to the payment manager |
| `WorkloadIdentityDetails` | workload_identity_details | `map` | computed |  |  |

Supports update: yes

Discovery: supported
