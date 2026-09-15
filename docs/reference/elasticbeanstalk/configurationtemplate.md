# aws.configurationtemplate

**CloudFormation type:** `AWS::ElasticBeanstalk::ConfigurationTemplate`

Resource Type definition for AWS::ElasticBeanstalk::ConfigurationTemplate

Region attribute: `region`

**Import ID:** `<region>/ApplicationName|TemplateName` (AWS::ElasticBeanstalk::ConfigurationTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationName` | application_name | `string` | required, replaces on change |  | The name of the Elastic Beanstalk application to associate with this configuration template. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | An optional description for this configuration. |
| `EnvironmentId` | environment_id | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The ID of an environment whose settings you want to use to create the configuration template. You must specify EnvironmentId if you don't specify PlatformArn, SolutionStackName, or SourceConfiguration. |
| `OptionSettings` | option_settings | `list` | optional, computed, provider-chosen |  | Option values for the Elastic Beanstalk configuration, such as the instance type. If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide. |
| `PlatformArn` | platform_arn | `string` | optional, computed, provider-chosen, replaces on change |  | The Amazon Resource Name (ARN) of the custom platform. For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the AWS Elastic Beanstalk Developer Guide. |
| `SolutionStackName` | solution_stack_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses. For example, 64bit Amazon Linux 2013.09 running Tomcat 7 Java 7. A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the AWS Elastic Beanstalk Developer Guide. |
| `SourceConfiguration` | source_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | An Elastic Beanstalk configuration template to base this one on. If specified, Elastic Beanstalk uses the configuration values from the specified configuration template to create a new configuration. |
| `TemplateName` | template_name | `string` | computed |  | The name of the configuration template |

Supports update: yes

Discovery: supported
