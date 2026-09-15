# aws.connectionfunction

**CloudFormation type:** `AWS::CloudFront::ConnectionFunction`

A connection function.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::ConnectionFunction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoPublish` | auto_publish | `boolean` | optional, computed, provider-chosen, write-only |  | A flag that determines whether to automatically publish the function to the ``LIVE`` stage when it’s created. To automatically publish to the ``LIVE`` stage, set this property to ``true``. |
| `ConnectionFunctionArn` | connection_function_arn | `string` | computed |  |  |
| `ConnectionFunctionCode` | connection_function_code | `string` | required |  | The code for the connection function. |
| `ConnectionFunctionConfig` | connection_function_config | `map` | required |  | Contains configuration information about a CloudFront function. |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `ETag` | e_tag | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  | The connection function name. |
| `Stage` |  | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A complex type that contains zero or more ``Tag`` elements. |

Supports update: yes

Discovery: supported
