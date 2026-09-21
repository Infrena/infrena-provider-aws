# aws.subscriptiontarget

**CloudFormation type:** `AWS::DataZone::SubscriptionTarget`

Subscription targets enables one to access the data to which you have subscribed in your projects.

Region attribute: `region`

**Import ID:** `<region>/DomainId|EnvironmentId|Id` (AWS::DataZone::SubscriptionTarget)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicableAssetTypes` | applicable_asset_types | `list` | required |  | The asset types that can be included in the subscription target. |
| `AuthorizedPrincipals` | authorized_principals | `list` | required |  | The authorized principals of the subscription target. |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp of when the subscription target was created. |
| `CreatedBy` | created_by | `string` | computed |  | The Amazon DataZone user who created the subscription target. |
| `DomainId` | domain_id | `string` | computed |  | The ID of the Amazon DataZone domain in which subscription target is created. |
| `DomainIdentifier` | domain_identifier | `string` | required, replaces on change, write-only |  | The ID of the Amazon DataZone domain in which subscription target would be created. |
| `EnvironmentId` | environment_id | `string` | computed |  | The ID of the environment in which subscription target is created. |
| `EnvironmentIdentifier` | environment_identifier | `string` | required, replaces on change, write-only |  | The ID of the environment in which subscription target would be created. |
| `Id` |  | `string` | computed |  | The ID of the subscription target. |
| `ManageAccessRole` | manage_access_role | `string` | optional, computed, provider-chosen |  | The manage access role that is used to create the subscription target. |
| `Name` |  | `string` | required |  | The name of the subscription target. |
| `ProjectId` | project_id | `string` | computed |  | The identifier of the project specified in the subscription target. |
| `Provider` | provider_value | `string` | optional, computed, provider-chosen |  | The provider of the subscription target. |
| `SubscriptionTargetConfig` | subscription_target_config | `list` | required |  | The configuration of the subscription target. |
| `Type` | type_value | `string` | required, replaces on change |  | The type of the subscription target. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp of when the subscription target was updated. |
| `UpdatedBy` | updated_by | `string` | computed |  | The Amazon DataZone user who updated the subscription target. |

Supports update: yes

Discovery: supported (parent resource required)
