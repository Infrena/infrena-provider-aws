# aws.lambda.function

**CloudFormation type:** `AWS::Lambda::Function`

The ``AWS::Lambda::Function`` resource creates a Lambda function. To create a function, you need a [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) and an [execution role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html). The deployment package is a .zip file archive or container image that contains your function code. The execution role grants the function permission to use AWS services, such as Amazon CloudWatch Logs for log streaming and AWS X-Ray for request tracing.

Region attribute: `region`

**Import ID:** `<region>/FunctionName` (AWS::Lambda::Function)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Architectures` |  | `list` | optional, computed, provider-chosen |  | The instruction set architecture that the function supports. Enter a string array with one of the valid values (arm64 or x86_64). The default value is ``x86_64``. |
| `Arn` |  | `string` | computed |  |  |
| `CapacityProviderConfig` | capacity_provider_config | `map` | optional, computed, provider-chosen |  | Configuration for the capacity provider that manages compute resources for Lambda functions. |
| `Code` |  | `map` | required |  | The [deployment package](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) for a Lambda function. To deploy a function defined as a container image, you specify the location of a container image in the Amazon ECR registry. For a .zip file deployment package, you can specify the location of an object in Amazon S3. For Node.js and Python functions, you can specify the function code inline in the template. |
| `CodeSigningConfigArn` | code_signing_config_arn | `string` | optional, computed, provider-chosen | aws.codesigningconfig.CodeSigningConfigArn | To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function. |
| `DeadLetterConfig` | dead_letter_config | `map` | optional, computed, provider-chosen |  | The [dead-letter queue](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async-retain-records.html#invocation-dlq) for failed asynchronous invocations. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the function. |
| `DurableConfig` | durable_config | `map` | optional, computed, provider-chosen |  | Configuration settings for [durable functions](https://docs.aws.amazon.com/lambda/latest/dg/durable-functions.html), including execution timeout, retention period for execution history, and an optional ARN of the KMSlong (KMS) customer managed key that is used to encrypt your durable execution's payload data, including input, output, and error payloads. |
| `Environment` |  | `map` | optional, computed, provider-chosen |  | A function's environment variable settings. You can use environment variables to adjust your function's behavior without updating code. An environment variable is a pair of strings that are stored in a function's version-specific configuration. |
| `EphemeralStorage` | ephemeral_storage | `map` | optional, computed, provider-chosen |  | The size of the function's ``/tmp`` directory in MB. The default value is 512, but it can be any whole number between 512 and 10,240 MB. |
| `FileSystemConfigs` | file_system_configs | `list` | optional, computed, provider-chosen |  | Connection settings for an Amazon EFS or Amazon S3 Files file system. To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an [AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html) or [AWS::S3Files::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-mounttarget.html) resource, you must also specify a ``DependsOn`` attribute to ensure that the mount target is created or updated before the function. |
| `FunctionName` | name, function_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the Lambda function, up to 64 characters in length. If you don't specify a name, CFN generates one. |
| `FunctionScalingConfig` | function_scaling_config | `map` | optional, computed, provider-chosen |  | Configuration that defines the scaling behavior for a Lambda Managed Instances function, including the minimum and maximum number of execution environments that can be provisioned. |
| `Handler` |  | `string` | optional, computed, provider-chosen |  | The name of the method within your code that Lambda calls to run your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see [Lambda programming model](https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html). |
| `ImageConfig` | image_config | `map` | optional, computed, provider-chosen |  | Configuration values that override the container image Dockerfile settings. For more information, see [Container image settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms). |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen |  | The ARN of the KMSlong (KMS) customer managed key that's used to encrypt the following resources: |
| `Layers` |  | `list` | optional, computed, provider-chosen |  | A list of [function layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version. |
| `LoggingConfig` | logging_config | `map` | optional, computed, provider-chosen |  | The function's Amazon CloudWatch Logs configuration settings. |
| `MemorySize` | memory_size | `integer` | optional, computed, provider-chosen |  | The amount of [memory available to the function](https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB. Note that new AWS accounts have reduced concurrency and memory quotas. AWS raises these quotas automatically based on your usage. You can also request a quota increase. |
| `PackageType` | package_type | `string` | optional, computed, provider-chosen, replaces on change |  | The type of deployment package. Set to ``Image`` for container image and set ``Zip`` for .zip file archive. |
| `PublishToLatestPublished` | publish_to_latest_published | `boolean` | optional, computed, provider-chosen, write-only |  |  |
| `RecursiveLoop` | recursive_loop | `string` | optional, computed, provider-chosen |  | The function recursion configuration. |
| `ReservedConcurrentExecutions` | reserved_concurrent_executions | `integer` | optional, computed, provider-chosen |  | The number of simultaneous executions to reserve for the function. |
| `Role` |  | `string` | required |  | The Amazon Resource Name (ARN) of the function's execution role. |
| `Runtime` |  | `string` | optional, computed, provider-chosen |  | The identifier of the function's [runtime](https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive. Specifying a runtime results in an error if you're deploying a function using a container image. |
| `RuntimeManagementConfig` | runtime_management_config | `map` | optional, computed, provider-chosen |  | Sets the runtime management configuration for a function's version. For more information, see [Runtime updates](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html). |
| `SnapStart` | snap_start | `map` | optional, computed, provider-chosen, write-only |  | The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting. |
| `SnapStartResponse` | snap_start_response | `map` | computed |  | The function's [SnapStart](https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function. |
| `TenancyConfig` | tenancy_config | `map` | optional, computed, provider-chosen, replaces on change |  | Specifies the tenant isolation mode configuration for a Lambda function. This allows you to configure specific tenant isolation strategies for your function invocations. Tenant isolation configuration cannot be modified after function creation. |
| `Timeout` |  | `integer` | optional, computed, provider-chosen |  | The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see [Lambda execution environment](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html). |
| `TracingConfig` | tracing_config | `map` | optional, computed, provider-chosen |  | The function's [](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html) tracing configuration. To sample and record incoming requests, set ``Mode`` to ``Active``. |
| `VpcConfig` | vpc_config | `map` | optional, computed, provider-chosen |  | The VPC security groups and subnets that are attached to a Lambda function. When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC. For more information, see [VPC Settings](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html). |

Supports update: yes

Discovery: supported
