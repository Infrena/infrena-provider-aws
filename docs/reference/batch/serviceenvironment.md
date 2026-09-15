# aws.serviceenvironment

**CloudFormation type:** `AWS::Batch::ServiceEnvironment`

Resource Type definition for AWS::Batch::ServiceEnvironment

Region attribute: `region`

**Import ID:** `<region>/ServiceEnvironmentArn` (AWS::Batch::ServiceEnvironment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CapacityLimits` | capacity_limits | `list` | required |  |  |
| `ServiceEnvironmentArn` | service_environment_arn | `string` | computed |  |  |
| `ServiceEnvironmentName` | service_environment_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ServiceEnvironmentType` | service_environment_type | `string` | required, replaces on change |  |  |
| `State` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
