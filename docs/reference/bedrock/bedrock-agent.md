# aws.bedrock.agent

**CloudFormation type:** `AWS::Bedrock::Agent`

Definition of AWS::Bedrock::Agent Resource Type

Region attribute: `region`

**Import ID:** `<region>/AgentId` (AWS::Bedrock::Agent)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActionGroups` | action_groups | `list` | optional, computed, provider-chosen |  | List of ActionGroups |
| `AgentArn` | agent_arn | `string` | computed |  | Arn representation of the Agent. |
| `AgentCollaboration` | agent_collaboration | `string` | optional, computed, provider-chosen |  | Agent collaboration state |
| `AgentCollaborators` | agent_collaborators | `list` | optional, computed, provider-chosen |  | List of Agent Collaborators |
| `AgentId` | agent_id | `string` | computed |  | Identifier for a resource. |
| `AgentName` | agent_name | `string` | required |  | Name for a resource. |
| `AgentResourceRoleArn` | agent_resource_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | ARN of a IAM role. |
| `AgentStatus` | agent_status | `string` | computed |  | Schema Type for Action APIs. |
| `AgentVersion` | agent_version | `string` | computed |  | Draft Agent Version. |
| `AutoPrepare` | auto_prepare | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether to automatically prepare after creating or updating the agent. |
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp. |
| `CustomOrchestration` | custom_orchestration | `map` | optional, computed, provider-chosen |  | Structure for custom orchestration |
| `CustomerEncryptionKeyArn` | customer_encryption_key_arn | `string` | optional, computed, provider-chosen |  | A KMS key ARN |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the Resource. |
| `FailureReasons` | failure_reasons | `list` | computed |  | Failure Reasons for Error. |
| `FoundationModel` | foundation_model | `string` | optional, computed, provider-chosen |  | ARN or name of a Bedrock model. |
| `GuardrailConfiguration` | guardrail_configuration | `map` | optional, computed, provider-chosen |  | Configuration for a guardrail. |
| `IdleSessionTTLInSeconds` | idle_session_ttl_in_seconds | `float` | optional, computed, provider-chosen |  | Max Session Time. |
| `Instruction` |  | `string` | optional, computed, provider-chosen |  | Instruction for the agent. |
| `KnowledgeBases` | knowledge_bases | `list` | optional, computed, provider-chosen |  | List of Agent Knowledge Bases |
| `MemoryConfiguration` | memory_configuration | `map` | optional, computed, provider-chosen |  | Configuration for memory storage |
| `OrchestrationType` | orchestration_type | `string` | optional, computed, provider-chosen |  | Types of orchestration strategy for agents |
| `PreparedAt` | prepared_at | `string` | computed |  | Time Stamp. |
| `PromptOverrideConfiguration` | prompt_override_configuration | `map` | optional, computed, provider-chosen |  | Configuration for prompt override. |
| `RecommendedActions` | recommended_actions | `list` | computed |  | The recommended actions users can take to resolve an error in failureReasons. |
| `SkipResourceInUseCheckOnDelete` | skip_resource_in_use_check_on_delete | `boolean` | optional, computed, provider-chosen, write-only |  | Specifies whether to allow deleting agent while it is in use. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `TestAliasTags` | test_alias_tags | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp. |

Supports update: yes

Discovery: supported
