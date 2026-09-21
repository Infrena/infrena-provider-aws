# aws.interconnect.connection

**CloudFormation type:** `AWS::Interconnect::Connection`

Resource Type definition for AWS::Interconnect::Connection. Creates a managed network connection between AWS and a partner cloud service provider.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Interconnect::Connection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ActivationKey` | activation_key | `string` | optional, computed, provider-chosen, replaces on change |  | The activation key for accepting a connection proposal from a partner CSP. Mutually exclusive with EnvironmentId. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the connection. |
| `AttachPoint` | attach_point | `map` | required, replaces on change |  | The logical attachment point in your AWS network where the managed connection will be connected. |
| `Bandwidth` |  | `string` | optional, computed, provider-chosen |  | The bandwidth of the connection (e.g., 50Mbps, 1Gbps). Required when creating a connection through AWS. |
| `BillingTier` | billing_tier | `integer` | computed |  | The billing tier for the connection. |
| `ConnectionId` | connection_id | `string` | computed |  | The unique identifier for the connection. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the connection. |
| `EnvironmentId` | environment_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the environment for the connection. Required when creating a connection through AWS. Mutually exclusive with ActivationKey. |
| `OwnerAccount` | owner_account | `string` | computed |  | The AWS account ID of the connection owner. |
| `Provider` | provider_value | `map` | computed |  | The partner cloud service provider. |
| `RemoteAccount` | remote_account | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | The remote account identifier for the connection. Required when creating a connection through AWS. Replaces RemoteOwnerAccount. |
| `RemoteOwnerAccount` | remote_owner_account | `string` | optional, computed, provider-chosen, write-only |  | Deprecated. Use RemoteAccount instead. The account ID of the remote owner. Required when creating a connection through AWS. |
| `SharedId` | shared_id | `string` | computed |  | The shared identifier for the connection pairing. |
| `State` |  | `string` | computed |  | The current state of the connection. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Type` | type_value | `string` | computed |  | The type of managed connection. |

Supports update: yes

Discovery: supported
