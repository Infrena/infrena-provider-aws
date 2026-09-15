# aws.discoverer

**CloudFormation type:** `AWS::EventSchemas::Discoverer`

Resource Type definition for AWS::EventSchemas::Discoverer

Region attribute: `region`

**Import ID:** `<region>/DiscovererArn` (AWS::EventSchemas::Discoverer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CrossAccount` | cross_account | `boolean` | optional, computed, provider-chosen |  | Defines whether event schemas from other accounts are discovered. Default is True. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the discoverer. |
| `DiscovererArn` | discoverer_arn | `string` | computed |  | The ARN of the discoverer. |
| `DiscovererId` | discoverer_id | `string` | computed |  | The Id of the discoverer. |
| `SourceArn` | source_arn | `string` | required, replaces on change |  | The ARN of the event bus. |
| `State` |  | `string` | computed |  | Defines the current state of the discoverer. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags associated with the resource. |

Supports update: yes

Discovery: supported
