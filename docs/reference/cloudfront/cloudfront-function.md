# aws.cloudfront.function

**CloudFormation type:** `AWS::CloudFront::Function`

Creates a CF function.

Global type (no region attribute)

**Import ID:** `global/FunctionARN` (AWS::CloudFront::Function)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutoPublish` | auto_publish | `boolean` | optional, computed, provider-chosen, write-only |  | A flag that determines whether to automatically publish the function to the ``LIVE`` stage when it’s created. To automatically publish to the ``LIVE`` stage, set this property to ``true``. |
| `FunctionARN` | function_arn | `string` | computed |  |  |
| `FunctionCode` | function_code | `string` | required |  | The function code. For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide*. |
| `FunctionConfig` | function_config | `map` | required |  | Contains configuration information about a CloudFront function. |
| `FunctionMetadata` | function_metadata | `map` | optional, computed, provider-chosen |  | Contains metadata about a CloudFront function. |
| `Name` |  | `string` | required, replaces on change |  | A name to identify the function. |
| `Stage` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A complex type that contains zero or more ``Tag`` elements. |

Supports update: yes

Discovery: supported
