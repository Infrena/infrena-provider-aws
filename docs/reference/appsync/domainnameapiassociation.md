# aws.domainnameapiassociation

**CloudFormation type:** `AWS::AppSync::DomainNameApiAssociation`

Resource Type definition for AWS::AppSync::DomainNameApiAssociation

Region attribute: `region`

**Import ID:** `<region>/ApiAssociationIdentifier` (AWS::AppSync::DomainNameApiAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiAssociationIdentifier` | api_association_identifier | `string` | computed |  |  |
| `ApiId` | api_id | `string` | required | aws.appsync.api.ApiId |  |
| `DomainName` | domain_name | `string` | required, replaces on change |  |  |

Supports update: yes

Discovery: not supported
