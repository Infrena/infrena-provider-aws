# aws.hoursofoperation

**CloudFormation type:** `AWS::Connect::HoursOfOperation`

Resource Type definition for AWS::Connect::HoursOfOperation

Region attribute: `region`

**Import ID:** `<region>/HoursOfOperationArn` (AWS::Connect::HoursOfOperation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ChildHoursOfOperations` | child_hours_of_operations | `list` | optional, computed, provider-chosen |  | List of child hours of operations. |
| `Config` |  | `list` | required |  | Configuration information for the hours of operation: day, start time, and end time. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the hours of operation. |
| `HoursOfOperationArn` | hours_of_operation_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the hours of operation. |
| `HoursOfOperationOverrides` | hours_of_operation_overrides | `list` | optional, computed, provider-chosen |  | One or more hours of operation overrides assigned to an hour of operation. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `Name` |  | `string` | required |  | The name of the hours of operation. |
| `ParentHoursOfOperations` | parent_hours_of_operations | `list` | optional, computed, provider-chosen |  | List of parent hours of operations. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |
| `TimeZone` | time_zone | `string` | required |  | The time zone of the hours of operation. |

Supports update: yes

Discovery: supported (parent resource required)
