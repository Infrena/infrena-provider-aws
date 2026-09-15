# aws.elasticbeanstalk.application

**CloudFormation type:** `AWS::ElasticBeanstalk::Application`

The AWS::ElasticBeanstalk::Application resource specifies an Elastic Beanstalk application.

Region attribute: `region`

**Import ID:** `<region>/ApplicationName` (AWS::ElasticBeanstalk::Application)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationName` | application_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the Elastic Beanstalk application. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Your description of the application. |
| `ResourceLifecycleConfig` | resource_lifecycle_config | `map` | optional, computed, provider-chosen |  | Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions. |

Supports update: yes

Discovery: supported
