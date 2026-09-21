# aws.sqlhastandbydetectedinstance

**CloudFormation type:** `AWS::EC2::SqlHaStandbyDetectedInstance`

Resource Type definition for AWS::EC2::SqlHaStandbyDetectedInstance

Region attribute: `region`

**Import ID:** `<region>/InstanceId` (AWS::EC2::SqlHaStandbyDetectedInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `HaStatus` | ha_status | `string` | computed |  | The SQL Server high availability status of the EC2 instance. |
| `InstanceId` | instance_id | `string` | required, replaces on change | aws.ec2.instance.InstanceId | The ID of the EC2 instance to enable for SQL Server high availability standby detection. |
| `LastUpdatedTime` | last_updated_time | `string` | computed |  | The timestamp when the EC2 instance's SQL Server high availability status was last updated. |
| `SqlServerCredentials` | sql_server_credentials | `string` | optional, computed, provider-chosen |  | The ARN of the AWS Secrets Manager secret containing SQL Server access credentials to the EC2 instance. If not specified, AWS Systems Manager agent will use default local user credentials. |
| `SqlServerLicenseUsage` | sql_server_license_usage | `string` | computed |  | The SQL Server license type of the EC2 instance. |

Supports update: yes

Discovery: supported
