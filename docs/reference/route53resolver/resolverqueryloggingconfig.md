# aws.resolverqueryloggingconfig

**CloudFormation type:** `AWS::Route53Resolver::ResolverQueryLoggingConfig`

Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Resolver::ResolverQueryLoggingConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn |
| `AssociationCount` | association_count | `integer` | computed |  | Count |
| `CreationTime` | creation_time | `string` | computed |  | Rfc3339TimeString |
| `CreatorRequestId` | creator_request_id | `string` | computed |  | The id of the creator request. |
| `DestinationArn` | destination_arn | `string` | optional, computed, provider-chosen, replaces on change |  | destination arn |
| `Id` |  | `string` | computed |  | ResourceId |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | ResolverQueryLogConfigName |
| `OwnerId` | owner_id | `string` | computed |  | AccountId |
| `ShareStatus` | share_status | `string` | computed |  | ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME. |
| `Status` |  | `string` | computed |  | ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: no

Discovery: supported
