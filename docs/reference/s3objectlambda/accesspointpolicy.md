# aws.accesspointpolicy

**CloudFormation type:** `AWS::S3ObjectLambda::AccessPointPolicy`

AWS::S3ObjectLambda::AccessPointPolicy resource is an Amazon S3ObjectLambda policy type that you can use to control permissions for your S3ObjectLambda

Region attribute: `region`

**Import ID:** `<region>/ObjectLambdaAccessPoint` (AWS::S3ObjectLambda::AccessPointPolicy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ObjectLambdaAccessPoint` | object_lambda_access_point | `string` | required, replaces on change |  | The name of the Amazon S3 ObjectLambdaAccessPoint to which the policy applies. |
| `PolicyDocument` | policy_document | `map` | required |  | A policy document containing permissions to add to the specified ObjectLambdaAccessPoint. For more information, see Access Policy Language Overview (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the Amazon Simple Storage Service Developer Guide. |

Supports update: yes

Discovery: not supported
