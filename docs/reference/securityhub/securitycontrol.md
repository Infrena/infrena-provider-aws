# aws.securitycontrol

**CloudFormation type:** `AWS::SecurityHub::SecurityControl`

A security control in Security Hub describes a security best practice related to a specific resource.

Region attribute: `region`

**Import ID:** `<region>/SecurityControlId` (AWS::SecurityHub::SecurityControl)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LastUpdateReason` | last_update_reason | `string` | optional, computed, provider-chosen |  | The most recent reason for updating the customizable properties of a security control. This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores. |
| `Parameters` |  | `map` | required |  | An object that identifies the name of a control parameter, its current value, and whether it has been customized. |
| `SecurityControlArn` | security_control_arn | `string` | optional, computed, provider-chosen | aws.securitycontrol.SecurityControlArn | The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard. |
| `SecurityControlId` | security_control_id | `string` | optional, computed, provider-chosen, replaces on change | aws.securitycontrol.SecurityControlId | The unique identifier of a security control across standards. Values for this field typically consist of an AWS service name and a number, such as APIGateway.3. |

Supports update: yes

Discovery: supported
