# aws.privateconnection

**CloudFormation type:** `AWS::DevOpsAgent::PrivateConnection`

Resource Type definition for AWS::DevOpsAgent::PrivateConnection

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::DevOpsAgent::PrivateConnection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Private Connection. |
| `Certificate` |  | `string` | optional, computed, provider-chosen, write-only |  | Certificate for the Private Connection. |
| `CertificateExpiryTime` | certificate_expiry_time | `string` | computed |  | The expiry time of the certificate associated with the Private Connection. |
| `ConnectionConfiguration` | connection_configuration | `map` | required, replaces on change |  | The connection configuration, either SelfManaged or ServiceManaged. |
| `Name` |  | `string` | required, replaces on change |  | Unique name for this Private Connection within the account. |
| `Status` |  | `string` | computed |  | The status of the Private Connection. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
