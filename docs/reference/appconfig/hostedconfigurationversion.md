# aws.hostedconfigurationversion

**CloudFormation type:** `AWS::AppConfig::HostedConfigurationVersion`

Resource Type definition for AWS::AppConfig::HostedConfigurationVersion

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|ConfigurationProfileId|VersionNumber` (AWS::AppConfig::HostedConfigurationVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.appconfig.application.ApplicationId | The application ID. |
| `ConfigurationProfileId` | configuration_profile_id | `string` | required, replaces on change | aws.configurationprofile.ConfigurationProfileId | The configuration profile ID. |
| `Content` |  | `string` | required, replaces on change |  | The content of the configuration or the configuration data. |
| `ContentType` | content_type | `string` | required, replaces on change |  | A standard MIME type describing the format of the configuration content. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description of the hosted configuration version. |
| `LatestVersionNumber` | latest_version_number | `integer` | optional, computed, provider-chosen, replaces on change, write-only |  | An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version. To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version. |
| `VersionLabel` | version_label | `string` | optional, computed, provider-chosen, replaces on change |  | A user-defined label for an AWS AppConfig hosted configuration version. |
| `VersionNumber` | version_number | `string` | computed |  | Current version number of hosted configuration version. |

Supports update: no

Discovery: supported (parent resource required)
