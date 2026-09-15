# aws.verifiedaccessgroup

**CloudFormation type:** `AWS::EC2::VerifiedAccessGroup`

The AWS::EC2::VerifiedAccessGroup resource creates an AWS EC2 Verified Access Group.

Region attribute: `region`

**Import ID:** `<region>/VerifiedAccessGroupId` (AWS::EC2::VerifiedAccessGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | Time this Verified Access Group was created. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description for the AWS Verified Access group. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | Time this Verified Access Group was last updated. |
| `Owner` |  | `string` | computed |  | The AWS account number that owns the group. |
| `PolicyDocument` | policy_document | `string` | optional, computed, provider-chosen |  | The AWS Verified Access policy document. |
| `PolicyEnabled` | policy_enabled | `boolean` | optional, computed, provider-chosen |  | The status of the Verified Access policy. |
| `SseSpecification` | sse_specification | `map` | optional, computed, provider-chosen |  | The configuration options for customer provided KMS encryption. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `VerifiedAccessGroupArn` | verified_access_group_arn | `string` | computed |  | The ARN of the Verified Access group. |
| `VerifiedAccessGroupId` | verified_access_group_id | `string` | computed |  | The ID of the AWS Verified Access group. |
| `VerifiedAccessInstanceId` | verified_access_instance_id | `string` | required | aws.verifiedaccessinstance.VerifiedAccessInstanceId | The ID of the AWS Verified Access instance. |

Supports update: yes

Discovery: supported
