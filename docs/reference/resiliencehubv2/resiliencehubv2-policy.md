# aws.resiliencehubv2.policy

**CloudFormation type:** `AWS::ResilienceHubV2::Policy`

Creates a resilience policy that defines availability and disaster recovery requirements.

Region attribute: `region`

**Import ID:** `<region>/PolicyArn` (AWS::ResilienceHubV2::Policy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociatedServiceCount` | associated_service_count | `integer` | computed |  | The number of services associated with this policy. |
| `AvailabilitySlo` | availability_slo | `map` | optional, computed, provider-chosen |  |  |
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the policy was created. |
| `DataRecovery` | data_recovery | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the policy. |
| `KmsKeyId` | kms_key_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The KMS key ID for encrypting policy data. |
| `MultiAz` | multi_az | `map` | optional, computed, provider-chosen |  |  |
| `MultiRegion` | multi_region | `map` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required, replaces on change |  | The name of the policy. |
| `PolicyArn` | policy_arn | `string` | computed |  | The ARN of the policy. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags assigned to the policy. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the policy was last updated. |

Supports update: yes

Discovery: supported
