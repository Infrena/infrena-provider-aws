# aws.templategroupaccesscontrolentry

**CloudFormation type:** `AWS::PCAConnectorAD::TemplateGroupAccessControlEntry`

Definition of AWS::PCAConnectorAD::TemplateGroupAccessControlEntry Resource Type

Region attribute: `region`

**Import ID:** `<region>/GroupSecurityIdentifier|TemplateArn` (AWS::PCAConnectorAD::TemplateGroupAccessControlEntry)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessRights` | access_rights | `map` | required, write-only |  |  |
| `GroupDisplayName` | group_display_name | `string` | required, write-only |  |  |
| `GroupSecurityIdentifier` | group_security_identifier | `string` | required, replaces on change |  |  |
| `TemplateArn` | template_arn | `string` | required, replaces on change | aws.pcaconnectorad.template.TemplateArn |  |

Supports update: yes

Discovery: supported (parent resource required)
