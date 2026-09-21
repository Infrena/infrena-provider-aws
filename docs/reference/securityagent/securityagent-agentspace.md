# aws.securityagent.agentspace

**CloudFormation type:** `AWS::SecurityAgent::AgentSpace`

Resource Type definition for AWS::SecurityAgent::AgentSpace

Region attribute: `region`

**Import ID:** `<region>/AgentSpaceId` (AWS::SecurityAgent::AgentSpace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentSpaceId` | agent_space_id | `string` | computed |  | Unique identifier of the agent space |
| `AwsResources` | aws_resources | `map` | optional, computed, provider-chosen |  | AWS resource configuration |
| `CodeReviewSettings` | code_review_settings | `map` | optional, computed, provider-chosen |  | Details of code review settings |
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the agent space was created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the agent space |
| `IntegratedResources` | integrated_resources | `list` | optional, computed, provider-chosen |  | Integrated Resources configuration |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  | Identifier of the KMS key used to encrypt data. Can be a key ID, key ARN, alias name, or alias ARN. If not specified, an AWS managed key is used. |
| `Name` |  | `string` | required |  | Name of the agent space |
| `Tags` |  | `map` | tags map |  | Tags for the agent space |
| `TargetDomainIds` | target_domain_ids | `list` | optional, computed, provider-chosen | aws.targetdomain.TargetDomainId | List of target domain identifiers registered with the agent space |
| `UpdatedAt` | updated_at | `string` | computed |  | Timestamp when the agent space was last updated |

Supports update: yes

Discovery: supported
