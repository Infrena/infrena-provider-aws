# aws.segmentdefinition

**CloudFormation type:** `AWS::CustomerProfiles::SegmentDefinition`

A segment definition resource of Amazon Connect Customer Profiles

Region attribute: `region`

**Import ID:** `<region>/DomainName|SegmentDefinitionName` (AWS::CustomerProfiles::SegmentDefinition)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The time of this segment definition got created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the segment definition. |
| `DisplayName` | display_name | `string` | required, replaces on change |  | The display name of the segment definition. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The unique name of the domain. |
| `SegmentDefinitionArn` | segment_definition_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the segment definition. |
| `SegmentDefinitionName` | segment_definition_name | `string` | required, replaces on change |  | The unique name of the segment definition. |
| `SegmentGroups` | segment_groups | `map` | optional, computed, provider-chosen, replaces on change |  | An array that defines the set of segment criteria to evaluate when handling segment groups for the segment. |
| `SegmentSort` | segment_sort | `map` | optional, computed, provider-chosen |  | Defines how segments should be sorted and ordered in the results. |
| `SegmentSqlQuery` | segment_sql_query | `string` | optional, computed, provider-chosen, replaces on change |  | The SQL query that defines the segment criteria. |
| `SegmentType` | segment_type | `string` | computed |  | The SQL query that defines the segment criteria. |
| `Tags` |  | `map` | tags map |  | The tags used to organize, track, or control access for this resource. |

Supports update: yes

Discovery: supported (parent resource required)
