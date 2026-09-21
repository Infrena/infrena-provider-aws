# aws.guardhook

**CloudFormation type:** `AWS::CloudFormation::GuardHook`

This is a CloudFormation resource for activating the first-party AWS::Hooks::GuardHook.

Region attribute: `region`

**Import ID:** `<region>/HookArn` (AWS::CloudFormation::GuardHook)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `string` | required, replaces on change |  | The typename alias for the hook. |
| `ExecutionRole` | execution_role | `string` | required, replaces on change |  | IAM Role ARN |
| `FailureMode` | failure_mode | `string` | required |  | Attribute to specify CloudFormation behavior on hook failure. |
| `HookArn` | hook_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the activated hook |
| `HookStatus` | hook_status | `string` | required |  | Attribute to specify which stacks this hook applies to or should get invoked for |
| `LogBucket` | log_bucket | `string` | optional, computed, provider-chosen |  | S3 Bucket where the guard validate report will be uploaded to |
| `Options` |  | `map` | optional, computed, provider-chosen |  |  |
| `RuleLocation` | rule_location | `map` | required |  | S3 Source Location for the Guard files. |
| `StackFilters` | stack_filters | `map` | optional, computed, provider-chosen |  | Filters to allow hooks to target specific stack attributes |
| `TargetFilters` | target_filters | `map` | optional, computed, provider-chosen |  | Attribute to specify which targets should invoke the hook |
| `TargetOperations` | target_operations | `list` | required |  | Which operations should this Hook run against? Resource changes, stacks or change sets. |

Supports update: yes

Discovery: supported
