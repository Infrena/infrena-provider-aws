# aws.iotsitewise.accesspolicy

**CloudFormation type:** `AWS::IoTSiteWise::AccessPolicy`

Resource schema for AWS::IoTSiteWise::AccessPolicy

Region attribute: `region`

**Import ID:** `<region>/AccessPolicyId` (AWS::IoTSiteWise::AccessPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessPolicyArn` | access_policy_arn | `string` | computed |  | The ARN of the access policy. |
| `AccessPolicyId` | access_policy_id | `string` | computed |  | The ID of the access policy. |
| `AccessPolicyIdentity` | access_policy_identity | `map` | required |  | The identity for this access policy. Choose either an SSO user or group or an IAM user or role. |
| `AccessPolicyPermission` | access_policy_permission | `string` | required |  | The permission level for this access policy. Valid values are ADMINISTRATOR or VIEWER. |
| `AccessPolicyResource` | access_policy_resource | `map` | required |  | The AWS IoT SiteWise Monitor resource for this access policy. Choose either portal or project but not both. |

Supports update: yes

Discovery: supported
