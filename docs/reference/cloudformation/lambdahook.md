# aws.lambdahook

**CloudFormation type:** `AWS::CloudFormation::LambdaHook`

This is a CloudFormation resource for the first-party AWS::Hooks::LambdaHook.

Region attribute: `region`

**Import ID:** `<region>/HookArn` (AWS::CloudFormation::LambdaHook)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | required, replaces on change |  | The typename alias for the hook. |
| `AutoUpdate` | auto_update | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Whether to automatically update the extension in this account and Region when a new minor version is published by the extension publisher. |
| `ExecutionRole` | execution_role | `string` | required |  | IAM Role ARN |
| `FailureMode` | failure_mode | `string` | required |  | Attribute to specify CloudFormation behavior on hook failure. |
| `HookArn` | hook_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the activated hook |
| `HookStatus` | hook_status | `string` | required |  | Attribute to specify which stacks this hook applies to or should get invoked for |
| `LambdaFunction` | lambda_function | `string` | required |  | Amazon Resource Name (ARN), Partial ARN, name, version, or alias of the Lambda function to invoke with this hook. |
| `LoggingConfig` | logging_config | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | Contains logging configuration information for an extension. |
| `StackFilters` | stack_filters | `map` | optional, computed, provider-chosen |  | Filters to allow hooks to target specific stack attributes |
| `TargetFilters` | target_filters | `map` | optional, computed, provider-chosen |  | Attribute to specify which targets should invoke the hook |
| `TargetOperations` | target_operations | `list` | required |  | Which operations should this Hook run against? Resource changes, stacks or change sets. |

Supports update: yes

Discovery: supported
