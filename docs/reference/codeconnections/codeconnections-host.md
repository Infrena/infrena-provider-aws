# aws.codeconnections.host

**CloudFormation type:** `AWS::CodeConnections::Host`

A resource that represents the infrastructure where a third-party provider is installed. You create one host for all connections to that provider.

Region attribute: `region`

**Import ID:** `<region>/HostArn` (AWS::CodeConnections::Host)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `HostArn` | host_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the host. |
| `HostId` | host_id | `string` | computed |  | The server-generated unique identifier for the host. |
| `Name` |  | `string` | required, replaces on change |  | The name of the host. |
| `ProviderEndpoint` | provider_endpoint | `string` | required |  | The endpoint of the infrastructure where your provider type is installed. |
| `ProviderType` | provider_type | `string` | required, replaces on change |  | The name of the installed provider to be associated with your connection. |
| `Status` |  | `string` | computed |  | The status of the host. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to apply to the host. |
| `VpcConfiguration` | vpc_configuration | `map` | optional, computed, provider-chosen |  | The VPC configuration provisioned for the host. |

Supports update: yes

Discovery: supported
