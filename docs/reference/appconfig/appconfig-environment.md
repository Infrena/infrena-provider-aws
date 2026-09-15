# aws.appconfig.environment

**CloudFormation type:** `AWS::AppConfig::Environment`

Resource Type definition for AWS::AppConfig::Environment

Region attribute: `region`

**Import ID:** `<region>/ApplicationId|EnvironmentId` (AWS::AppConfig::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required, replaces on change | aws.appconfig.application.ApplicationId | The application ID. |
| `DeletionProtectionCheck` | deletion_protection_check | `string` | optional, computed, provider-chosen, write-only |  | On resource deletion this controls whether the Deletion Protection check should be applied, bypassed, or (the default) whether the behavior should be controlled by the account-level Deletion Protection setting. See https://docs.aws.amazon.com/appconfig/latest/userguide/deletion-protection.html |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the environment. |
| `EnvironmentId` | environment_id | `string` | computed |  | The environment ID. |
| `Monitors` |  | `list` | optional, computed, provider-chosen |  | Amazon CloudWatch alarms to monitor during the deployment process. |
| `Name` |  | `string` | required |  | A name for the environment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Metadata to assign to the environment. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. |

Supports update: yes

Discovery: supported (parent resource required)
