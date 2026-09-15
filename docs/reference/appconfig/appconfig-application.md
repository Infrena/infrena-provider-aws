# aws.appconfig.application

**CloudFormation type:** `AWS::AppConfig::Application`

Resource Type definition for AWS::AppConfig::Application

Region attribute: `region`

**Import ID:** `<region>/ApplicationId` (AWS::AppConfig::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | computed |  | The application Id |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the application. |
| `Name` |  | `string` | required |  | A name for the application. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Metadata to assign to the application. Tags help organize and categorize your AWS AppConfig resources. Each tag consists of a key and an optional value, both of which you define. |

Supports update: yes

Discovery: supported
