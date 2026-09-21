# aws.grafana.workspace

**CloudFormation type:** `AWS::Grafana::Workspace`

Definition of AWS::Grafana::Workspace Resource Type

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Grafana::Workspace)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountAccessType` | account_access_type | `string` | required |  | These enums represent valid account access types. Specifically these enums determine whether the workspace can access AWS resources in the AWS account only, or whether it can also access resources in other accounts in the same organization. If the value CURRENT_ACCOUNT is used, a workspace role ARN must be provided. If the value is ORGANIZATION, a list of organizational units must be provided. |
| `AuthenticationProviders` | authentication_providers | `list` | required |  | List of authentication providers to enable. |
| `ClientToken` | client_token | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | A unique, case-sensitive, user-provided identifier to ensure the idempotency of the request. |
| `CreationTimestamp` | creation_timestamp | `string` | computed |  | Timestamp when the workspace was created. |
| `DataSources` | data_sources | `list` | optional, computed, provider-chosen |  | List of data sources on the service managed IAM role. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Description of a workspace. |
| `Endpoint` |  | `string` | computed |  | Endpoint for the Grafana workspace. |
| `GrafanaVersion` | grafana_version | `string` | optional, computed, provider-chosen |  | The version of Grafana to support in your workspace. |
| `Id` |  | `string` | computed |  | The id that uniquely identifies a Grafana workspace. |
| `ModificationTimestamp` | modification_timestamp | `string` | computed |  | Timestamp when the workspace was last modified |
| `Name` |  | `string` | optional, computed, provider-chosen |  | The user friendly name of a workspace. |
| `NetworkAccessControl` | network_access_control | `map` | optional, computed, provider-chosen |  | The configuration settings for Network Access Control. |
| `NotificationDestinations` | notification_destinations | `list` | optional, computed, provider-chosen |  | List of notification destinations on the customers service managed IAM role that the Grafana workspace can query. |
| `OrganizationRoleName` | organization_role_name | `string` | optional, computed, provider-chosen |  | The name of an IAM role that already exists to use with AWS Organizations to access AWS data sources and notification channels in other accounts in an organization. |
| `OrganizationalUnits` | organizational_units | `list` | optional, computed, provider-chosen |  | List of Organizational Units containing AWS accounts the Grafana workspace can pull data from. |
| `PermissionType` | permission_type | `string` | required |  | These enums represent valid permission types to use when creating or configuring a Grafana workspace. The SERVICE_MANAGED permission type means the Managed Grafana service will create a workspace IAM role on your behalf. The CUSTOMER_MANAGED permission type means that the customer is expected to provide an IAM role that the Grafana workspace can use to query data sources. |
| `PluginAdminEnabled` | plugin_admin_enabled | `boolean` | optional, computed, provider-chosen |  | Allow workspace admins to install plugins |
| `RoleArn` | role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn | IAM Role that will be used to grant the Grafana workspace access to a customers AWS resources. |
| `SamlConfiguration` | saml_configuration | `map` | optional, computed, provider-chosen |  | SAML configuration data associated with an AMG workspace. |
| `SamlConfigurationStatus` | saml_configuration_status | `string` | computed |  | Valid SAML configuration statuses. |
| `SsoClientId` | sso_client_id | `string` | computed |  | The client ID of the AWS SSO Managed Application. |
| `StackSetName` | stack_set_name | `string` | optional, computed, provider-chosen |  | The name of the AWS CloudFormation stack set to use to generate IAM roles to be used for this workspace. |
| `Status` |  | `string` | computed |  | These enums represent the status of a workspace. |
| `Tags` |  | `map` | tags map |  | The list of tags associated with the workspace. |
| `VpcConfiguration` | vpc_configuration | `map` | optional, computed, provider-chosen |  | The configuration settings for an Amazon VPC that contains data sources for your Grafana workspace to connect to. |

Supports update: yes

Discovery: supported
