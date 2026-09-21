# aws.tlsinspectionconfiguration

**CloudFormation type:** `AWS::NetworkFirewall::TLSInspectionConfiguration`

Resource type definition for AWS::NetworkFirewall::TLSInspectionConfiguration

Region attribute: `region`

**Import ID:** `<region>/TLSInspectionConfigurationArn` (AWS::NetworkFirewall::TLSInspectionConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `TLSInspectionConfiguration` | tls_inspection_configuration | `map` | required |  |  |
| `TLSInspectionConfigurationArn` | tls_inspection_configuration_arn | `string` | computed |  | A resource ARN. |
| `TLSInspectionConfigurationId` | tls_inspection_configuration_id | `string` | computed |  |  |
| `TLSInspectionConfigurationName` | tls_inspection_configuration_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
