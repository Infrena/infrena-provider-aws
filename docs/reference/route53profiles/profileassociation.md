# aws.profileassociation

**CloudFormation type:** `AWS::Route53Profiles::ProfileAssociation`

Resource Type definition for AWS::Route53Profiles::ProfileAssociation

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Profiles::ProfileAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | optional, computed, provider-chosen, write-only |  | The Amazon Resource Name (ARN) of the profile association. |
| `Id` |  | `string` | computed |  | Primary Identifier for  Profile Association |
| `Name` |  | `string` | required, replaces on change |  | The name of an association between a  Profile and a VPC. |
| `ProfileId` | profile_id | `string` | required, replaces on change | aws.route53profiles.profile.Id | The ID of the  profile that you associated with the resource that is specified by ResourceId. |
| `ResourceId` | resource_id | `string` | required, replaces on change |  | The resource that you associated the  profile with. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
