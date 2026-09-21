# aws.profileresourceassociation

**CloudFormation type:** `AWS::Route53Profiles::ProfileResourceAssociation`

Resource Type definition for AWS::Route53Profiles::ProfileResourceAssociation

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Profiles::ProfileResourceAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  | Primary Identifier for  Profile Resource Association |
| `Name` |  | `string` | required, replaces on change |  | The name of an association between the  Profile and resource. |
| `ProfileId` | profile_id | `string` | required, replaces on change | aws.route53profiles.profile.Id | The ID of the  profile that you associated the resource to that is specified by ResourceArn. |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The arn of the resource that you associated to the  Profile. |
| `ResourceProperties` | resource_properties | `string` | optional, computed, provider-chosen |  | A JSON-formatted string with key-value pairs specifying the properties of the associated resource. |
| `ResourceType` | resource_type | `string` | computed |  | The type of the resource associated to the  Profile. |

Supports update: yes

Discovery: supported (parent resource required)
