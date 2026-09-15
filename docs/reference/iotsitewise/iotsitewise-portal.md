# aws.iotsitewise.portal

**CloudFormation type:** `AWS::IoTSiteWise::Portal`

Resource schema for AWS::IoTSiteWise::Portal

Region attribute: `region`

**Import ID:** `<region>/PortalId` (AWS::IoTSiteWise::Portal)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alarms` |  | `map` | optional, computed, provider-chosen |  | Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal. You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range. |
| `NotificationSenderEmail` | notification_sender_email | `string` | optional, computed, provider-chosen |  | The email address that sends alarm notifications. |
| `PortalArn` | portal_arn | `string` | computed |  | The ARN of the portal, which has the following format. |
| `PortalAuthMode` | portal_auth_mode | `string` | optional, computed, provider-chosen, replaces on change |  | The service to use to authenticate users to the portal. Choose from SSO or IAM. You can't change this value after you create a portal. |
| `PortalClientId` | portal_client_id | `string` | computed |  | The AWS SSO application generated client ID (used with AWS SSO APIs). |
| `PortalContactEmail` | portal_contact_email | `string` | required |  | The AWS administrator's contact email address. |
| `PortalDescription` | portal_description | `string` | optional, computed, provider-chosen |  | A description for the portal. |
| `PortalId` | portal_id | `string` | computed |  | The ID of the portal. |
| `PortalName` | portal_name | `string` | required |  | A friendly name for the portal. |
| `PortalStartUrl` | portal_start_url | `string` | computed |  | The public root URL for the AWS IoT AWS IoT SiteWise Monitor application portal. |
| `PortalType` | portal_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of portal |
| `PortalTypeConfiguration` | portal_type_configuration | `map` | optional, computed, provider-chosen |  | Map to associate detail of configuration related with a PortalType. |
| `RoleArn` | role_arn | `string` | required | aws.role.Arn | The ARN of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs that contain metadata for the portal. |

Supports update: yes

Discovery: supported
