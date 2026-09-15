# aws.cloudconnector

**CloudFormation type:** `AWS::SSM::CloudConnector`

Resource Type definition for AWS::SSM::CloudConnector. Enables AWS Systems Manager to manage resources in external cloud providers.

Region attribute: `region`

**Import ID:** `<region>/CloudConnectorArn` (AWS::SSM::CloudConnector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CloudConnectorArn` | cloud_connector_arn | `string` | computed |  | The ARN of the cloud connector. |
| `CloudConnectorId` | cloud_connector_id | `string` | computed |  | The unique identifier of the cloud connector. |
| `ConfigConnectorArn` | config_connector_arn | `string` | required, replaces on change |  | The ARN of the AWS Config connector. |
| `Configuration` |  | `map` | required |  | The configuration for the cloud connector. Currently supports Azure. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the cloud connector was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the cloud connector. |
| `DisplayName` | display_name | `string` | required |  | The display name of the cloud connector. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The IAM role ARN used by the cloud connector. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to apply to the cloud connector. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the cloud connector was last updated. |

Supports update: yes

Discovery: supported
