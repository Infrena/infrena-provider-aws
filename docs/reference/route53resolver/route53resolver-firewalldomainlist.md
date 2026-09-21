# aws.route53resolver.firewalldomainlist

**CloudFormation type:** `AWS::Route53Resolver::FirewallDomainList`

Resource schema for AWS::Route53Resolver::FirewallDomainList.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Resolver::FirewallDomainList)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Arn |
| `CreationTime` | creation_time | `string` | computed |  | Rfc3339TimeString |
| `CreatorRequestId` | creator_request_id | `string` | computed |  | The id of the creator request. |
| `DomainCount` | domain_count | `integer` | computed |  | Count |
| `DomainFileUrl` | domain_file_url | `string` | optional, computed, provider-chosen, write-only |  | S3 URL to import domains from. |
| `Domains` |  | `list` | optional, computed, provider-chosen, write-only |  | An inline list of domains to use for this domain list. |
| `Id` |  | `string` | computed |  | ResourceId |
| `ManagedOwnerName` | managed_owner_name | `string` | computed |  | ServicePrincipal |
| `ModificationTime` | modification_time | `string` | computed |  | Rfc3339TimeString |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | FirewallDomainListName |
| `Status` |  | `string` | computed |  | ResolverFirewallDomainList, possible values are COMPLETE, DELETING, UPDATING, COMPLETE_IMPORT_FAILED, IMPORTING, and INACTIVE_OWNER_ACCOUNT_CLOSED. |
| `StatusMessage` | status_message | `string` | computed |  | FirewallDomainListAssociationStatus |
| `Tags` |  | `map` | tags map |  | Tags |

Supports update: yes

Discovery: supported
