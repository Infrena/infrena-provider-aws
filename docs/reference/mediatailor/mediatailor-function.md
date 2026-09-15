# aws.mediatailor.function

**CloudFormation type:** `AWS::MediaTailor::Function`

Resource Type definition for AWS::MediaTailor::Function

Region attribute: `region`

**Import ID:** `<region>/FunctionId` (AWS::MediaTailor::Function)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the function. |
| `CustomOutputConfiguration` | custom_output_configuration | `map` | optional, computed, provider-chosen |  | Configuration for custom output functions. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the function. |
| `FunctionId` | function_id | `string` | required, replaces on change | aws.mediatailor.function.FunctionId | The unique identifier for the function. |
| `FunctionType` | function_type | `string` | required |  | The type of the function. Determines which configuration object is used. |
| `HttpRequestConfiguration` | http_request_configuration | `map` | optional, computed, provider-chosen |  | Configuration for HTTP request functions. |
| `SequentialExecutorConfiguration` | sequential_executor_configuration | `map` | optional, computed, provider-chosen |  | Configuration for sequential executor functions. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to assign to the function resource. |

Supports update: yes

Discovery: supported
