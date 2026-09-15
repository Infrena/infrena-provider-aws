# aws.organizationalunitassociation

**CloudFormation type:** `AWS::Notifications::OrganizationalUnitAssociation`

Resource Type definition for AWS::Notifications::OrganizationalUnitAssociation

Region attribute: `region`

**Import ID:** `<region>/NotificationConfigurationArn|OrganizationalUnitId` (AWS::Notifications::OrganizationalUnitAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `NotificationConfigurationArn` | notification_configuration_arn | `string` | required, replaces on change | aws.notificationconfiguration.Arn | ARN identifier of the NotificationConfiguration. |
| `OrganizationalUnitId` | organizational_unit_id | `string` | required, replaces on change |  | The ID of the organizational unit. |

Supports update: no

Discovery: supported (parent resource required)
