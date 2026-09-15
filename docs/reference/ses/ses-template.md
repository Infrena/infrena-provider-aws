# aws.ses.template

**CloudFormation type:** `AWS::SES::Template`

Resource Type definition for AWS::SES::Template

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::SES::Template)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the email template. |
| `Template` |  | `map` | optional, computed, provider-chosen |  | The content of the email, composed of a subject line, an HTML part, and a text-only part |

Supports update: yes

Discovery: supported
