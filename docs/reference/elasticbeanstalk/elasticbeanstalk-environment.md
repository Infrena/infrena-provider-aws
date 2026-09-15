# aws.elasticbeanstalk.environment

**CloudFormation type:** `AWS::ElasticBeanstalk::Environment`

Resource Type definition for AWS::ElasticBeanstalk::Environment

Region attribute: `region`

**Import ID:** `<region>/EnvironmentName` (AWS::ElasticBeanstalk::Environment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationName` | application_name | `string` | required, replaces on change |  | The name of the application that is associated with this environment. |
| `CNAMEPrefix` | cname_prefix | `string` | optional, computed, provider-chosen, replaces on change |  | If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | Your description for this environment. |
| `EndpointURL` | endpoint_url | `string` | computed |  |  |
| `EnvironmentName` | environment_name | `string` | optional, computed, provider-chosen, replaces on change |  | A unique name for the environment. |
| `OperationsRole` | operations_role | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role. |
| `OptionSettings` | option_settings | `list` | optional, computed, provider-chosen, write-only |  | Key-value pairs defining configuration options for this environment, such as the instance type. |
| `PlatformArn` | platform_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the custom platform to use with the environment. |
| `SolutionStackName` | solution_stack_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of an Elastic Beanstalk solution stack (platform version) to use with the environment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Specifies the tags applied to resources in the environment. |
| `TemplateName` | template_name | `string` | optional, computed, provider-chosen, write-only |  | The name of the Elastic Beanstalk configuration template to use with the environment. |
| `Tier` |  | `map` | optional, computed, provider-chosen |  | Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks. |
| `VersionLabel` | version_label | `string` | optional, computed, provider-chosen |  | The name of the application version to deploy. |

Supports update: yes

Discovery: supported
