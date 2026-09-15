# aws.domainnameaccessassociation

**CloudFormation type:** `AWS::ApiGateway::DomainNameAccessAssociation`

Resource Type definition for AWS::ApiGateway::DomainNameAccessAssociation.

Region attribute: `region`

**Import ID:** `<region>/DomainNameAccessAssociationArn` (AWS::ApiGateway::DomainNameAccessAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessAssociationSource` | access_association_source | `string` | required, replaces on change |  | The source of the domain name access association resource. |
| `AccessAssociationSourceType` | access_association_source_type | `string` | required, replaces on change |  | The source type of the domain name access association resource. |
| `DomainNameAccessAssociationArn` | domain_name_access_association_arn | `string` | computed |  | The amazon resource name (ARN) of the domain name access association resource. |
| `DomainNameArn` | domain_name_arn | `string` | required, replaces on change | aws.apigateway.domainname.DomainNameArn | The amazon resource name (ARN) of the domain name resource. |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  | An array of arbitrary tags (key-value pairs) to associate with the domainname access association. |

Supports update: no

Discovery: supported
