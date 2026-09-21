# aws.webaclassociation

**CloudFormation type:** `AWS::WAFv2::WebACLAssociation`

Associates WebACL to Application Load Balancer, CloudFront or API Gateway.

Region attribute: `region`

**Import ID:** `<region>/ResourceArn|WebACLArn` (AWS::WAFv2::WebACLAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  |  |
| `WebACLArn` | web_acl_arn | `string` | required, replaces on change | aws.webacl.Arn |  |

Supports update: yes

Discovery: not supported
