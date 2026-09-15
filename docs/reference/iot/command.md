# aws.command

**CloudFormation type:** `AWS::IoT::Command`

Represents the resource definition of AWS IoT Command.

Region attribute: `region`

**Import ID:** `<region>/CommandId` (AWS::IoT::Command)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CommandArn` | command_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the command. |
| `CommandId` | command_id | `string` | required, replaces on change | aws.command.CommandId | The unique identifier for the command. |
| `CreatedAt` | created_at | `string` | optional, computed, provider-chosen |  | The date and time when the command was created. |
| `Deprecated` |  | `boolean` | optional, computed, provider-chosen |  | A flag indicating whether the command is deprecated. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the command. |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | The display name for the command. |
| `LastUpdatedAt` | last_updated_at | `string` | optional, computed, provider-chosen, write-only |  | The date and time when the command was last updated. |
| `MandatoryParameters` | mandatory_parameters | `list` | optional, computed, provider-chosen |  | The list of mandatory parameters for the command. |
| `Namespace` |  | `string` | optional, computed, provider-chosen |  | The namespace to which the command belongs. |
| `Payload` |  | `map` | optional, computed, provider-chosen |  | The payload associated with the command. |
| `PayloadTemplate` | payload_template | `string` | optional, computed, provider-chosen, replaces on change |  | The payload template associated with the command. |
| `PendingDeletion` | pending_deletion | `boolean` | optional, computed, provider-chosen |  | A flag indicating whether the command is pending deletion. |
| `Preprocessor` |  | `map` | optional, computed, provider-chosen, replaces on change |  | The command preprocessor configuration. |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | The customer role associated with the command. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to be associated with the command. |

Supports update: yes

Discovery: supported
