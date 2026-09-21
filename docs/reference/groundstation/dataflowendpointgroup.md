# aws.dataflowendpointgroup

**CloudFormation type:** `AWS::GroundStation::DataflowEndpointGroup`

AWS Ground Station DataflowEndpointGroup schema for CloudFormation

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::GroundStation::DataflowEndpointGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ContactPostPassDurationSeconds` | contact_post_pass_duration_seconds | `integer` | optional, computed, provider-chosen, replaces on change |  | Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state. |
| `ContactPrePassDurationSeconds` | contact_pre_pass_duration_seconds | `integer` | optional, computed, provider-chosen, replaces on change |  | Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state. |
| `EndpointDetails` | endpoint_details | `list` | required, replaces on change |  |  |
| `Id` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
