# aws.codeinterpretercustom

**CloudFormation type:** `AWS::BedrockAgentCore::CodeInterpreterCustom`

Resource definition for AWS::BedrockAgentCore::CodeInterpreterCustom

Region attribute: `region`

**Import ID:** `<region>/CodeInterpreterId` (AWS::BedrockAgentCore::CodeInterpreterCustom)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Certificates` |  | `list` | optional, computed, provider-chosen, replaces on change |  | List of root CA certificates. |
| `CodeInterpreterArn` | code_interpreter_arn | `string` | computed |  | The ARN of a CodeInterpreter resource. |
| `CodeInterpreterId` | code_interpreter_id | `string` | computed |  | The id of the code interpreter. |
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the code interpreter was created. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description of the code interpreter. |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The ARN of the IAM role. |
| `FailureReason` | failure_reason | `string` | computed |  | The reason for failure if the code interpreter creation or operation failed. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | Timestamp when the code interpreter was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the sandbox. |
| `NetworkConfiguration` | network_configuration | `map` | required, replaces on change |  | Network configuration for code interpreter |
| `Status` |  | `string` | computed |  | Status of Code interpreter |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |

Supports update: yes

Discovery: supported
