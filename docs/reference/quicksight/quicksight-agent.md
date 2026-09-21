# aws.quicksight.agent

**CloudFormation type:** `AWS::QuickSight::Agent`

Resource Type definition for AWS::QuickSight::Agent

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|AgentId` (AWS::QuickSight::Agent)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionConnectors` | action_connectors | `list` | optional, computed, provider-chosen |  | A list of ActionConnector ARNs (max 10) attached to the agent. |
| `AgentId` | agent_id | `string` | required, replaces on change | aws.quicksight.agent.AgentId | The unique identifier for the agent. |
| `AgentLifecycle` | agent_lifecycle | `string` | optional, computed, provider-chosen, replaces on change |  | The lifecycle stage of the agent. PREVIEW or PUBLISHED. |
| `AgentStatus` | agent_status | `string` | computed |  | The current status of the agent. One of ACTIVE, CREATING, UPDATING, or FAILED. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the agent. |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  | The ID of the Amazon Web Services account where the agent is being created. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time the agent was created. |
| `Creator` |  | `string` | computed |  | The ARN of the user who created the agent. |
| `CustomPromptInput` | custom_prompt_input | `map` | optional, computed, provider-chosen, write-only |  | Custom prompt configuration. Specify either ExistingPrompt or NewPrompt. |
| `CustomPromptInterface` | custom_prompt_interface | `map` | computed |  | Read-only view of the resolved custom prompt interface for the agent. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the agent. |
| `ErrorMessage` | error_message | `string` | computed |  | The error message if the agent is in FAILED status. |
| `IconId` | icon_id | `string` | optional, computed, provider-chosen |  | The icon identifier for the agent. |
| `Name` |  | `string` | required |  | The display name of the agent. |
| `Spaces` |  | `list` | optional, computed, provider-chosen |  | A list of Space ARNs (max 10) attached to the agent. |
| `StarterPrompts` | starter_prompts | `list` | optional, computed, provider-chosen |  | A list of up to 3 starter prompts displayed to users. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs to associate with the agent resource. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time the agent was last updated. |
| `WelcomeMessage` | welcome_message | `string` | optional, computed, provider-chosen |  | The welcome message displayed when a user opens the agent. |

Supports update: yes

Discovery: supported (parent resource required)
