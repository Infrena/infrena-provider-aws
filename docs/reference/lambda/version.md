# aws.version

**CloudFormation type:** `AWS::Lambda::Version`

Resource Type definition for AWS::Lambda::Version

Region attribute: `region`

**Import ID:** `<region>/FunctionArn` (AWS::Lambda::Version)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CodeSha256` | code_sha256 | `string` | optional, computed, provider-chosen, replaces on change |  | Only publish a version if the hash value matches the value that's specified. Use this option to avoid publishing a version if the function code has changed since you last updated it. Updates are not supported for this property. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description for the version to override the description in the function configuration. Updates are not supported for this property. |
| `FunctionArn` | function_arn | `string` | computed |  | The ARN of the version. |
| `FunctionName` | function_name | `string` | required, replaces on change |  | The name of the Lambda function. |
| `FunctionScalingConfig` | function_scaling_config | `map` | optional, computed, provider-chosen |  | Configuration that defines the scaling behavior for a Lambda Managed Instances function, including the minimum and maximum number of execution environments that can be provisioned. |
| `ProvisionedConcurrencyConfig` | provisioned_concurrency_config | `map` | optional, computed, provider-chosen, replaces on change |  | A provisioned concurrency configuration for a function's version. |
| `RuntimePolicy` | runtime_policy | `map` | optional, computed, provider-chosen, replaces on change |  | Runtime Management Config of a function. |
| `Version` |  | `string` | computed |  | The version number. |

Supports update: yes

Discovery: supported (parent resource required)
