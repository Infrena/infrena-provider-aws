# aws.glue.session

**CloudFormation type:** `AWS::Glue::Session`

Resource Type definition for AWS::Glue::Session. Sessions provide an on-demand, serverless Apache Spark runtime environment for building, testing, and running data preparation and analytics applications.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Glue::Session)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the session. |
| `Command` |  | `map` | required, replaces on change |  | The SessionCommand that runs the job. |
| `Connections` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies the connections used by the session. |
| `CreatedOn` | created_on | `string` | computed |  | The time and date when the session was created. |
| `DefaultArguments` | default_arguments | `map` | optional, computed, provider-chosen, replaces on change |  | A map array of key-value pairs. Max is 75 pairs. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the session. |
| `GlueVersion` | glue_version | `string` | optional, computed, provider-chosen, replaces on change |  | The Glue version determines the versions of Apache Spark and Python that Glue supports. The GlueVersion must be greater than 2.0. |
| `Id` |  | `string` | required, replaces on change |  | The ID of the session. |
| `IdleTimeout` | idle_timeout | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of minutes when idle before session times out. Default is the value of Timeout. |
| `MaxCapacity` | max_capacity | `float` | optional, computed, provider-chosen, replaces on change |  | The number of Glue data processing units (DPUs) that can be allocated when the job runs. |
| `NumberOfWorkers` | number_of_workers | `integer` | optional, computed, provider-chosen, replaces on change |  | The number of workers of a defined WorkerType to use for the session. |
| `Progress` |  | `float` | computed |  | The code execution progress of the session. |
| `RequestOrigin` | request_origin | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The origin of the request. |
| `Role` |  | `string` | required, replaces on change |  | The IAM Role ARN. |
| `SecurityConfiguration` | security_configuration | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the SecurityConfiguration structure to be used with the session. |
| `Status` |  | `string` | computed |  | The session status. |
| `Tags` |  | `map` | tags map |  | The tags belonging to the session. |
| `Timeout` |  | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | The number of minutes before session times out. |
| `WorkerType` | worker_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of predefined worker that is allocated when a session runs. |

Supports update: yes

Discovery: supported
