# aws.eventstream

**CloudFormation type:** `AWS::CustomerProfiles::EventStream`

An Event Stream resource of Amazon Connect Customer Profiles

Region attribute: `region`

**Import ID:** `<region>/DomainName|EventStreamName` (AWS::CustomerProfiles::EventStream)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the export was created. |
| `DestinationDetails` | destination_details | `map` | computed |  | Details regarding the Kinesis stream. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The unique name of the domain. |
| `EventStreamArn` | event_stream_arn | `string` | computed |  | A unique identifier for the event stream. |
| `EventStreamName` | event_stream_name | `string` | required, replaces on change |  | The name of the event stream. |
| `State` |  | `string` | computed |  | The operational state of destination stream for export. |
| `Tags` |  | `map` | tags map |  | The tags used to organize, track, or control access for this resource. |
| `Uri` |  | `string` | required, replaces on change |  | The StreamARN of the destination to deliver profile events to. For example, arn:aws:kinesis:region:account-id:stream/stream-name |

Supports update: yes

Discovery: supported (parent resource required)
