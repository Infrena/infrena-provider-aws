# aws.simulation

**CloudFormation type:** `AWS::SimSpaceWeaver::Simulation`

AWS::SimSpaceWeaver::Simulation resource creates an AWS Simulation.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::SimSpaceWeaver::Simulation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DescribePayload` | describe_payload | `string` | computed |  | Json object with all simulation details |
| `MaximumDuration` | maximum_duration | `string` | optional, computed, provider-chosen, replaces on change |  | The maximum running time of the simulation. |
| `Name` |  | `string` | required, replaces on change |  | The name of the simulation. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | Role ARN. |
| `SchemaS3Location` | schema_s3_location | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SnapshotS3Location` | snapshot_s3_location | `map` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
