# aws.locationazureblob

**CloudFormation type:** `AWS::DataSync::LocationAzureBlob`

Resource Type definition for AWS::DataSync::LocationAzureBlob.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationAzureBlob)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentArns` | agent_arns | `list` | optional, computed, provider-chosen | aws.datasync.agent.AgentArn | Specifies the Amazon Resource Name (ARN) of the DataSync agent that can connect with your Azure Blob Storage container. If you are setting up an agentless cross-cloud transfer, you do not need to specify a value for this parameter. |
| `AzureAccessTier` | azure_access_tier | `string` | optional, computed, provider-chosen |  | Specifies an access tier for the objects you're transferring into your Azure Blob Storage container. |
| `AzureBlobAuthenticationType` | azure_blob_authentication_type | `string` | required |  | The specific authentication type that you want DataSync to use to access your Azure Blob Container. |
| `AzureBlobContainerUrl` | azure_blob_container_url | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The URL of the Azure Blob container that was described. |
| `AzureBlobSasConfiguration` | azure_blob_sas_configuration | `map` | optional, computed, provider-chosen, write-only |  | Specifies the shared access signature (SAS) that DataSync uses to access your Azure Blob Storage container. |
| `AzureBlobType` | azure_blob_type | `string` | optional, computed, provider-chosen |  | Specifies a blob type for the objects you're transferring into your Azure Blob Storage container. |
| `CmkSecretConfig` | cmk_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key. |
| `CustomSecretConfig` | custom_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the Azure Blob Location that is created. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the Azure Blob Location that was described. |
| `ManagedSecretConfig` | managed_secret_config | `map` | computed |  | Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | The subdirectory in the Azure Blob Container that is used to read data from the Azure Blob Source Location. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
