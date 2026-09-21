# aws.apidestination

**CloudFormation type:** `AWS::Events::ApiDestination`

Resource Type definition for AWS::Events::ApiDestination.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Events::ApiDestination)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The arn of the api destination. |
| `ArnForPolicy` | arn_for_policy | `string` | computed |  | The arn of the api destination to be used in IAM policies. |
| `ConnectionArn` | connection_arn | `string` | required | aws.events.connection.Arn | The arn of the connection. |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `HttpMethod` | http_method | `string` | required |  |  |
| `InvocationEndpoint` | invocation_endpoint | `string` | required |  | Url endpoint to invoke. |
| `InvocationRateLimitPerSecond` | invocation_rate_limit_per_second | `integer` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the apiDestination. |

Supports update: yes

Discovery: supported
