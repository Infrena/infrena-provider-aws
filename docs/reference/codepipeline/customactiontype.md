# aws.customactiontype

**CloudFormation type:** `AWS::CodePipeline::CustomActionType`

The AWS::CodePipeline::CustomActionType resource creates a custom action for activities that aren't included in the CodePipeline default actions, such as running an internally developed build process or a test suite. You can use these custom actions in the stage of a pipeline.

Region attribute: `region`

**Import ID:** `<region>/Category|Provider|Version` (AWS::CodePipeline::CustomActionType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Category` |  | `string` | required, replaces on change |  | The category of the custom action, such as a build action or a test action. |
| `ConfigurationProperties` | configuration_properties | `list` | optional, computed, provider-chosen, replaces on change |  | The configuration properties for the custom action. |
| `Id` |  | `string` | computed |  |  |
| `InputArtifactDetails` | input_artifact_details | `map` | required, replaces on change |  | Returns information about the details of an artifact. |
| `OutputArtifactDetails` | output_artifact_details | `map` | required, replaces on change |  | Returns information about the details of an artifact. |
| `Provider` | provider_value | `string` | required, replaces on change |  | The provider of the service used in the custom action, such as AWS CodeDeploy. |
| `Settings` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Settings is a property of the AWS::CodePipeline::CustomActionType resource that provides URLs that users can access to view information about the CodePipeline custom action. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | Any tags assigned to the custom action. |
| `Version` |  | `string` | required, replaces on change |  | The version identifier of the custom action. |

Supports update: yes

Discovery: supported
