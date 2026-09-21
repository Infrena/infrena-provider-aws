# aws.workloadidentity

**CloudFormation type:** `AWS::BedrockAgentCore::WorkloadIdentity`

Definition of AWS::BedrockAgentCore::WorkloadIdentity Resource Type

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::BedrockAgentCore::WorkloadIdentity)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedResourceOauth2ReturnUrls` | allowed_resource_oauth2_return_urls | `list` | optional, computed, provider-chosen |  | The list of allowed OAuth2 return URLs for resources associated with this workload identity. |
| `CreatedTime` | created_time | `float` | computed |  | The timestamp when the workload identity was created. |
| `LastUpdatedTime` | last_updated_time | `float` | computed |  | The timestamp when the workload identity was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the workload identity. The name must be unique within your account. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `WorkloadIdentityArn` | workload_identity_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the workload identity. |

Supports update: yes

Discovery: supported
