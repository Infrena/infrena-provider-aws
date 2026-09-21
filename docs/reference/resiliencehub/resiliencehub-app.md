# aws.resiliencehub.app

**CloudFormation type:** `AWS::ResilienceHub::App`

Resource Type Definition for AWS::ResilienceHub::App.

Region attribute: `region`

**Import ID:** `<region>/AppArn` (AWS::ResilienceHub::App)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppArn` | app_arn | `string` | computed |  | Amazon Resource Name (ARN) of the App. |
| `AppAssessmentSchedule` | app_assessment_schedule | `string` | optional, computed, provider-chosen |  | Assessment execution schedule. |
| `AppTemplateBody` | app_template_body | `string` | required |  | A string containing full ResilienceHub app template body. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | App description. |
| `DriftStatus` | drift_status | `string` | computed |  | Indicates if compliance drifts (deviations) were detected while running an assessment for your application. |
| `EventSubscriptions` | event_subscriptions | `list` | optional, computed, provider-chosen |  | The list of events you would like to subscribe and get notification for. |
| `Name` |  | `string` | required, replaces on change |  | Name of the app. |
| `PermissionModel` | permission_model | `map` | optional, computed, provider-chosen |  | Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment. |
| `ResiliencyPolicyArn` | resiliency_policy_arn | `string` | optional, computed, provider-chosen | aws.resiliencypolicy.PolicyArn | Amazon Resource Name (ARN) of the Resiliency Policy. |
| `ResourceMappings` | resource_mappings | `list` | required |  | An array of ResourceMapping objects. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
