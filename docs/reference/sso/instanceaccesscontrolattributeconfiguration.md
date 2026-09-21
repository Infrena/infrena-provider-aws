# aws.instanceaccesscontrolattributeconfiguration

**CloudFormation type:** `AWS::SSO::InstanceAccessControlAttributeConfiguration`

Resource Type definition for SSO InstanceAccessControlAttributeConfiguration

Region attribute: `region`

**Import ID:** `<region>/InstanceArn` (AWS::SSO::InstanceAccessControlAttributeConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessControlAttributes` | access_control_attributes | `list` | optional, computed, provider-chosen |  |  |
| `InstanceAccessControlAttributeConfiguration` | instance_access_control_attribute_configuration | `map` | optional, computed, provider-chosen |  | The InstanceAccessControlAttributeConfiguration property has been deprecated but is still supported for backwards compatibility purposes. We recomend that you use  AccessControlAttributes property instead. |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.sso.instance.InstanceArn | The ARN of the AWS SSO instance under which the operation will be executed. |

Supports update: yes

Discovery: supported (parent resource required)
