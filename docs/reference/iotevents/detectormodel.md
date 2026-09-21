# aws.detectormodel

**CloudFormation type:** `AWS::IoTEvents::DetectorModel`

The AWS::IoTEvents::DetectorModel resource creates a detector model. You create a *detector model* (a model of your equipment or process) using *states*. For each state, you define conditional (Boolean) logic that evaluates the incoming inputs to detect significant events. When an event is detected, it can change the state or trigger custom-built or predefined actions using other AWS services. You can define additional events that trigger actions when entering or exiting a state and, optionally, when a condition is met. For more information, see [How to Use](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/DetectorModelName` (AWS::IoTEvents::DetectorModel)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DetectorModelDefinition` | detector_model_definition | `map` | required |  | Information that defines how a detector operates. |
| `DetectorModelDescription` | detector_model_description | `string` | optional, computed, provider-chosen |  | A brief description of the detector model. |
| `DetectorModelName` | detector_model_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the detector model. |
| `EvaluationMethod` | evaluation_method | `string` | optional, computed, provider-chosen |  | Information about the order in which events are evaluated and how actions are executed. |
| `Key` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The value used to identify a detector instance. When a device or system sends input, a new detector instance with a unique key value is created. ITE can continue to route input to its corresponding detector instance based on this identifying information. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The ARN of the role that grants permission to ITE to perform its operations. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
