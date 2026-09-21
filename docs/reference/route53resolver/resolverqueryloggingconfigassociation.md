# aws.resolverqueryloggingconfigassociation

**CloudFormation type:** `AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation`

Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Resolver::ResolverQueryLoggingConfigAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | Rfc3339TimeString |
| `Error` |  | `string` | computed |  | ResolverQueryLogConfigAssociationError |
| `ErrorMessage` | error_message | `string` | computed |  | ResolverQueryLogConfigAssociationErrorMessage |
| `Id` |  | `string` | computed |  | Id |
| `ResolverQueryLogConfigId` | resolver_query_log_config_id | `string` | optional, computed, provider-chosen, replaces on change |  | ResolverQueryLogConfigId |
| `ResourceId` | resource_id | `string` | optional, computed, provider-chosen, replaces on change |  | ResourceId |
| `Status` |  | `string` | computed |  | ResolverQueryLogConfigAssociationStatus |

Supports update: no

Discovery: supported
