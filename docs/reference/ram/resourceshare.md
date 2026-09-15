# aws.resourceshare

**CloudFormation type:** `AWS::RAM::ResourceShare`

Resource type definition for AWS::RAM::ResourceShare

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RAM::ResourceShare)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowExternalPrincipals` | allow_external_principals | `boolean` | optional, computed, provider-chosen |  | Specifies whether principals outside your organization in AWS Organizations can be associated with a resource share. A value of `true` lets you share with individual AWS accounts that are not in your organization. A value of `false` only has meaning if your account is a member of an AWS Organization. The default value is `true`. |
| `Arn` |  | `string` | computed |  |  |
| `CreationTime` | creation_time | `string` | computed |  | The date and time when the resource share was created. |
| `FeatureSet` | feature_set | `string` | computed |  | The feature set of the resource share. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The date and time when the resource share was last updated. |
| `Name` |  | `string` | required |  | Specifies the name of the resource share. |
| `OwningAccountId` | owning_account_id | `string` | computed |  | The ID of the AWS account that owns the resource share. |
| `PermissionArns` | permission_arns | `list` | optional, computed, provider-chosen, write-only | aws.ram.permission.Arn | Specifies the [Amazon Resource Names (ARNs)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the AWS RAM permission to associate with the resource share. If you do not specify an ARN for the permission, AWS RAM automatically attaches the default version of the permission for each resource type. You can associate only one permission with each resource type included in the resource share. |
| `Principals` |  | `list` | optional, computed, provider-chosen, write-only |  | Specifies the principals to associate with the resource share. The possible values are: |
| `ResourceArns` | resource_arns | `list` | optional, computed, provider-chosen, write-only |  | Specifies a list of one or more ARNs of the resources to associate with the resource share. |
| `ResourceShareConfiguration` | resource_share_configuration | `map` | optional, computed, provider-chosen |  | The configuration for a resource share. |
| `Sources` |  | `list` | optional, computed, provider-chosen, write-only |  | Specifies from which source accounts the service principal has access to the resources in this resource share. |
| `Status` |  | `string` | computed |  | The current status of the resource share. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Specifies one or more tags to attach to the resource share itself. It doesn't attach the tags to the resources associated with the resource share. |

Supports update: yes

Discovery: supported
