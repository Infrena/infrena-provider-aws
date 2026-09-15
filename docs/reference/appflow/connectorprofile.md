# aws.connectorprofile

**CloudFormation type:** `AWS::AppFlow::ConnectorProfile`

Resource Type definition for AWS::AppFlow::ConnectorProfile

Region attribute: `region`

**Import ID:** `<region>/ConnectorProfileName` (AWS::AppFlow::ConnectorProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectionMode` | connection_mode | `string` | required |  | Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular |
| `ConnectorLabel` | connector_label | `string` | optional, computed, provider-chosen, replaces on change |  | The label of the connector. The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for CUSTOMCONNECTOR connector type/. |
| `ConnectorProfileArn` | connector_profile_arn | `string` | computed |  | Unique identifier for connector profile resources |
| `ConnectorProfileConfig` | connector_profile_config | `map` | optional, computed, provider-chosen, write-only |  | Connector specific configurations needed to create connector profile |
| `ConnectorProfileName` | connector_profile_name | `string` | required, replaces on change |  | The maximum number of items to retrieve in a single batch. |
| `ConnectorType` | connector_type | `string` | required, replaces on change |  | List of Saas providers that need connector profile to be created |
| `CredentialsArn` | credentials_arn | `string` | computed |  | A unique Arn for Connector-Profile resource |
| `KMSArn` | kms_arn | `string` | optional, computed, provider-chosen, write-only |  | The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key. |

Supports update: yes

Discovery: supported
