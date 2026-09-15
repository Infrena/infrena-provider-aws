# aws.dataflowendpointgroupv2

**CloudFormation type:** `AWS::GroundStation::DataflowEndpointGroupV2`

Resource Type definition for AWS Ground Station DataflowEndpointGroupV2

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::GroundStation::DataflowEndpointGroupV2)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ContactPostPassDurationSeconds` | contact_post_pass_duration_seconds | `integer` | optional, computed, provider-chosen, replaces on change |  | Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state. |
| `ContactPrePassDurationSeconds` | contact_pre_pass_duration_seconds | `integer` | optional, computed, provider-chosen, replaces on change |  | Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state. |
| `EndpointDetails` | endpoint_details | `list` | computed |  |  |
| `Endpoints` |  | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Id` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
