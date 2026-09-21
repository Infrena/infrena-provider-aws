# aws.connect.queue

**CloudFormation type:** `AWS::Connect::Queue`

Resource Type definition for AWS::Connect::Queue

Region attribute: `region`

**Import ID:** `<region>/QueueArn` (AWS::Connect::Queue)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalEmailAddresses` | additional_email_addresses | `list` | optional, computed, provider-chosen |  | The email addresses that agents can use when replying to or initiating email contacts |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the queue. |
| `HoursOfOperationArn` | hours_of_operation_arn | `string` | required | aws.hoursofoperation.HoursOfOperationArn | The identifier for the hours of operation. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `LastModifiedRegion` | last_modified_region | `string` | computed |  | The AWS Region where this resource was last modified. |
| `LastModifiedTime` | last_modified_time | `float` | computed |  | The timestamp when this resource was last modified. |
| `MaxContacts` | max_contacts | `integer` | optional, computed, provider-chosen |  | The maximum number of contacts that can be in the queue before it is considered full. |
| `Name` |  | `string` | required |  | The name of the queue. |
| `OutboundCallerConfig` | outbound_caller_config | `map` | optional, computed, provider-chosen |  | The outbound caller ID name, number, and outbound whisper flow. |
| `OutboundEmailConfig` | outbound_email_config | `map` | optional, computed, provider-chosen |  | The outbound email address ID. |
| `QueueArn` | queue_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the queue. |
| `QuickConnectArns` | quick_connect_arns | `list` | optional, computed, provider-chosen | aws.quickconnect.QuickConnectArn | The quick connects available to agents who are working the queue. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the queue. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | computed |  | The type of queue. |

Supports update: yes

Discovery: supported (parent resource required)
