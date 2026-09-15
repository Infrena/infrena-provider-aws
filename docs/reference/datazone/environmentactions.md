# aws.environmentactions

**CloudFormation type:** `AWS::DataZone::EnvironmentActions`

Definition of AWS::DataZone::EnvironmentActions Resource Type

Region attribute: `region`

**Import ID:** `<region>/DomainId|EnvironmentId|Id` (AWS::DataZone::EnvironmentActions)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the Amazon DataZone environment action. |
| `DomainId` | domain_id | `string` | computed |  | The identifier of the Amazon DataZone domain in which the environment is created. |
| `DomainIdentifier` | domain_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier of the Amazon DataZone domain in which the environment would be created. |
| `EnvironmentId` | environment_id | `string` | computed |  | The identifier of the Amazon DataZone environment in which the action is taking place |
| `EnvironmentIdentifier` | environment_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier of the Amazon DataZone environment in which the action is taking place |
| `Id` |  | `string` | computed |  | The ID of the Amazon DataZone environment action. |
| `Identifier` |  | `string` | optional, computed, provider-chosen, write-only |  | The ID of the Amazon DataZone environment action. |
| `Name` |  | `string` | required |  | The name of the environment action. |
| `Parameters` |  | `map` | optional, computed, provider-chosen |  | The parameters of the console link specified as part of the environment action |

Supports update: yes

Discovery: supported (parent resource required)
