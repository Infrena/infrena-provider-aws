# aws.refactorspaces.route

**CloudFormation type:** `AWS::RefactorSpaces::Route`

Definition of AWS::RefactorSpaces::Route Resource Type

Region attribute: `region`

**Import ID:** `<region>/EnvironmentIdentifier|ApplicationIdentifier|RouteIdentifier` (AWS::RefactorSpaces::Route)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationIdentifier` | application_identifier | `string` | required, replaces on change |  |  |
| `Arn` |  | `string` | computed |  |  |
| `DefaultRoute` | default_route | `map` | optional, computed, provider-chosen, write-only |  |  |
| `EnvironmentIdentifier` | environment_identifier | `string` | required, replaces on change |  |  |
| `PathResourceToId` | path_resource_to_id | `string` | computed |  |  |
| `RouteIdentifier` | route_identifier | `string` | computed |  |  |
| `RouteType` | route_type | `string` | required, replaces on change, write-only |  |  |
| `ServiceIdentifier` | service_identifier | `string` | required, replaces on change, write-only |  |  |
| `Tags` |  | `list` | optional, computed, provider-chosen |  | Metadata that you can assign to help organize the frameworks that you create. Each tag is a key-value pair. |
| `UriPathRoute` | uri_path_route | `map` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported (parent resource required)
