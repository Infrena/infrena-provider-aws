# AWS Lambda

AWS Lambda lets you run code without provisioning or managing servers. Define a function with code in S3 or as a container image, grant it an execution role, and invoke it via API, events, or a function URL.

## Lambda Function

A Lambda function is the core compute unit. It requires deployment code (in S3 or as a container image in ECR), an execution role, and a runtime (for zip-based functions).

### Deployment code

The deployment code must already exist before you declare the function. For zip-based functions, upload a `.zip` file to S3 and reference its bucket and key. For container images, push the image to ECR and reference its URI. The plugin does not upload or push code for you.

### Key attributes

- **name** (function_name): Name of the function, up to 64 characters. Defaults to auto-generated name if omitted.
- **role**: ARN of the execution role the function assumes (required). This role grants the function permissions to call AWS services.
- **code**: Deployment package configuration (required). Specify `S3Bucket` and `S3Key` for zip files, or `ImageUri` for container images.
- **package_type**: Deployment type (`Zip` or `Image`). Defaults to `Zip`.
- **handler**: Handler method (required for zip functions). Format: `filename.function_name` (e.g., `index.handler` for Node.js).
- **runtime**: Runtime identifier (required for zip functions). Examples: `python3.12`, `nodejs20.x`, `go1.x`.
- **memory_size**: Memory in MB (128–10,240). Defaults to 128 MB. Increasing memory also increases CPU.
- **timeout**: Execution timeout in seconds (1–900). Defaults to 3 seconds.
- **environment**: Environment variables as key-value pairs to customize function behavior.
- **vpc_config**: VPC configuration (subnets and security groups) if the function needs access to resources in a VPC.
- **layers**: List of Lambda layer ARNs to attach (optional).
- **tracing_config**: X-Ray tracing configuration (`Active` or `PassThrough`).
- **architectures**: Instruction set (`x86_64` or `arm64`). Defaults to `x86_64`.
- **ephemeral_storage**: `/tmp` directory size in MB (512–10,240). Defaults to 512 MB.
- **tags**: Metadata tags for the function

### Example

```yaml
lambda_exec_role:
  type: aws.role
  name: lambda-execution-role
  assume_role_policy: |
    {
      "Version": "2012-10-17",
      "Statement": [
        {
          "Effect": "Allow",
          "Principal": {
            "Service": "lambda.amazonaws.com"
          },
          "Action": "sts:AssumeRole"
        }
      ]
    }
  managed_policy_arns:
    - arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole

api_function:
  type: aws.lambda.function
  name: api-handler
  role: ${lambda_exec_role.Arn}
  handler: index.handler
  runtime: python3.12
  memory_size: 512
  timeout: 30
  code:
    S3Bucket: my-lambda-code
    S3Key: api-handler.zip
  environment:
    Variables:
      ENVIRONMENT: production
      LOG_LEVEL: INFO
  tags:
    Application: myapp
```

## Lambda Permission

A Lambda permission grants an AWS service, account, or principal permission to invoke your function. Permissions are required before external services (like API Gateway, SNS, S3) can invoke your function.

### Key attributes

- **function_name**: Name or ARN of the Lambda function to grant permission to (required). Can include a version or alias.
- **principal**: The AWS service, account, or principal allowed to invoke (required). Use `s3.amazonaws.com` for S3, `apigateway.amazonaws.com` for API Gateway, or an AWS account ID.
- **action**: The action allowed (required). Usually `lambda:InvokeFunction`.
- **source_arn**: ARN of the resource invoking the function (optional). For example, an S3 bucket ARN or API Gateway ARN. Restricts which sources can invoke.
- **source_account**: AWS account ID of the invoker (optional). Use with `source_arn` to restrict access to a specific account.
- **principal_org_id**: Organization ID for cross-account permissions within an AWS Organization (optional).

### Example

```yaml
api_permission:
  type: aws.lambda.permission
  function_name: ${api_function.FunctionName}
  principal: apigateway.amazonaws.com
  action: lambda:InvokeFunction
  source_arn: arn:aws:execute-api:us-east-1:123456789012:abcdefg1234/*/*/*
```

## Lambda Function URL

A Lambda function URL creates an HTTPS endpoint for direct invocation without API Gateway. Function URLs are useful for lightweight APIs, webhooks, and external integrations.

### Key attributes

- **target_function_arn**: ARN of the function to expose (required). Must reference an `aws.lambda.function` resource.
- **auth_type**: Authentication mode (required). Set to `AWS_IAM` for IAM-authenticated access or `NONE` for public endpoints.
- **cors**: CORS configuration map (optional) specifying allowed origins, methods, and headers.
- **invoke_mode**: Invocation mode (`BUFFERED` or `RESPONSE_STREAM`). Defaults to `BUFFERED`. Set to `RESPONSE_STREAM` for streaming responses.
- **qualifier**: Alias or version qualifier (optional). Restricts the URL to a specific version or alias.

### Example

```yaml
api_url:
  type: aws.url
  target_function_arn: ${api_function.Arn}
  auth_type: NONE
  cors:
    AllowOrigins:
      - "*"
    AllowMethods:
      - GET
      - POST
    AllowHeaders:
      - Content-Type
```

## Common pitfalls

- **Deployment code must already exist**: The plugin does not upload code to S3 or push images to ECR. Upload or push before declaring the function.
- **Missing execution role**: Functions require an execution role ARN, not just a role name. Use `${role.Arn}` to pass the ARN.
- **Handler required for zip functions only**: Zip-based functions need a handler; container images do not. Omit `handler` and `runtime` for container images.
- **Runtime required for zip functions**: Zip-based functions need a runtime identifier; container images do not. Omit `runtime` for container images.
- **Memory affects CPU and cost**: Increasing `memory_size` also increases CPU allocation and billing. Tune based on actual requirements.
- **Permissions are not optional for external invocations**: If API Gateway, S3, or another AWS service invokes the function, add a `aws.lambda.permission` resource.
- **Function URLs are public by default**: A function URL with `auth_type: NONE` is publicly accessible. Use `auth_type: AWS_IAM` to restrict to authenticated users.

## Reference pages

- [aws.lambda.function](../reference/lambda/lambda-function.md) — Function configuration and attributes
- [aws.lambda.permission](../reference/lambda/lambda-permission.md) — Permission schema
- [aws.url](../reference/lambda/url.md) — Function URL configuration
