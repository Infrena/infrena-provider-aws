# aws.locationsmb

**CloudFormation type:** `AWS::DataSync::LocationSMB`

Resource Type definition for AWS::DataSync::LocationSMB.

Region attribute: `region`

**Import ID:** `<region>/LocationArn` (AWS::DataSync::LocationSMB)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AgentArns` | agent_arns | `list` | required | aws.datasync.agent.AgentArn | The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location. |
| `AuthenticationType` | authentication_type | `string` | optional, computed, provider-chosen |  | The authentication mode used to determine identity of user. |
| `CmkSecretConfig` | cmk_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a DataSync-managed secret, such as a password or set of credentials that DataSync uses to access a specific transfer location, and a customer-managed AWS KMS key. |
| `CustomSecretConfig` | custom_secret_config | `map` | optional, computed, provider-chosen |  | Specifies configuration information for a customer-managed secret, such as a password or set of credentials that DataSync uses to access a specific transfer location, and an IAM role that DataSync can assume and access the customer-managed secret. |
| `DnsIpAddresses` | dns_ip_addresses | `list` | optional, computed, provider-chosen |  | Specifies the IPv4 addresses for the DNS servers that your SMB file server belongs to. This parameter applies only if AuthenticationType is set to KERBEROS. If you have multiple domains in your environment, configuring this parameter makes sure that DataSync connects to the right SMB file server. |
| `Domain` |  | `string` | optional, computed, provider-chosen |  | The name of the Windows domain that the SMB server belongs to. |
| `KerberosKeytab` | kerberos_keytab | `string` | optional, computed, provider-chosen, write-only |  | The Base64 string representation of the Keytab file. Specifies your Kerberos key table (keytab) file, which includes mappings between your service principal name (SPN) and encryption keys. To avoid task execution errors, make sure that the SPN in the keytab file matches exactly what you specify for KerberosPrincipal and in your krb5.conf file. |
| `KerberosKrb5Conf` | kerberos_krb5_conf | `string` | optional, computed, provider-chosen, write-only |  | The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket. Specifies a Kerberos configuration file (krb5.conf) that defines your Kerberos realm configuration. To avoid task execution errors, make sure that the service principal name (SPN) in the krb5.conf file matches exactly what you specify for KerberosPrincipal and in your keytab file. |
| `KerberosPrincipal` | kerberos_principal | `string` | optional, computed, provider-chosen |  | Specifies a service principal name (SPN), which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server. SPNs are case sensitive and must include a prepended cifs/. For example, an SPN might look like cifs/kerberosuser@EXAMPLE.COM. Your task execution will fail if the SPN that you provide for this parameter doesn't match exactly what's in your keytab or krb5.conf files. |
| `LocationArn` | location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the SMB location that is created. |
| `LocationUri` | location_uri | `string` | computed |  | The URL of the SMB location that was described. |
| `ManagedSecretConfig` | managed_secret_config | `map` | computed |  | Specifies configuration information for a DataSync-managed secret, such as a password or set of credentials that DataSync uses to access a specific transfer location. DataSync uses the default AWS-managed KMS key to encrypt this secret in AWS Secrets Manager. |
| `MountOptions` | mount_options | `map` | optional, computed, provider-chosen |  | The mount options used by DataSync to access the SMB server. |
| `Password` |  | `string` | optional, computed, provider-chosen, sensitive, write-only |  | The password of the user who can mount the share and has the permissions to access files and folders in the SMB share. |
| `ServerHostname` | server_hostname | `string` | optional, computed, provider-chosen, write-only |  | The name of the SMB server. This value is the IP address or Domain Name Service (DNS) name of the SMB server. |
| `Subdirectory` |  | `string` | optional, computed, provider-chosen, write-only |  | The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `User` |  | `string` | optional, computed, provider-chosen |  | The user who can mount the share, has the permissions to access files and folders in the SMB share. |

Supports update: yes

Discovery: supported
