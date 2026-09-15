# aws.codedeploy.application

**CloudFormation type:** `AWS::CodeDeploy::Application`

The AWS::CodeDeploy::Application resource creates an AWS CodeDeploy application

Region attribute: `region`

**Import ID:** `<region>/ApplicationName` (AWS::CodeDeploy::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationName` | application_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name. |
| `ComputePlatform` | compute_platform | `string` | optional, computed, provider-chosen, replaces on change |  | The compute platform that CodeDeploy deploys the application to. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The metadata that you apply to CodeDeploy applications to help you organize and categorize them. Each tag consists of a key and an optional value, both of which you define. |

Supports update: yes

Discovery: supported
