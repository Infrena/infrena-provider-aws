# aws.license

**CloudFormation type:** `AWS::LicenseManager::License`

Resource Type definition for AWS::LicenseManager::License

Region attribute: `region`

**Import ID:** `<region>/LicenseArn` (AWS::LicenseManager::License)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Beneficiary` |  | `string` | required |  | Beneficiary of the license. |
| `ConsumptionConfiguration` | consumption_configuration | `map` | required |  |  |
| `Entitlements` |  | `list` | required |  |  |
| `HomeRegion` | home_region | `string` | required |  | Home region for the created license. |
| `Issuer` |  | `map` | required |  |  |
| `LicenseArn` | license_arn | `string` | computed |  | Amazon Resource Name is a unique name for each resource. |
| `LicenseMetadata` | license_metadata | `list` | optional, computed, provider-chosen |  |  |
| `LicenseName` | license_name | `string` | required |  | Name for the created license. |
| `ProductName` | product_name | `string` | required |  | Product name for the created license. |
| `ProductSKU` | product_sku | `string` | required |  | ProductSKU of the license. |
| `Status` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to attach. |
| `Validity` |  | `map` | required |  |  |
| `Version` |  | `string` | computed |  | The version of the license. |

Supports update: yes

Discovery: supported
