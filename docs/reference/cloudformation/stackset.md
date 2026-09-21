# aws.stackset

**CloudFormation type:** `AWS::CloudFormation::StackSet`

StackSet as a resource provides one-click experience for provisioning a StackSet and StackInstances

Region attribute: `region`

**Import ID:** `<region>/StackSetId` (AWS::CloudFormation::StackSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdministrationRoleARN` | administration_role_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Number (ARN) of the IAM role to use to create this stack set. Specify an IAM role only if you are using customized administrator roles to control which users or groups can manage specific stack sets within the same administrator account. |
| `AutoDeployment` | auto_deployment | `map` | optional, computed, provider-chosen |  | Describes whether StackSets automatically deploys to AWS Organizations accounts that are added to the target organization or organizational unit (OU). Specify only if PermissionModel is SERVICE_MANAGED. |
| `CallAs` | call_as | `string` | optional, computed, provider-chosen, write-only |  | Specifies the AWS account that you are acting from. By default, SELF is specified. For self-managed permissions, specify SELF; for service-managed permissions, if you are signed in to the organization's management account, specify SELF. If you are signed in to a delegated administrator account, specify DELEGATED_ADMIN. |
| `Capabilities` |  | `list` | optional, computed, provider-chosen |  | In some cases, you must explicitly acknowledge that your stack set template contains certain capabilities in order for AWS CloudFormation to create the stack set and related stack instances. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the stack set. You can use the description to identify the stack set's purpose or other important information. |
| `ExecutionRoleName` | execution_role_name | `string` | optional, computed, provider-chosen |  | The name of the IAM execution role to use to create the stack set. If you do not specify an execution role, AWS CloudFormation uses the AWSCloudFormationStackSetExecutionRole role for the stack set operation. |
| `ManagedExecution` | managed_execution | `map` | optional, computed, provider-chosen |  | Describes whether StackSets performs non-conflicting operations concurrently and queues conflicting operations. |
| `OperationPreferences` | operation_preferences | `map` | optional, computed, provider-chosen, write-only |  | The user-specified preferences for how AWS CloudFormation performs a stack set operation. |
| `Parameters` |  | `list` | optional, computed, provider-chosen |  | The input parameters for the stack set template. |
| `PermissionModel` | permission_model | `string` | required, replaces on change |  | Describes how the IAM roles required for stack set operations are created. By default, SELF-MANAGED is specified. |
| `StackInstancesGroup` | stack_instances_group | `list` | optional, computed, provider-chosen, write-only |  | A group of stack instances with parameters in some specific accounts and regions. |
| `StackSetId` | stack_set_id | `string` | computed |  | The ID of the stack set that you're creating. |
| `StackSetName` | stack_set_name | `string` | required, replaces on change |  | The name to associate with the stack set. The name must be unique in the Region where you create your stack set. |
| `Tags` |  | `map` | tags map |  | The key-value pairs to associate with this stack set and the stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the stacks. A maximum number of 50 tags can be specified. |
| `TemplateBody` | template_body | `string` | optional, computed, provider-chosen |  | The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes. |
| `TemplateURL` | template_url | `string` | optional, computed, provider-chosen, write-only |  | Location of file containing the template body. The URL must point to a template (max size: 460,800 bytes) that is located in an Amazon S3 bucket. |

Supports update: yes

Discovery: supported
