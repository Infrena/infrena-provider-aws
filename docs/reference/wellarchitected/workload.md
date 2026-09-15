# aws.workload

**CloudFormation type:** `AWS::WellArchitected::Workload`

Definition of AWS::WellArchitected::Workload Resource Type

Region attribute: `region`

**Import ID:** `<region>/WorkloadArn` (AWS::WellArchitected::Workload)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountIds` | account_ids | `list` | optional, computed, provider-chosen |  | The list of Amazon Web Services account IDs associated with the workload. |
| `ArchitecturalDesign` | architectural_design | `string` | optional, computed, provider-chosen |  | The URL of the architectural design for the workload. |
| `AwsRegions` | aws_regions | `list` | optional, computed, provider-chosen |  | The list of Amazon Web Services Regions associated with the workload. |
| `Description` |  | `string` | required |  | The description for the workload. |
| `DiscoveryConfig` | discovery_config | `map` | optional, computed, provider-chosen |  | Discovery configuration associated to the workload. |
| `Environment` |  | `string` | required |  | The environment for the workload. |
| `ImprovementStatus` | improvement_status | `string` | computed |  | The improvement status for a workload. |
| `Industry` |  | `string` | optional, computed, provider-chosen |  | The industry for the workload. |
| `IndustryType` | industry_type | `string` | optional, computed, provider-chosen |  | The industry type for the workload. |
| `Lenses` |  | `list` | required, replaces on change |  | The list of lenses associated with the workload. |
| `NonAwsRegions` | non_aws_regions | `list` | optional, computed, provider-chosen |  | The list of non-Amazon Web Services Regions associated with the workload. |
| `Notes` |  | `string` | optional, computed, provider-chosen |  | The notes associated with the workload. |
| `ReviewOwner` | review_owner | `string` | optional, computed, provider-chosen |  | The review owner of the workload. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with the workload. |
| `WorkloadArn` | workload_arn | `string` | computed |  | The ARN for the workload. |
| `WorkloadId` | workload_id | `string` | computed |  | The ID assigned to the workload. |
| `WorkloadName` | workload_name | `string` | required |  | The name of the workload. |

Supports update: yes

Discovery: supported
