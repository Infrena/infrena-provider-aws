# aws.resolverruleassociation

**CloudFormation type:** `AWS::Route53Resolver::ResolverRuleAssociation`

In the response to an [AssociateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverRule.html), [DisassociateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverRule.html), or [ListResolverRuleAssociations](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRuleAssociations.html) request, provides information about an association between a resolver rule and a VPC. The association determines which DNS queries that originate in the VPC are forwarded to your network.

Region attribute: `region`

**Import ID:** `<region>/ResolverRuleAssociationId` (AWS::Route53Resolver::ResolverRuleAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of an association between a Resolver rule and a VPC. |
| `ResolverRuleAssociationId` | resolver_rule_association_id | `string` | computed |  |  |
| `ResolverRuleId` | resolver_rule_id | `string` | required, replaces on change | aws.resolverrule.ResolverRuleId | The ID of the Resolver rule that you associated with the VPC that is specified by ``VPCId``. |
| `VPCId` | vpc_id | `string` | required, replaces on change | aws.vpc.VpcId | The ID of the VPC that you associated the Resolver rule with. |

Supports update: no

Discovery: supported
