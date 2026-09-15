# aws.refactorspaces.service

**CloudFormation type:** `AWS::RefactorSpaces::Service`

Definition of AWS::RefactorSpaces::Service Resource Type

Region attribute: `region`

**Import ID:** `<region>/EnvironmentIdentifier|ApplicationIdentifier|ServiceIdentifier` (AWS::RefactorSpaces::Service)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationIdentifier` | application_identifier | `string` | required, replaces on change |  |  |
| `Arn` |  | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `EndpointType` | endpoint_type | `string` | required, replaces on change, write-only |  |  |
| `EnvironmentIdentifier` | environment_identifier | `string` | required, replaces on change |  |  |
| `LambdaEndpoint` | lambda_endpoint | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Name` |  | `string` | required, replaces on change, write-only |  |  |
| `ServiceIdentifier` | service_identifier | `string` | computed |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen, replaces on change |  | Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair. |
| `UrlEndpoint` | url_endpoint | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `VpcId` | vpc_id | `string` | optional, computed, provider-chosen, replaces on change, write-only | aws.vpc.VpcId |  |

Supports update: no

Discovery: supported (parent resource required)
