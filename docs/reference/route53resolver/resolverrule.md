# aws.resolverrule

**CloudFormation type:** `AWS::Route53Resolver::ResolverRule`

Resource Type definition for AWS::Route53Resolver::ResolverRule

Region attribute: `region`

**Import ID:** `<region>/ResolverRuleId` (AWS::Route53Resolver::ResolverRule)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the resolver rule. |
| `DelegationRecord` | delegation_record | `string` | optional, computed, provider-chosen |  | The name server domain for queries to be delegated to if a query matches the delegation record. |
| `DomainName` | domain_name | `string` | optional, computed, provider-chosen |  | DNS queries for this domain name are forwarded to the IP addresses that are specified in TargetIps |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The name for the Resolver rule |
| `ResolverEndpointId` | resolver_endpoint_id | `string` | optional, computed, provider-chosen | aws.resolverendpoint.ResolverEndpointId | The ID of the endpoint that the rule is associated with. |
| `ResolverRuleId` | resolver_rule_id | `string` | computed |  | The ID of the endpoint that the rule is associated with. |
| `RuleType` | rule_type | `string` | required, replaces on change |  | When you want to forward DNS queries for specified domain name to resolvers on your network, specify FORWARD. When you have a forwarding rule to forward DNS queries for a domain to your network and you want Resolver to process queries for a subdomain of that domain, specify SYSTEM. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `TargetIps` | target_ips | `list` | optional, computed, provider-chosen |  | An array that contains the IP addresses and ports that an outbound endpoint forwards DNS queries to. Typically, these are the IP addresses of DNS resolvers on your network. Specify IPv4 addresses. IPv6 is not supported. |

Supports update: yes

Discovery: supported
