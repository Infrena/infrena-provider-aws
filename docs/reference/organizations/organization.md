# aws.organization

**CloudFormation type:** `AWS::Organizations::Organization`

Resource schema for AWS::Organizations::Organization

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::Organizations::Organization)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of an organization. |
| `FeatureSet` | feature_set | `string` | optional, computed, provider-chosen |  | Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality. |
| `Id` |  | `string` | computed |  | The unique identifier (ID) of an organization. |
| `ManagementAccountArn` | management_account_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the account that is designated as the management account for the organization. |
| `ManagementAccountEmail` | management_account_email | `string` | computed |  | The email address that is associated with the AWS account that is designated as the management account for the organization. |
| `ManagementAccountId` | management_account_id | `string` | computed |  | The unique identifier (ID) of the management account of an organization. |
| `RootId` | root_id | `string` | computed |  | The unique identifier (ID) for the root. |

Supports update: yes

Discovery: supported
