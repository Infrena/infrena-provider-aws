# aws.server

**CloudFormation type:** `AWS::Transfer::Server`

Definition of AWS::Transfer::Server Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Transfer::Server)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `As2ServiceManagedEgressIpAddresses` | as2_service_managed_egress_ip_addresses | `list` | computed |  | The list of egress IP addresses of this server. These IP addresses are only relevant for servers that use the AS2 protocol. They are used for sending asynchronous MDNs. These IP addresses are assigned automatically when you create an AS2 server. Additionally, if you update an existing server and add the AS2 protocol, static IP addresses are assigned as well. |
| `Certificate` |  | `string` | optional, computed, provider-chosen |  |  |
| `Domain` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `EndpointDetails` | endpoint_details | `map` | optional, computed, provider-chosen |  |  |
| `EndpointType` | endpoint_type | `string` | optional, computed, provider-chosen |  |  |
| `IdentityProviderDetails` | identity_provider_details | `map` | optional, computed, provider-chosen |  |  |
| `IdentityProviderType` | identity_provider_type | `string` | optional, computed, provider-chosen |  |  |
| `IpAddressType` | ip_address_type | `string` | optional, computed, provider-chosen |  |  |
| `LoggingRole` | logging_role | `string` | optional, computed, provider-chosen |  |  |
| `PostAuthenticationLoginBanner` | post_authentication_login_banner | `string` | optional, computed, provider-chosen |  |  |
| `PreAuthenticationLoginBanner` | pre_authentication_login_banner | `string` | optional, computed, provider-chosen |  |  |
| `ProtocolDetails` | protocol_details | `map` | optional, computed, provider-chosen |  |  |
| `Protocols` |  | `list` | optional, computed, provider-chosen |  |  |
| `S3StorageOptions` | s3_storage_options | `map` | optional, computed, provider-chosen |  |  |
| `SecurityPolicyName` | security_policy_name | `string` | optional, computed, provider-chosen |  |  |
| `ServerId` | server_id | `string` | computed |  |  |
| `State` |  | `string` | computed |  |  |
| `StructuredLogDestinations` | structured_log_destinations | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `WorkflowDetails` | workflow_details | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
