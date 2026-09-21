# aws.intelligentpromptrouter

**CloudFormation type:** `AWS::Bedrock::IntelligentPromptRouter`

Definition of AWS::Bedrock::IntelligentPromptRouter Resource Type

Region attribute: `region`

**Import ID:** `<region>/PromptRouterArn` (AWS::Bedrock::IntelligentPromptRouter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | Time Stamp |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Description of the Prompt Router. |
| `FallbackModel` | fallback_model | `map` | required, replaces on change |  | Model configuration |
| `Models` |  | `list` | required, replaces on change |  | List of model configuration |
| `PromptRouterArn` | prompt_router_arn | `string` | computed |  | Arn of the Prompt Router. |
| `PromptRouterName` | prompt_router_name | `string` | required, replaces on change |  | Name of the Prompt Router. |
| `RoutingCriteria` | routing_criteria | `map` | required, replaces on change |  | Represents the criteria used for routing requests. |
| `Status` |  | `string` | computed |  | Status of a PromptRouter |
| `Tags` |  | `map` | tags map |  | List of Tags |
| `Type` | type_value | `string` | computed |  | Type of a Prompt Router |
| `UpdatedAt` | updated_at | `string` | computed |  | Time Stamp |

Supports update: yes

Discovery: supported
