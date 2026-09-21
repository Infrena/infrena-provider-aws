# aws.input

**CloudFormation type:** `AWS::IoTEvents::Input`

The AWS::IoTEvents::Input resource creates an input. To monitor your devices and processes, they must have a way to get telemetry data into ITE. This is done by sending messages as *inputs* to ITE. For more information, see [How to Use](https://docs.aws.amazon.com/iotevents/latest/developerguide/how-to-use-iotevents.html) in the *Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/InputName` (AWS::IoTEvents::Input)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `InputDefinition` | input_definition | `map` | required |  | The definition of the input. |
| `InputDescription` | input_description | `string` | optional, computed, provider-chosen |  | A brief description of the input. |
| `InputName` | input_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the input. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
