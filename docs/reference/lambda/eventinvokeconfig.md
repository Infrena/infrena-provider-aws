# aws.eventinvokeconfig

**CloudFormation type:** `AWS::Lambda::EventInvokeConfig`

The AWS::Lambda::EventInvokeConfig resource configures options for asynchronous invocation on a version or an alias.

Region attribute: `region`

**Import ID:** `<region>/FunctionName|Qualifier` (AWS::Lambda::EventInvokeConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DestinationConfig` | destination_config | `map` | optional, computed, provider-chosen |  | A destination for events after they have been sent to a function for processing. |
| `FunctionName` | function_name | `string` | required, replaces on change |  | The name of the Lambda function. |
| `MaximumEventAgeInSeconds` | maximum_event_age_in_seconds | `integer` | optional, computed, provider-chosen |  | The maximum age of a request that Lambda sends to a function for processing. |
| `MaximumRetryAttempts` | maximum_retry_attempts | `integer` | optional, computed, provider-chosen |  | The maximum number of times to retry when the function returns an error. |
| `Qualifier` |  | `string` | required, replaces on change |  | The identifier of a version or alias. |

Supports update: yes

Discovery: supported (parent resource required)
