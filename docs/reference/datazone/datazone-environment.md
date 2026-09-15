# aws.datazone.environment

**CloudFormation type:** `AWS::DataZone::Environment`

Definition of AWS::DataZone::Environment Resource Type

Region attribute: `region`

**Import ID:** `<region>/DomainId|Id` (AWS::DataZone::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AwsAccountId` | aws_account_id | `string` | computed |  | The AWS account in which the Amazon DataZone environment is created. |
| `AwsAccountRegion` | aws_account_region | `string` | computed |  | The AWS region in which the Amazon DataZone environment is created. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the environment was created. |
| `CreatedBy` | created_by | `string` | computed |  | The Amazon DataZone user who created the environment. |
| `DeploymentOrder` | deployment_order | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | The deployment order for the environment. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the Amazon DataZone environment. |
| `DomainId` | domain_id | `string` | computed |  | The identifier of the Amazon DataZone domain in which the environment is created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  | The identifier of the Amazon DataZone domain in which the environment would be created. |
| `EnvironmentAccountIdentifier` | environment_account_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The AWS account in which the Amazon DataZone environment is created. |
| `EnvironmentAccountRegion` | environment_account_region | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The AWS region in which the Amazon DataZone environment is created. |
| `EnvironmentBlueprintId` | environment_blueprint_id | `string` | computed |  | The ID of the blueprint with which the Amazon DataZone environment was created. |
| `EnvironmentBlueprintIdentifier` | environment_blueprint_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier of the environment blueprint. |
| `EnvironmentConfigurationId` | environment_configuration_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The identifier of the environment configuration. |
| `EnvironmentProfileId` | environment_profile_id | `string` | computed |  | The ID of the environment profile with which the Amazon DataZone environment was created. |
| `EnvironmentProfileIdentifier` | environment_profile_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of the environment profile with which the Amazon DataZone environment would be created. |
| `EnvironmentRoleArn` | environment_role_arn | `string` | optional, computed, provider-chosen, write-only | aws.role.Arn | Environment role arn for custom aws environment permissions |
| `GlossaryTerms` | glossary_terms | `list` | optional, computed, provider-chosen |  | The glossary terms that can be used in the Amazon DataZone environment. |
| `Id` |  | `string` | computed |  | The ID of the Amazon DataZone environment. |
| `Name` |  | `string` | required |  | The name of the environment. |
| `ProjectId` | project_id | `string` | computed |  | The ID of the Amazon DataZone project in which the environment is created. |
| `ProjectIdentifier` | project_identifier | `string` | required, replaces on change, write-only |  | The ID of the Amazon DataZone project in which the environment would be created. |
| `Provider` | provider_value | `string` | computed |  | The provider of the Amazon DataZone environment. |
| `Status` |  | `string` | computed |  | The status of the Amazon DataZone environment. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp of when the environment was updated. |
| `UserParameters` | user_parameters | `list` | optional, computed, provider-chosen, replaces on change |  | The user parameters of the Amazon DataZone environment. |

Supports update: yes

Discovery: supported (parent resource required)
