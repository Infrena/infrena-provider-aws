# aws.wickr.network

**CloudFormation type:** `AWS::Wickr::Network`

Resource Type definition for AWS::Wickr::Network. Creates and manages an AWS Wickr network for secure enterprise communications.

Region attribute: `region`

**Import ID:** `<region>/NetworkArn` (AWS::Wickr::Network)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessLevel` | access_level | `string` | required, replaces on change |  | The access level of the network, which determines available features and capabilities. |
| `AwsAccountId` | aws_account_id | `string` | computed |  | The AWS account ID that owns the network. |
| `MigrationState` | migration_state | `integer` | computed |  | The SSO redirect URI migration state. Values: 0 (not started), 1 (in progress), or 2 (completed). |
| `NetworkArn` | network_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the network. |
| `NetworkId` | network_id | `string` | computed |  | The unique identifier of the network. |
| `NetworkName` | network_name | `string` | required |  | The name of the network. Must be between 1 and 20 characters. |
| `Standing` |  | `integer` | computed |  | The current standing or status of the network. |

Supports update: yes

Discovery: supported
