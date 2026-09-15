# aws.transfer.connector

**CloudFormation type:** `AWS::Transfer::Connector`

Resource Type definition for AWS::Transfer::Connector

Region attribute: `region`

**Import ID:** `<region>/ConnectorId` (AWS::Transfer::Connector)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessRole` | access_role | `string` | required |  | Specifies the access role for the connector. |
| `Arn` |  | `string` | computed |  | Specifies the unique Amazon Resource Name (ARN) for the connector. |
| `As2Config` | as2_config | `map` | optional, computed, provider-chosen |  | Configuration for an AS2 connector. |
| `ConnectorId` | connector_id | `string` | computed |  | A unique identifier for the connector. |
| `EgressConfig` | egress_config | `map` | optional, computed, provider-chosen |  | Egress configuration for the connector. |
| `EgressType` | egress_type | `string` | optional, computed, provider-chosen |  | Specifies the egress type for the connector. |
| `ErrorMessage` | error_message | `string` | computed |  | Detailed error message when Connector in ERRORED status |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  | IP address type for Connector |
| `LoggingRole` | logging_role | `string` | optional, computed, provider-chosen |  | Specifies the logging role for the connector. |
| `SecurityPolicyName` | security_policy_name | `string` | optional, computed, provider-chosen |  | Security policy for SFTP Connector |
| `ServiceManagedEgressIpAddresses` | service_managed_egress_ip_addresses | `list` | computed |  | The list of egress IP addresses of this connector. These IP addresses are assigned automatically when you create the connector. |
| `SftpConfig` | sftp_config | `map` | optional, computed, provider-chosen |  | Configuration for an SFTP connector. |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Key-value pairs that can be used to group and search for connectors. Tags are metadata attached to connectors for any purpose. |
| `Url` |  | `string` | optional, computed, provider-chosen |  | URL for Connector |

Supports update: yes

Discovery: supported
