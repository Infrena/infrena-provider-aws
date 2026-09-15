# aws.s3tableintegration

**CloudFormation type:** `AWS::ObservabilityAdmin::S3TableIntegration`

Resource Type definition for a CloudWatch Observability Admin S3 Table Integration.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ObservabilityAdmin::S3TableIntegration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of the S3 Table Integration |
| `Encryption` |  | `map` | required, replaces on change |  | Encryption configuration for the S3 Table Integration |
| `LogSources` | log_sources | `list` | optional, computed, provider-chosen |  | The CloudWatch Logs data sources to associate with the S3 Table Integration |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The ARN of the role used to access the S3 Table Integration |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource |

Supports update: yes

Discovery: supported
