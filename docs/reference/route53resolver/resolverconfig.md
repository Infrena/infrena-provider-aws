# aws.resolverconfig

**CloudFormation type:** `AWS::Route53Resolver::ResolverConfig`

Resource schema for AWS::Route53Resolver::ResolverConfig.

Region attribute: `region`

**Import ID:** `<region>/ResourceId` (AWS::Route53Resolver::ResolverConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AutodefinedReverse` | autodefined_reverse | `string` | computed |  | ResolverAutodefinedReverseStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED. |
| `AutodefinedReverseFlag` | autodefined_reverse_flag | `string` | required, replaces on change |  | Represents the desired status of AutodefinedReverse. The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED). |
| `Id` |  | `string` | computed |  | Id |
| `OwnerId` | owner_id | `string` | computed |  | AccountId |
| `ResourceId` | resource_id | `string` | required, replaces on change |  | ResourceId |

Supports update: no

Discovery: supported
