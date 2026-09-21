# aws.accessgrantslocation

**CloudFormation type:** `AWS::S3::AccessGrantsLocation`

The AWS::S3::AccessGrantsLocation resource is an Amazon S3 resource type hosted in an access grants instance which can be the target of S3 access grants.

Region attribute: `region`

**Import ID:** `<region>/AccessGrantsLocationId` (AWS::S3::AccessGrantsLocation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessGrantsLocationArn` | access_grants_location_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the specified Access Grants location. |
| `AccessGrantsLocationId` | access_grants_location_id | `string` | computed |  | The unique identifier for the specified Access Grants location. |
| `IamRoleArn` | iam_role_arn | `string` | required | aws.role.Arn | The Amazon Resource Name (ARN) of the access grant location's associated IAM role. |
| `LocationScope` | location_scope | `string` | required |  | Descriptor for where the location actually points |
| `Tags` |  | `map` | replaces on change, tags map |  |  |

Supports update: yes

Discovery: supported
