# aws.meteredproduct

**CloudFormation type:** `AWS::Deadline::MeteredProduct`

Resource Type definition for AWS::Deadline::MeteredProduct

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Deadline::MeteredProduct)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Family` |  | `string` | computed |  |  |
| `LicenseEndpointId` | license_endpoint_id | `string` | optional, computed, provider-chosen, replaces on change | aws.licenseendpoint.LicenseEndpointId |  |
| `Port` |  | `integer` | computed |  |  |
| `ProductId` | product_id | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Vendor` |  | `string` | computed |  |  |

Supports update: no

Discovery: supported (parent resource required)
