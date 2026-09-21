# aws.refactorspaces.application

**CloudFormation type:** `AWS::RefactorSpaces::Application`

Definition of AWS::RefactorSpaces::Application Resource Type

Region attribute: `region`

**Import ID:** `<region>/EnvironmentIdentifier|ApplicationIdentifier` (AWS::RefactorSpaces::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiGatewayId` | api_gateway_id | `string` | computed |  |  |
| `ApiGatewayProxy` | api_gateway_proxy | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `ApplicationIdentifier` | application_identifier | `string` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `EnvironmentIdentifier` | environment_identifier | `string` | required, replaces on change |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `NlbArn` | nlb_arn | `string` | computed |  |  |
| `NlbName` | nlb_name | `string` | computed |  |  |
| `ProxyType` | proxy_type | `string` | required, replaces on change |  |  |
| `ProxyUrl` | proxy_url | `string` | computed |  |  |
| `StageName` | stage_name | `string` | computed |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen, replaces on change |  | Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair. |
| `VpcId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId |  |
| `VpcLinkId` | vpc_link_id | `string` | computed |  |  |

Supports update: no

Discovery: supported (parent resource required)
