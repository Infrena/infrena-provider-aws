# aws.linkroutingrule

**CloudFormation type:** `AWS::RTBFabric::LinkRoutingRule`

Resource Type definition for AWS::RTBFabric::LinkRoutingRule. A routing rule on a link within RTB Fabric that controls request routing based on conditions such as host headers, path matching, and query string parameters.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::RTBFabric::LinkRoutingRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Conditions` |  | `map` | required |  | Conditions for a routing rule. All non-null fields must match (AND logic). At least one field must be set. HostHeader and HostHeaderWildcard are mutually exclusive. PathPrefix and PathExact are mutually exclusive. |
| `CreatedTimestamp` | created_timestamp | `string` | computed |  |  |
| `GatewayId` | gateway_id | `string` | required |  |  |
| `LinkId` | link_id | `string` | required | aws.rtbfabric.link.LinkId |  |
| `Priority` |  | `integer` | required |  |  |
| `RuleId` | rule_id | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Tags to assign to the LinkRoutingRule. |
| `UpdatedTimestamp` | updated_timestamp | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
