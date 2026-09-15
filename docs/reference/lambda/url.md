# aws.url

**CloudFormation type:** `AWS::Lambda::Url`

Resource Type definition for AWS::Lambda::Url

Region attribute: `region`

**Import ID:** `<region>/FunctionArn` (AWS::Lambda::Url)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AuthType` | auth_type | `string` | required |  | Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL. |
| `Cors` |  | `map` | optional, computed, provider-chosen |  |  |
| `FunctionArn` | function_arn | `string` | computed |  | The full Amazon Resource Name (ARN) of the function associated with the Function URL. |
| `FunctionUrl` | function_url | `string` | computed |  | The generated url for this resource. |
| `InvokeMode` | invoke_mode | `string` | optional, computed, provider-chosen |  | The invocation mode for the function's URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED. |
| `Qualifier` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed. |
| `TargetFunctionArn` | target_function_arn | `string` | required, replaces on change | aws.lambda.function.Arn | The Amazon Resource Name (ARN) of the function associated with the Function URL. |

Supports update: yes

Discovery: supported (parent resource required)
