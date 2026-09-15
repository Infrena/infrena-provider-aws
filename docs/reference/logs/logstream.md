# aws.logstream

**CloudFormation type:** `AWS::Logs::LogStream`

Resource Type definition for AWS::Logs::LogStream

Region attribute: `region`

**Import ID:** `<region>/LogGroupName|LogStreamName` (AWS::Logs::LogStream)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LogGroupName` | log_group_name | `string` | required, replaces on change |  | The name of the log group where the log stream is created. |
| `LogStreamName` | log_stream_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the log stream. The name must be unique wihtin the log group. |

Supports update: no

Discovery: supported (parent resource required)
