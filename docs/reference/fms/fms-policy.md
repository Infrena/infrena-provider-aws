# aws.fms.policy

**CloudFormation type:** `AWS::FMS::Policy`

Creates an AWS Firewall Manager policy.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::FMS::Policy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | A resource ARN. |
| `DeleteAllPolicyResources` | delete_all_policy_resources | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `ExcludeMap` | exclude_map | `map` | optional, computed, provider-chosen |  | An FMS includeMap or excludeMap. |
| `ExcludeResourceTags` | exclude_resource_tags | `boolean` | required |  |  |
| `Id` |  | `string` | computed |  |  |
| `IncludeMap` | include_map | `map` | optional, computed, provider-chosen |  | An FMS includeMap or excludeMap. |
| `PolicyDescription` | policy_description | `string` | optional, computed, provider-chosen |  |  |
| `PolicyName` | policy_name | `string` | required |  |  |
| `RemediationEnabled` | remediation_enabled | `boolean` | required |  |  |
| `ResourceSetIds` | resource_set_ids | `list` | optional, computed, provider-chosen | aws.fms.resourceset.Id |  |
| `ResourceTagLogicalOperator` | resource_tag_logical_operator | `string` | optional, computed, provider-chosen |  |  |
| `ResourceTags` | resource_tags | `list` | optional, computed, provider-chosen |  |  |
| `ResourceType` | resource_type | `string` | optional, computed, provider-chosen |  | An AWS resource type |
| `ResourceTypeList` | resource_type_list | `list` | optional, computed, provider-chosen |  |  |
| `ResourcesCleanUp` | resources_clean_up | `boolean` | optional, computed, provider-chosen |  |  |
| `SecurityServicePolicyData` | security_service_policy_data | `map` | required |  | Firewall security service policy data. |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
