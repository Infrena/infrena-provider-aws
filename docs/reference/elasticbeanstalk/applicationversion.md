# aws.applicationversion

**CloudFormation type:** `AWS::ElasticBeanstalk::ApplicationVersion`

Resource Type definition for AWS::ElasticBeanstalk::ApplicationVersion

Region attribute: `region`

**Import ID:** `<region>/ApplicationName|Id` (AWS::ElasticBeanstalk::ApplicationVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationName` | application_name | `string` | required, replaces on change |  | The name of the Elastic Beanstalk application that is associated with this application version. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of this application version. |
| `Id` |  | `string` | computed |  |  |
| `SourceBundle` | source_bundle | `map` | required, replaces on change |  | The Amazon S3 bucket and key that identify the location of the source bundle for this version. |

Supports update: yes

Discovery: supported
