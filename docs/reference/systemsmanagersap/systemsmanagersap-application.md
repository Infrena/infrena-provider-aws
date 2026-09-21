# aws.systemsmanagersap.application

**CloudFormation type:** `AWS::SystemsManagerSAP::Application`

Resource schema for AWS::SystemsManagerSAP::Application

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SystemsManagerSAP::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationId` | application_id | `string` | required | aws.systemsmanagersap.application.ApplicationId |  |
| `ApplicationType` | application_type | `string` | required |  |  |
| `Arn` |  | `string` | computed |  | The ARN of the SSM-SAP application |
| `ComponentsInfo` | components_info | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | This is an optional parameter for component details to which the SAP ABAP application is attached, such as Web Dispatcher. |
| `Credentials` |  | `list` | optional, computed, provider-chosen, replaces on change, sensitive, write-only |  |  |
| `DatabaseArn` | database_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ARN of the SAP HANA database |
| `Instances` |  | `list` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `SapInstanceNumber` | sap_instance_number | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Sid` |  | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Tags` |  | `map` | tags map |  | The tags of a SystemsManagerSAP application. |

Supports update: yes

Discovery: supported
