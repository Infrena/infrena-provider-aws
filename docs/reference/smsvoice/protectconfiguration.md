# aws.protectconfiguration

**CloudFormation type:** `AWS::SMSVOICE::ProtectConfiguration`

Resource Type definition for AWS::SMSVOICE::ProtectConfiguration

Region attribute: `region`

**Import ID:** `<region>/ProtectConfigurationId` (AWS::SMSVOICE::ProtectConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the protect configuration. |
| `CountryRuleSet` | country_rule_set | `map` | optional, computed, provider-chosen |  | An array of CountryRule containing the rules for the NumberCapability. |
| `DeletionProtectionEnabled` | deletion_protection_enabled | `boolean` | optional, computed, provider-chosen |  | When set to true deletion protection is enabled and protect configuration cannot be deleted. By default this is set to false. |
| `ProtectConfigurationId` | protect_configuration_id | `string` | computed |  | The unique identifier for the protect configuration. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
