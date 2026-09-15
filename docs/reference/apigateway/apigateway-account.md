# aws.apigateway.account

**CloudFormation type:** `AWS::ApiGateway::Account`

The ``AWS::ApiGateway::Account`` resource specifies the IAM role that Amazon API Gateway uses to write API logs to Amazon CloudWatch Logs. To avoid overwriting other roles, you should only have one ``AWS::ApiGateway::Account`` resource per region per account. 

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ApiGateway::Account)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CloudWatchRoleArn` | cloud_watch_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `Id` |  | `string` | computed |  |  |

Supports update: yes

Discovery: not supported
