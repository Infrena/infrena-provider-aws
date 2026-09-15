# aws.opensearchservice.application

**CloudFormation type:** `AWS::OpenSearchService::Application`

Amazon OpenSearchService application resource

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::OpenSearchService::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppConfigs` | app_configs | `list` | optional, computed, provider-chosen |  | List of application configurations. |
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN) format. |
| `DataSources` | data_sources | `list` | optional, computed, provider-chosen |  | List of data sources. |
| `Endpoint` |  | `string` | optional, computed, provider-chosen |  | The endpoint for the application. |
| `IamIdentityCenterOptions` | iam_identity_center_options | `map` | optional, computed, provider-chosen |  | Options for configuring IAM Identity Center |
| `Id` |  | `string` | computed |  | The identifier of the application. |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ARN of the KMS key used to encrypt the application. |
| `Name` |  | `string` | required, replaces on change |  | The name of the application. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An arbitrary set of tags (key-value pairs) for this application. |

Supports update: yes

Discovery: supported
