# aws.browsercustom

**CloudFormation type:** `AWS::BedrockAgentCore::BrowserCustom`

Resource definition for AWS::BedrockAgentCore::BrowserCustom

Region attribute: `region`

**Import ID:** `<region>/BrowserId` (AWS::BedrockAgentCore::BrowserCustom)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `BrowserArn` | browser_arn | `string` | computed |  | The ARN of a Browser resource. |
| `BrowserId` | browser_id | `string` | computed |  | The id of the browser. |
| `BrowserSigning` | browser_signing | `map` | optional, computed, provider-chosen, replaces on change |  | Browser signing configuration |
| `Certificates` |  | `list` | optional, computed, provider-chosen, replaces on change |  | List of root CA certificates. |
| `CreatedAt` | created_at | `string` | computed |  | Timestamp when the browser was created. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The description of the browser. |
| `EnterprisePolicies` | enterprise_policies | `list` | optional, computed, provider-chosen, replaces on change |  | List of browser enterprise policies. |
| `ExecutionRoleArn` | execution_role_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role. |
| `FailureReason` | failure_reason | `string` | computed |  | The reason for failure if the browser creation or operation failed. |
| `LastUpdatedAt` | last_updated_at | `string` | computed |  | Timestamp when the browser was last updated. |
| `Name` |  | `string` | required, replaces on change |  | The name of the browser. |
| `NetworkConfiguration` | network_configuration | `map` | required, replaces on change |  | Network configuration for browser |
| `RecordingConfig` | recording_config | `map` | optional, computed, provider-chosen, replaces on change |  | Recording configuration for browser |
| `Status` |  | `string` | computed |  | Status of browser |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A map of tag keys and values |

Supports update: yes

Discovery: supported
