# aws.provisioningtemplate

**CloudFormation type:** `AWS::IoT::ProvisioningTemplate`

Creates a fleet provisioning template.

Region attribute: `region`

**Import ID:** `<region>/TemplateName` (AWS::IoT::ProvisioningTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  |  |
| `PreProvisioningHook` | pre_provisioning_hook | `map` | optional, computed, provider-chosen |  |  |
| `ProvisioningRoleArn` | provisioning_role_arn | `string` | required | aws.role.Arn |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `TemplateArn` | template_arn | `string` | computed |  |  |
| `TemplateBody` | template_body | `string` | required |  |  |
| `TemplateName` | template_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `TemplateType` | template_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
