# aws.environmentprofile

**CloudFormation type:** `AWS::DataZone::EnvironmentProfile`

AWS Datazone Environment Profile is pre-configured set of resources and blueprints that provide reusable templates for creating environments.

Region attribute: `region`

**Import ID:** `<region>/DomainId|Id` (AWS::DataZone::EnvironmentProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AwsAccountId` | aws_account_id | `string` | required |  | The AWS account in which the Amazon DataZone environment is created. |
| `AwsAccountRegion` | aws_account_region | `string` | required |  | The AWS region in which this environment profile is created. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when this environment profile was created. |
| `CreatedBy` | created_by | `string` | computed |  | The Amazon DataZone user who created this environment profile. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of this Amazon DataZone environment profile. |
| `DomainId` | domain_id | `string` | computed |  | The ID of the Amazon DataZone domain in which this environment profile is created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  | The ID of the Amazon DataZone domain in which this environment profile is created. |
| `EnvironmentBlueprintId` | environment_blueprint_id | `string` | computed |  | The ID of the blueprint with which this environment profile is created. |
| `EnvironmentBlueprintIdentifier` | environment_blueprint_identifier | `string` | required, replaces on change, write-only |  | The ID of the blueprint with which this environment profile is created. |
| `Id` |  | `string` | computed |  | The ID of this Amazon DataZone environment profile. |
| `Name` |  | `string` | required |  | The name of this Amazon DataZone environment profile. |
| `ProjectId` | project_id | `string` | computed |  | The identifier of the project in which to create the environment profile. |
| `ProjectIdentifier` | project_identifier | `string` | required, replaces on change, write-only |  | The identifier of the project in which to create the environment profile. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp of when this environment profile was updated. |
| `UserParameters` | user_parameters | `list` | optional, computed, provider-chosen |  | The user parameters of this Amazon DataZone environment profile. |

Supports update: yes

Discovery: supported (parent resource required)
