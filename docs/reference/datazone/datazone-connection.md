# aws.datazone.connection

**CloudFormation type:** `AWS::DataZone::Connection`

Connections enables users to connect their DataZone resources (domains, projects, and environments) to external resources/services (data, compute, etc)

Region attribute: `region`

**Import ID:** `<region>/DomainId|ConnectionId` (AWS::DataZone::Connection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AwsLocation` | aws_location | `map` | optional, computed, provider-chosen, write-only |  | AWS Location of project |
| `Configurations` |  | `list` | optional, computed, provider-chosen, write-only |  | The configurations of the connection. |
| `ConnectionId` | connection_id | `string` | computed |  | The ID of the connection. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the connection. |
| `DomainId` | domain_id | `string` | computed |  | The ID of the domain in which the connection is created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  | The identifier of the domain in which the connection is created. |
| `DomainUnitId` | domain_unit_id | `string` | computed |  | The ID of the domain unit in which the connection is created. |
| `EnableTrustedIdentityPropagation` | enable_trusted_identity_propagation | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  | Specifies whether the trusted identity propagation is enabled |
| `EnvironmentId` | environment_id | `string` | computed |  | The ID of the environment in which the connection is created. |
| `EnvironmentIdentifier` | environment_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier of the environment in which the connection is created. |
| `EnvironmentUserRole` | environment_user_role | `string` | computed |  | The role of the user in the environment. |
| `Name` |  | `string` | required, replaces on change |  | The name of the connection. |
| `ProjectId` | project_id | `string` | computed |  | The ID of the project in which the connection is created. |
| `ProjectIdentifier` | project_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier of the project in which the connection should be created. If |
| `Props` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Scope` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The scope of the connection. |
| `Type` | type_value | `string` | computed |  | Connection Type |

Supports update: yes

Discovery: supported (parent resource required)
