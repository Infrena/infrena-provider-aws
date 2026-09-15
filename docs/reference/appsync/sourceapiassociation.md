# aws.sourceapiassociation

**CloudFormation type:** `AWS::AppSync::SourceApiAssociation`

Resource Type definition for AWS::AppSync::SourceApiAssociation

Region attribute: `region`

**Import ID:** `<region>/AssociationArn` (AWS::AppSync::SourceApiAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationArn` | association_arn | `string` | computed |  | ARN of the SourceApiAssociation. |
| `AssociationId` | association_id | `string` | computed |  | Id of the SourceApiAssociation. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of the SourceApiAssociation. |
| `LastSuccessfulMergeDate` | last_successful_merge_date | `string` | computed |  | Date of last schema successful merge. |
| `MergedApiArn` | merged_api_arn | `string` | computed |  | ARN of the Merged API in the association. |
| `MergedApiId` | merged_api_id | `string` | computed |  | GraphQLApiId of the Merged API in the association. |
| `MergedApiIdentifier` | merged_api_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Identifier of the Merged GraphQLApi to associate. It could be either GraphQLApi ApiId or ARN |
| `SourceApiArn` | source_api_arn | `string` | computed |  | ARN of the source API in the association. |
| `SourceApiAssociationConfig` | source_api_association_config | `map` | optional, computed, provider-chosen |  | Customized configuration for SourceApiAssociation. |
| `SourceApiAssociationStatus` | source_api_association_status | `string` | computed |  | Current status of SourceApiAssociation. |
| `SourceApiAssociationStatusDetail` | source_api_association_status_detail | `string` | computed |  | Current SourceApiAssociation status details. |
| `SourceApiId` | source_api_id | `string` | computed |  | GraphQLApiId of the source API in the association. |
| `SourceApiIdentifier` | source_api_identifier | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | Identifier of the Source GraphQLApi to associate. It could be either GraphQLApi ApiId or ARN |

Supports update: yes

Discovery: supported (parent resource required)
