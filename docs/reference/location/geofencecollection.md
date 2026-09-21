# aws.geofencecollection

**CloudFormation type:** `AWS::Location::GeofenceCollection`

Definition of AWS::Location::GeofenceCollection Resource Type

Region attribute: `region`

**Import ID:** `<region>/CollectionName` (AWS::Location::GeofenceCollection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `CollectionArn` | collection_arn | `string` | computed |  |  |
| `CollectionName` | collection_name | `string` | required, replaces on change |  |  |
| `CreateTime` | create_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PricingPlan` | pricing_plan | `string` | optional, computed, provider-chosen |  |  |
| `PricingPlanDataSource` | pricing_plan_data_source | `string` | optional, computed, provider-chosen |  | This shape is deprecated since 2022-02-01: Deprecated. No longer allowed. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `UpdateTime` | update_time | `string` | computed |  | The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ) |

Supports update: yes

Discovery: supported
