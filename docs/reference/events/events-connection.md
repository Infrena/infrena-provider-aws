# aws.events.connection

**CloudFormation type:** `AWS::Events::Connection`

Resource Type definition for AWS::Events::Connection.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Events::Connection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The arn of the connection resource. |
| `ArnForPolicy` | arn_for_policy | `string` | computed |  | The arn of the connection resource to be used in IAM policies. |
| `AuthParameters` | auth_parameters | `map` | optional, computed, provider-chosen |  |  |
| `AuthorizationType` | authorization_type | `string` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the connection. |
| `InvocationConnectivityParameters` | invocation_connectivity_parameters | `map` | optional, computed, provider-chosen |  | The private resource the HTTP request will be sent to. |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Name of the connection. |
| `SecretArn` | secret_arn | `string` | computed |  | The arn of the secrets manager secret created in the customer account. |

Supports update: yes

Discovery: supported
