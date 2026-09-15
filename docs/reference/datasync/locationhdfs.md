# aws.locationhdfs

**CloudFormation type:** `AWS::DataSync::LocationHDFS`

Resource Type definition for AWS::DataSync::LocationHDFS.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationHDFS)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentArns` | agent_arns | `list` | required | aws.datasync.agent.AgentArn | ARN(s) of the agent(s) to use for an HDFS location. |
| `AuthenticationType` | authentication_type | `string` | required |  | The authentication mode used to determine identity of user. |
| `BlockSize` | block_size | `integer` | optional, computed, provider-chosen |  | Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster. |
| `CmkSecretConfig` | cmk_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key. |
| `CustomSecretConfig` | custom_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a customer-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret. |
| `KerberosKeytab` | kerberos_keytab | `string` | optional, computed, provider-chosen, write-only |  | The Base64 string representation of the Keytab file. |
| `KerberosKrb5Conf` | kerberos_krb5_conf | `string` | optional, computed, provider-chosen, write-only |  | The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket. |
| `KerberosPrincipal` | kerberos_principal | `string` | optional, computed, provider-chosen |  | The unique identity, or principal, to which Kerberos can assign tickets. |
| `KmsKeyProviderUri` | kms_key_provider_uri | `string` | optional, computed, provider-chosen |  | The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the HDFS location. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the HDFS location that was described. |
| `ManagedSecretConfig` | managed_secret_config | `map` | computed |  | Specifies configuration information for a DataSync-managed secret, such as an authentication token or set of credentials that DataSync uses to access a specific transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager. |
| `NameNodes` | name_nodes | `list` | required |  | An array of Name Node(s) of the HDFS location. |
| `QopConfiguration` | qop_configuration | `map` | optional, computed, provider-chosen |  | Configuration information for RPC Protection and Data Transfer Protection. These parameters can be set to AUTHENTICATION, INTEGRITY, or PRIVACY. The default value is PRIVACY. |
| `ReplicationFactor` | replication_factor | `integer` | optional, computed, provider-chosen |  | Number of copies of each block that exists inside the HDFS cluster. |
| `SimpleUser` | simple_user | `string` | optional, computed, provider-chosen |  | The user name that has read and write permissions on the specified HDFS cluster. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
