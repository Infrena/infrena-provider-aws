# aws.dataautomationproject

**CloudFormation type:** `AWS::Bedrock::DataAutomationProject`

Definition of AWS::Bedrock::DataAutomationProject Resource Type

Region attribute: `region`

**Import ID:** `<region>/ProjectArn` (AWS::Bedrock::DataAutomationProject)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | Time Stamp |
| `CustomOutputConfiguration` | custom_output_configuration | `map` | optional, computed, provider-chosen |  | Custom output configuration |
| `KmsEncryptionContext` | kms_encryption_context | `map` | optional, computed, provider-chosen |  | KMS encryption context |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen |  | KMS key identifier |
| `LastModifiedTime` | last_modified_time | `string` | computed |  | Time Stamp |
| `OverrideConfiguration` | override_configuration | `map` | optional, computed, provider-chosen |  | Override configuration |
| `ProjectArn` | project_arn | `string` | computed |  | ARN of a DataAutomationProject |
| `ProjectDescription` | project_description | `string` | optional, computed, provider-chosen |  | Description of the DataAutomationProject |
| `ProjectName` | project_name | `string` | required, replaces on change |  | Name of the DataAutomationProject |
| `ProjectStage` | project_stage | `string` | computed |  | Stage of the Project |
| `ProjectType` | project_type | `string` | optional, computed, provider-chosen, replaces on change |  | Type of the DataAutomationProject - Sync or Async |
| `StandardOutputConfiguration` | standard_output_configuration | `map` | optional, computed, provider-chosen |  | Standard output configuration |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | List of Tags |

Supports update: yes

Discovery: supported
