# aws.lambda.permission

**CloudFormation type:** `AWS::Lambda::Permission`

The ``AWS::Lambda::Permission`` resource grants an AWS service or another account permission to use a function. You can apply the policy at the function level, or specify a qualifier to restrict access to a single version or alias. If you use a qualifier, the invoker must use the full Amazon Resource Name (ARN) of that version or alias to invoke the function.

Region attribute: `region`

**Import ID:** `<region>/FunctionName|Id` (AWS::Lambda::Permission)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Action` |  | `string` | required, replaces on change |  | The action that the principal can use on the function. For example, ``lambda:InvokeFunction`` or ``lambda:GetFunction``. |
| `EventSourceToken` | event_source_token | `string` | optional, computed, provider-chosen, replaces on change |  | For Alexa Smart Home functions, a token that the invoker must supply. |
| `FunctionName` | function_name | `string` | required, replaces on change |  | The name or ARN of the Lambda function, version, or alias. |
| `FunctionUrlAuthType` | function_url_auth_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of authentication that your function URL uses. Set to ``AWS_IAM`` if you want to restrict access to authenticated users only. Set to ``NONE`` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Control access to Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html). |
| `Id` |  | `string` | computed |  |  |
| `InvokedViaFunctionUrl` | invoked_via_function_url | `boolean` | optional, computed, provider-chosen, replaces on change |  | Indicates whether the permission applies when the function is invoked through a function URL. |
| `Principal` |  | `string` | required, replaces on change |  | The AWS-service, AWS-account, IAM user, or IAM role that invokes the function. If you specify a service, use ``SourceArn`` or ``SourceAccount`` to limit who can invoke the function through that service. |
| `PrincipalOrgID` | principal_org_id | `string` | optional, computed, provider-chosen, replaces on change |  | The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization. |
| `SourceAccount` | source_account | `string` | optional, computed, provider-chosen, replaces on change |  | For AWS-service, the ID of the AWS-account that owns the resource. Use this together with ``SourceArn`` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account. |
| `SourceArn` | source_arn | `string` | optional, computed, provider-chosen, replaces on change |  | For AWS-services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic. |

Supports update: no

Discovery: supported (parent resource required)
