# aws.appimageconfig

**CloudFormation type:** `AWS::SageMaker::AppImageConfig`

Resource Type definition for AWS::SageMaker::AppImageConfig

Region attribute: `region`

**Import ID:** `<region>/AppImageConfigName` (AWS::SageMaker::AppImageConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AppImageConfigArn` | app_image_config_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the AppImageConfig. |
| `AppImageConfigName` | app_image_config_name | `string` | required, replaces on change |  | The Name of the AppImageConfig. |
| `CodeEditorAppImageConfig` | code_editor_app_image_config | `map` | optional, computed, provider-chosen |  | The configuration for the kernels in a SageMaker image running as a CodeEditor app. |
| `JupyterLabAppImageConfig` | jupyter_lab_app_image_config | `map` | optional, computed, provider-chosen |  | The configuration for the kernels in a SageMaker image running as a JupyterLab app. |
| `KernelGatewayImageConfig` | kernel_gateway_image_config | `map` | optional, computed, provider-chosen |  | The configuration for the file system and kernels in a SageMaker image running as a KernelGateway app. |
| `Tags` |  | `map` | tags map |  | A list of tags to apply to the AppImageConfig. |

Supports update: yes

Discovery: supported
