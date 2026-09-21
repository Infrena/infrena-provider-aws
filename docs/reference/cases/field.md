# aws.field

**CloudFormation type:** `AWS::Cases::Field`

A field in the Cases domain. This field is used to define the case object model (that is, defines what data can be captured on cases) in a Cases domain.

Region attribute: `region`

**Import ID:** `<region>/FieldArn` (AWS::Cases::Field)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Attributes` |  | `map` | optional, computed, provider-chosen |  | Union of field attributes |
| `CreatedTime` | created_time | `string` | computed |  | The time at which the field was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description explaining the purpose and usage of this field in cases. Helps agents and administrators understand what information should be captured in this field. |
| `DomainId` | domain_id | `string` | optional, computed, provider-chosen, replaces on change | aws.cases.domain.DomainId | The unique identifier of the Cases domain. |
| `FieldArn` | field_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the field. |
| `FieldId` | field_id | `string` | computed |  | The unique identifier of a field. |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | The time at which the field was created or last modified. |
| `Name` |  | `string` | required |  | The display name of the field as it appears to agents in the case interface. Should be descriptive and user-friendly (e.g., 'Customer Priority Level', 'Issue Category'). |
| `Namespace` |  | `string` | computed |  | Indicates whether this is a System field (predefined by AWS) or a Custom field (created by your organization). System fields cannot be modified or deleted. |
| `Tags` |  | `map` | tags map |  | The tags that you attach to this field. |
| `Type` | type_value | `string` | required, replaces on change |  | The data type of the field, which determines validation rules, input constraints, and display format. Each type has specific constraints: Text (string input), Number (numeric values), Boolean (true/false), DateTime (date/time picker), SingleSelect (dropdown options), Url (URL validation), User (Amazon Connect user selection). |

Supports update: yes

Discovery: supported (parent resource required)
