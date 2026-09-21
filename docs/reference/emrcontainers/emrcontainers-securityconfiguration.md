# aws.emrcontainers.securityconfiguration

**CloudFormation type:** `AWS::EMRContainers::SecurityConfiguration`

Resource Schema of AWS::EMRContainers::SecurityConfiguration Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::EMRContainers::SecurityConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the security configuration. |
| `ContainerProvider` | container_provider | `map` | optional, computed, provider-chosen, replaces on change |  | Container provider information. |
| `Id` |  | `string` | computed |  | The ID of the security configuration. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the security configuration. |
| `SecurityConfigurationData` | security_configuration_data | `map` | required, replaces on change, write-only |  | Security configuration data containing encryption and authorization settings. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this security configuration. |

Supports update: yes

Discovery: supported
