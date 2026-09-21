# aws.glue.blueprint

**CloudFormation type:** `AWS::Glue::Blueprint`

Resource Type definition for AWS::Glue::Blueprint. Registers a blueprint with AWS Glue.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Glue::Blueprint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the blueprint. |
| `BlueprintLocation` | blueprint_location | `string` | required |  | Specifies a path in Amazon S3 where the blueprint is published. |
| `CreatedOn` | created_on | `string` | computed |  | The date and time the blueprint was registered. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the blueprint. |
| `LastModifiedOn` | last_modified_on | `string` | computed |  | The date and time the blueprint was last modified. |
| `Name` |  | `string` | required, replaces on change |  | The name of the blueprint. |
| `ParameterSpec` | parameter_spec | `string` | computed |  | A JSON string that indicates the list of parameter specifications for the blueprint. |
| `Status` |  | `string` | computed |  | The status of the blueprint registration. |
| `Tags` |  | `map` | tags map |  | The tags to be applied to this blueprint. |

Supports update: yes

Discovery: supported
