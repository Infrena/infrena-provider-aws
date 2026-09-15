# aws.locationobjectstorage

**CloudFormation type:** `AWS::DataSync::LocationObjectStorage`

Resource Type definition for AWS::DataSync::LocationObjectStorage.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationObjectStorage)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessKey` | access_key | `string` | optional, computed, provider-chosen |  | Optional. The access key is used if credentials are required to access the self-managed object storage server. |
| `AgentArns` | agent_arns | `list` | optional, computed, provider-chosen | aws.datasync.agent.AgentArn | Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can connect with your object storage system. If you are setting up an agentless cross-cloud transfer, you do not need to specify a value for this parameter. |
| `BucketName` | bucket_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The name of the bucket on the self-managed object storage server. |
| `CmkSecretConfig` | cmk_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key. |
| `CustomSecretConfig` | custom_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the location that is created. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the object storage location that was described. |
| `ManagedSecretConfig` | managed_secret_config | `map` | computed |  | Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager. |
| `SecretKey` | secret_key | `string` | optional, computed, provider-chosen, sensitive, write-only |  | Optional. The secret key is used if credentials are required to access the self-managed object storage server. |
| `ServerCertificate` | server_certificate | `string` | optional, computed, provider-chosen |  | X.509 PEM content containing a certificate authority or chain to trust. |
| `ServerHostname` | server_hostname | `string` | optional, computed, provider-chosen, write-only |  | The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. |
| `ServerPort` | server_port | `integer` | optional, computed, provider-chosen |  | The port that your self-managed server accepts inbound network traffic on. |
| `ServerProtocol` | server_protocol | `string` | optional, computed, provider-chosen |  | The protocol that the object storage server uses to communicate. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | The subdirectory in the self-managed object storage server that is used to read data from. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
