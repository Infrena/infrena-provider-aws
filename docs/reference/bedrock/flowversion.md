# aws.flowversion

**CloudFormation type:** `AWS::Bedrock::FlowVersion`

Definition of AWS::Bedrock::FlowVersion Resource Type

Region attribute: `region`

**Import ID:** `<region>/FlowArn|Version` (AWS::Bedrock::FlowVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp. |
| `CustomerEncryptionKeyArn` | customer_encryption_key_arn | `string` | computed |  | A KMS key ARN |
| `Definition` |  | `map` | computed |  | Flow definition |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description of the flow version |
| `ExecutionRoleArn` | execution_role_arn | `string` | computed |  | ARN of a IAM role |
| `FlowArn` | flow_arn | `string` | required, replaces on change | aws.bedrock.flow.Arn | Arn representation of the Flow |
| `FlowId` | flow_id | `string` | computed |  | Identifier for a Flow |
| `Name` |  | `string` | computed |  | Name for the flow |
| `Status` |  | `string` | computed |  | Schema Type for Flow APIs |
| `Version` |  | `string` | computed |  | Numerical Version. |

Supports update: yes

Discovery: supported (parent resource required)
