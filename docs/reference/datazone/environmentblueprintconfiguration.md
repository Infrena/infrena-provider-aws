# aws.environmentblueprintconfiguration

**CloudFormation type:** `AWS::DataZone::EnvironmentBlueprintConfiguration`

Definition of AWS::DataZone::EnvironmentBlueprintConfiguration Resource Type

Region attribute: `region`

**Import ID:** `<region>/DomainId|EnvironmentBlueprintId` (AWS::DataZone::EnvironmentBlueprintConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  |  |
| `DomainId` | domain_id | `string` | computed |  |  |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  |  |
| `EnabledRegions` | enabled_regions | `list` | required |  |  |
| `EnvironmentBlueprintId` | environment_blueprint_id | `string` | computed |  |  |
| `EnvironmentBlueprintIdentifier` | environment_blueprint_identifier | `string` | required, replaces on change, write-only |  |  |
| `EnvironmentRolePermissionBoundary` | environment_role_permission_boundary | `string` | optional, computed, provider-chosen, write-only |  |  |
| `GlobalParameters` | global_parameters | `map` | optional, computed, provider-chosen, write-only |  | Region-agnostic environment blueprint parameters. |
| `ManageAccessRoleArn` | manage_access_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `ProvisioningConfigurations` | provisioning_configurations | `list` | optional, computed, provider-chosen, write-only |  |  |
| `ProvisioningRoleArn` | provisioning_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `RegionalParameters` | regional_parameters | `list` | optional, computed, provider-chosen |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
