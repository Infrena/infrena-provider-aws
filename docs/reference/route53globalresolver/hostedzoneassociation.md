# aws.hostedzoneassociation

**CloudFormation type:** `AWS::Route53GlobalResolver::HostedZoneAssociation`

Resource schema for AWS::Route53GlobalResolver::HostedZoneAssociation

Region attribute: `region`

**Import ID:** `<region>/HostedZoneAssociationId` (AWS::Route53GlobalResolver::HostedZoneAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  |  |
| `HostedZoneAssociationId` | hosted_zone_association_id | `string` | computed |  |  |
| `HostedZoneId` | hosted_zone_id | `string` | required, replaces on change | aws.hostedzone.Id |  |
| `HostedZoneName` | hosted_zone_name | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  |  |
| `Status` |  | `string` | computed |  |  |
| `UpdatedAt` | updated_at | `string` | computed |  |  |

Supports update: yes

Discovery: supported (parent resource required)
