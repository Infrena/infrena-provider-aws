# aws.pcaconnectorad.template

**CloudFormation type:** `AWS::PCAConnectorAD::Template`

Represents a template that defines certificate configurations, both for issuance and client handling

Region attribute: `region`

**Import ID:** `<region>/TemplateArn` (AWS::PCAConnectorAD::Template)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ConnectorArn` | connector_arn | `string` | required, replaces on change | aws.pcaconnectorad.connector.ConnectorArn |  |
| `Definition` |  | `string` | required |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `ReenrollAllCertificateHolders` | reenroll_all_certificate_holders | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `TemplateArn` | template_arn | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
