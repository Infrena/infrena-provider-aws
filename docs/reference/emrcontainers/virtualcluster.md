# aws.virtualcluster

**CloudFormation type:** `AWS::EMRContainers::VirtualCluster`

Resource Schema of AWS::EMRContainers::VirtualCluster Type

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EMRContainers::VirtualCluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ContainerProvider` | container_provider | `map` | required, replaces on change |  | Container provider of the virtual cluster. |
| `Id` |  | `string` | computed |  | Id of the virtual cluster. |
| `Name` |  | `string` | required, replaces on change |  | Name of the virtual cluster. |
| `SecurityConfigurationId` | security_configuration_id | `string` | optional, computed, provider-chosen, replaces on change | aws.emrcontainers.securityconfiguration.Id | The ID of the security configuration. |
| `SessionEnabled` | session_enabled | `boolean` | optional, computed, provider-chosen, replaces on change |  | Whether the virtual cluster is session-enabled for Spark Connect. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this virtual cluster. |

Supports update: yes

Discovery: supported
