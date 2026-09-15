# aws.receiptfilter

**CloudFormation type:** `AWS::SES::ReceiptFilter`

Resource Type definition for AWS::SES::ReceiptFilter

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::SES::ReceiptFilter)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Filter` |  | `map` | required, replaces on change |  | A structure that describes the IP address filter to create, which consists of a name, an IP address range, and whether to allow or block mail from it. |
| `Id` |  | `string` | computed |  |  |

Supports update: no

Discovery: supported
