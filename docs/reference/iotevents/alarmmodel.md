# aws.alarmmodel

**CloudFormation type:** `AWS::IoTEvents::AlarmModel`

Represents an alarm model to monitor an ITE input attribute. You can use the alarm to get notified when the value is outside a specified range. For more information, see [Create an alarm model](https://docs.aws.amazon.com/iotevents/latest/developerguide/create-alarms.html) in the *Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/AlarmModelName` (AWS::IoTEvents::AlarmModel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AlarmCapabilities` | alarm_capabilities | `map` | optional, computed, provider-chosen |  | Contains the configuration information of alarm state changes. |
| `AlarmEventActions` | alarm_event_actions | `map` | optional, computed, provider-chosen |  | Contains information about one or more alarm actions. |
| `AlarmModelDescription` | alarm_model_description | `string` | optional, computed, provider-chosen |  | The description of the alarm model. |
| `AlarmModelName` | alarm_model_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the alarm model. |
| `AlarmRule` | alarm_rule | `map` | required |  | Defines when your alarm is invoked. |
| `Key` |  | `string` | optional, computed, provider-chosen, replaces on change |  | An input attribute used as a key to create an alarm. ITE routes [inputs](https://docs.aws.amazon.com/iotevents/latest/apireference/API_Input.html) associated with this key to the alarm. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The ARN of the IAM role that allows the alarm to perform actions and access AWS resources. For more information, see [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the *General Reference*. |
| `Severity` |  | `integer` | optional, computed, provider-chosen |  | A non-negative integer that reflects the severity level of the alarm. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the alarm model. The tags help you manage the alarm model. For more information, see [Tagging your resources](https://docs.aws.amazon.com/iotevents/latest/developerguide/tagging-iotevents.html) in the *Developer Guide*. |

Supports update: yes

Discovery: supported
