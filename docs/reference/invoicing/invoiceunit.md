# aws.invoiceunit

**CloudFormation type:** `AWS::Invoicing::InvoiceUnit`

An invoice unit is a set of mutually exclusive accounts that correspond to your business entity. Invoice units allow you to separate AWS account costs and configures your invoice for each business entity.

Region attribute: `region`

**Import ID:** `<region>/InvoiceUnitArn` (AWS::Invoicing::InvoiceUnit)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `InvoiceReceiver` | invoice_receiver | `string` | required, replaces on change |  |  |
| `InvoiceUnitArn` | invoice_unit_arn | `string` | computed |  |  |
| `LastModified` | last_modified | `float` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `ResourceTags` | resource_tags | `map` | tags map |  |  |
| `Rule` |  | `map` | required |  |  |
| `TaxInheritanceDisabled` | tax_inheritance_disabled | `boolean` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
