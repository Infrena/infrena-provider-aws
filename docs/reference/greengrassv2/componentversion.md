# aws.componentversion

**CloudFormation type:** `AWS::GreengrassV2::ComponentVersion`

Resource for Greengrass component version.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::GreengrassV2::ComponentVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `ComponentName` | component_name | `string` | computed |  |  |
| `ComponentVersion` | component_version | `string` | computed |  |  |
| `InlineRecipe` | inline_recipe | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `LambdaFunction` | lambda_function | `map` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
