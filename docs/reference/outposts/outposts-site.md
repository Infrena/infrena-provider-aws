# aws.outposts.site

**CloudFormation type:** `AWS::Outposts::Site`

Definition of AWS::Outposts::Site Resource Type

Region attribute: `region`

**Import ID:** `<region>/SiteArn` (AWS::Outposts::Site)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `Notes` |  | `string` | optional, computed, provider-chosen |  |  |
| `OperatingAddress` | operating_address | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `RackPhysicalProperties` | rack_physical_properties | `map` | optional, computed, provider-chosen |  |  |
| `ShippingAddress` | shipping_address | `map` | optional, computed, provider-chosen, replaces on change |  |  |
| `SiteArn` | site_arn | `string` | computed |  |  |
| `SiteId` | site_id | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
