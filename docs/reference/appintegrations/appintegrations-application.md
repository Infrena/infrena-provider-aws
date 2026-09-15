# aws.appintegrations.application

**CloudFormation type:** `AWS::AppIntegrations::Application`

Resource Type definition for AWS:AppIntegrations::Application

Region attribute: `region`

**Import ID:** `<region>/ApplicationArn` (AWS::AppIntegrations::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationArn` | application_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the application. |
| `ApplicationConfig` | application_config | `map` | optional, computed, provider-chosen |  | The application configuration. Cannot be used when IsService is true. |
| `ApplicationSourceConfig` | application_source_config | `map` | required |  | Application source config |
| `ApplicationType` | application_type | `string` | optional, computed, provider-chosen |  | The type of application |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The application description. |
| `Id` |  | `string` | computed |  | The id of the application. |
| `IframeConfig` | iframe_config | `map` | optional, computed, provider-chosen |  | The iframe configuration |
| `InitializationTimeout` | initialization_timeout | `integer` | optional, computed, provider-chosen |  | The initialization timeout in milliseconds. Required when IsService is true. |
| `IsService` | is_service | `boolean` | optional, computed, provider-chosen |  | Indicates if the application is a service |
| `Name` |  | `string` | required |  | The name of the application. |
| `Namespace` |  | `string` | required |  | The namespace of the application. |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  | The configuration of events or requests that the application has access to. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the application. |

Supports update: yes

Discovery: supported
