# Lambda Function Module

Creates a Lambda function with execution role and CloudWatch Logs log group.

## Resources

- IAM role for Lambda execution (with basic execution policy)
- CloudWatch Logs log group
- Lambda function

## Inputs

| Input | Type | Default | Description |
|-------|------|---------|-------------|
| `name` | string | (required) | Lambda function name |
| `s3_bucket` | string | (required) | S3 bucket containing the function code |
| `s3_key` | string | (required) | S3 key (path) to the function code .zip file |
| `handler` | string | (required) | Handler function (e.g., `index.handler`) |
| `runtime` | string | `python3.12` | Lambda runtime (python3.12, nodejs20.x, etc.) |
| `memory` | integer | `256` | Memory allocation in MB (128-10240) |
| `timeout` | integer | `30` | Function timeout in seconds (1-900) |

## Outputs

| Output | Description |
|--------|-------------|
| `function_name` | Lambda function name |
| `function_arn` | Lambda function ARN |

## Usage

```yaml
modules:
  - ./modules/lambda-function

resources:
  api_handler:
    type: module.lambda_function
    name: my-api-handler
    s3_bucket: my-code-bucket
    s3_key: api-handler.zip
    handler: index.handler
    runtime: python3.12
    memory: 512
    timeout: 60
```

Note: Function code must be uploaded to the specified S3 bucket before deployment. The function uses the basic Lambda execution role which grants CloudWatch Logs permissions.
