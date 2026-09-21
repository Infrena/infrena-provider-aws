# aws.servicediscovery.service

**CloudFormation type:** `AWS::ServiceDiscovery::Service`

Resource Type definition for AWS::ServiceDiscovery::Service

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ServiceDiscovery::Service)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the service. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the service. |
| `DnsConfig` | dns_config | `map` | optional, computed, provider-chosen |  | DNS configuration settings for the service. |
| `HealthCheckConfig` | health_check_config | `map` | optional, computed, provider-chosen |  | Configuration for health checks for the service. |
| `HealthCheckCustomConfig` | health_check_custom_config | `map` | optional, computed, provider-chosen, replaces on change |  | Configurations for custom health checks for the service. |
| `Id` |  | `string` | computed |  | The unique identifier for the service. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the service. |
| `NamespaceId` | namespace_id | `string` | optional, computed, provider-chosen, replaces on change |  | The ID of the namespace in which the service is created. |
| `ServiceAttributes` | service_attributes | `map` | optional, computed, provider-chosen |  | A string map that contains attributes and values for the service. You can specify a maximum of 30 key-value pairs. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to associate with the service. |
| `Type` | type_value | `string` | optional, computed, provider-chosen, replaces on change |  | The type of service. Supported values are HTTP or DNS. |

Supports update: yes

Discovery: supported
