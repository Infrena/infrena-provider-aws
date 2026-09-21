# aws.actionconnector

**CloudFormation type:** `AWS::QuickSight::ActionConnector`

Definition of the AWS::QuickSight::ActionConnector Resource Type.

Region attribute: `region`

**Import ID:** `<region>/ActionConnectorId|AwsAccountId` (AWS::QuickSight::ActionConnector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionConnectorId` | action_connector_id | `string` | required, replaces on change | aws.actionconnector.ActionConnectorId |  |
| `Arn` |  | `string` | computed |  |  |
| `AuthenticationConfig` | authentication_config | `map` | required, write-only |  |  |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `EnabledActions` | enabled_actions | `list` | computed |  |  |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `Type` | type_value | `string` | required, replaces on change |  |  |
| `VpcConnectionArn` | vpc_connection_arn | `string` | optional, computed, provider-chosen, write-only | aws.quicksight.vpcconnection.Arn |  |

Supports update: yes

Discovery: supported
