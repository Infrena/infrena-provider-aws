# aws.framework

**CloudFormation type:** `AWS::Backup::Framework`

Contains detailed information about a framework. Frameworks contain controls, which evaluate and report on your backup events and resources. Frameworks generate daily compliance results.

Region attribute: `region`

**Import ID:** `<region>/FrameworkArn` (AWS::Backup::Framework)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The date and time that a framework is created, in ISO 8601 representation. The value of CreationTime is accurate to milliseconds. For example, 2020-07-10T15:00:00.000-08:00 represents the 10th of July 2020 at 3:00 PM 8 hours behind UTC. |
| `DeploymentStatus` | deployment_status | `string` | computed |  | The deployment status of a framework. The statuses are: `CREATE_IN_PROGRESS \| UPDATE_IN_PROGRESS \| DELETE_IN_PROGRESS \| COMPLETED \| FAILED` |
| `FrameworkArn` | framework_arn | `string` | computed |  | An Amazon Resource Name (ARN) that uniquely identifies Framework as a resource |
| `FrameworkControls` | framework_controls | `list` | required |  | Contains detailed information about all of the controls of a framework. Each framework must contain at least one control. |
| `FrameworkDescription` | framework_description | `string` | optional, computed, provider-chosen |  | An optional description of the framework with a maximum 1,024 characters. |
| `FrameworkName` | framework_name | `string` | optional, computed, provider-chosen, replaces on change |  | The unique name of a framework. This name is between 1 and 256 characters, starting with a letter, and consisting of letters (a-z, A-Z), numbers (0-9), and underscores (_). |
| `FrameworkStatus` | framework_status | `string` | computed |  | A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. The statuses are: |
| `FrameworkTags` | framework_tags | `map` | tags map |  | Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair. |

Supports update: yes

Discovery: supported
