# aws.container

**CloudFormation type:** `AWS::Lightsail::Container`

Resource Type definition for AWS::Lightsail::Container

Region attribute: `region`

**Import ID:** `<region>/ServiceName` (AWS::Lightsail::Container)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContainerArn` | container_arn | `string` | computed |  |  |
| `ContainerServiceDeployment` | container_service_deployment | `map` | optional, computed, provider-chosen |  | Describes a container deployment configuration of an Amazon Lightsail container service. |
| `IsDisabled` | is_disabled | `boolean` | optional, computed, provider-chosen |  | A Boolean value to indicate whether the container service is disabled. |
| `Power` |  | `string` | required |  | The power specification for the container service. |
| `PrincipalArn` | principal_arn | `string` | computed |  | The principal ARN of the container service. |
| `PrivateRegistryAccess` | private_registry_access | `map` | optional, computed, provider-chosen |  | An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. |
| `PublicDomainNames` | public_domain_names | `list` | optional, computed, provider-chosen |  | The public domain names to use with the container service, such as example.com and www.example.com. |
| `Scale` |  | `integer` | required |  | The scale specification for the container service. |
| `ServiceName` | service_name | `string` | required, replaces on change |  | The name for the container service. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Url` |  | `string` | computed |  | The publicly accessible URL of the container service. |

Supports update: yes

Discovery: supported
