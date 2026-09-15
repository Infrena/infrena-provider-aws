# aws.msk.configuration

**CloudFormation type:** `AWS::MSK::Configuration`

Resource Type definition for AWS::MSK::Configuration

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::MSK::Configuration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `KafkaVersionsList` | kafka_versions_list | `list` | optional, computed, provider-chosen, replaces on change |  |  |
| `LatestRevision` | latest_revision | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `ServerProperties` | server_properties | `string` | required, write-only |  |  |

Supports update: yes

Discovery: supported
