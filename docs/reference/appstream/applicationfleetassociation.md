# aws.applicationfleetassociation

**CloudFormation type:** `AWS::AppStream::ApplicationFleetAssociation`

Resource Type definition for AWS::AppStream::ApplicationFleetAssociation

Region attribute: `region`

**Import ID:** `<region>/FleetName|ApplicationArn` (AWS::AppStream::ApplicationFleetAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationArn` | application_arn | `string` | required, replaces on change | aws.appstream.application.Arn |  |
| `FleetName` | fleet_name | `string` | required, replaces on change |  |  |

Supports update: no

Discovery: not supported
