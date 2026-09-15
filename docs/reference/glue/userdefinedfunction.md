# aws.userdefinedfunction

**CloudFormation type:** `AWS::Glue::UserDefinedFunction`

Represents a user-defined function (UDF) definition in the AWS Glue Data Catalog.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Glue::UserDefinedFunction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the user-defined function. |
| `ClassName` | class_name | `string` | optional, computed, provider-chosen |  | The Java class that contains the function code. |
| `DatabaseName` | database_name | `string` | required, replaces on change |  | The name of the catalog database in which the function is located. |
| `FunctionName` | function_name | `string` | required, replaces on change |  | The name of the function. |
| `FunctionType` | function_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of the function. |
| `OwnerName` | owner_name | `string` | optional, computed, provider-chosen |  | The owner of the function. |
| `OwnerType` | owner_type | `string` | optional, computed, provider-chosen |  | The owner type. |
| `ResourceUris` | resource_uris | `list` | optional, computed, provider-chosen |  | The resource URIs for the function. |

Supports update: yes

Discovery: supported (parent resource required)
