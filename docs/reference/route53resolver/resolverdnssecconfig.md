# aws.resolverdnssecconfig

**CloudFormation type:** `AWS::Route53Resolver::ResolverDNSSECConfig`

Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Resolver::ResolverDNSSECConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  | Id |
| `OwnerId` | owner_id | `string` | computed |  | AccountId |
| `ResourceId` | resource_id | `string` | optional, computed, provider-chosen, replaces on change |  | ResourceId |
| `ValidationStatus` | validation_status | `string` | computed |  | ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED. |

Supports update: no

Discovery: supported
