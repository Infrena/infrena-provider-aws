# aws.policyassociation

**CloudFormation type:** `AWS::SecurityHub::PolicyAssociation`

The AWS::SecurityHub::PolicyAssociation resource represents the AWS Security Hub Central Configuration Policy associations in your Target. Only the AWS Security Hub delegated administrator can create the resouce from the home region.

Region attribute: `region`

**Import ID:** `<region>/AssociationIdentifier` (AWS::SecurityHub::PolicyAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationIdentifier` | association_identifier | `string` | computed |  | A unique identifier to indicates if the target has an association |
| `AssociationStatus` | association_status | `string` | computed |  | The current status of the association between the specified target and the configuration |
| `AssociationStatusMessage` | association_status_message | `string` | computed |  | An explanation for a FAILED value for AssociationStatus |
| `AssociationType` | association_type | `string` | computed |  | Indicates whether the association between the specified target and the configuration was directly applied by the Security Hub delegated administrator or inherited from a parent |
| `ConfigurationPolicyId` | configuration_policy_id | `string` | required | aws.configurationpolicy.Id | The universally unique identifier (UUID) of the configuration policy or a value of SELF_MANAGED_SECURITY_HUB for a self-managed configuration |
| `TargetId` | target_id | `string` | required, replaces on change |  | The identifier of the target account, organizational unit, or the root |
| `TargetType` | target_type | `string` | required, replaces on change |  | Indicates whether the target is an AWS account, organizational unit, or the organization root |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time, in UTC and ISO 8601 format, that the configuration policy association was last updated |

Supports update: yes

Discovery: supported
