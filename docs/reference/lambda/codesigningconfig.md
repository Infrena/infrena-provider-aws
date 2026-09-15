# aws.codesigningconfig

**CloudFormation type:** `AWS::Lambda::CodeSigningConfig`

Resource Type definition for AWS::Lambda::CodeSigningConfig.

Region attribute: `region`

**Import ID:** `<region>/CodeSigningConfigArn` (AWS::Lambda::CodeSigningConfig)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AllowedPublishers` | allowed_publishers | `map` | required |  | When the CodeSigningConfig is later on attached to a function, the function code will be expected to be signed by profiles from this list |
| `CodeSigningConfigArn` | code_signing_config_arn | `string` | computed |  | A unique Arn for CodeSigningConfig resource |
| `CodeSigningConfigId` | code_signing_config_id | `string` | computed |  | A unique identifier for CodeSigningConfig resource |
| `CodeSigningPolicies` | code_signing_policies | `map` | optional, computed, provider-chosen |  | Policies to control how to act if a signature is invalid |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the CodeSigningConfig |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of tags to apply to CodeSigningConfig resource |

Supports update: yes

Discovery: supported
