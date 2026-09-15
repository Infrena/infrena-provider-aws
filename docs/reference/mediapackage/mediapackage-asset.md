# aws.mediapackage.asset

**CloudFormation type:** `AWS::MediaPackage::Asset`

Resource schema for AWS::MediaPackage::Asset

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::MediaPackage::Asset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the Asset. |
| `CreatedAt` | created_at | `string` | computed |  | The time the Asset was initially submitted for Ingest. |
| `EgressEndpoints` | egress_endpoints | `list` | optional, computed, provider-chosen, replaces on change |  | The list of egress endpoints available for the Asset. |
| `Id` |  | `string` | required, replaces on change |  | The unique identifier for the Asset. |
| `PackagingGroupId` | packaging_group_id | `string` | required, replaces on change | aws.packaginggroup.Id | The ID of the PackagingGroup for the Asset. |
| `ResourceId` | resource_id | `string` | optional, computed, provider-chosen, replaces on change |  | The resource ID to include in SPEKE key requests. |
| `SourceArn` | source_arn | `string` | required, replaces on change |  | ARN of the source object in S3. |
| `SourceRoleArn` | source_role_arn | `string` | required, replaces on change | aws.role.Arn | The IAM role_arn used to access the source S3 bucket. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | A collection of tags associated with a resource |

Supports update: no

Discovery: supported
