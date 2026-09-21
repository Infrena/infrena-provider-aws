# aws.iot.authorizer

**CloudFormation type:** `AWS::IoT::Authorizer`

Creates an authorizer.

Region attribute: `region`

**Import ID:** `<region>/AuthorizerName` (AWS::IoT::Authorizer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AuthorizerFunctionArn` | authorizer_function_arn | `string` | required |  |  |
| `AuthorizerName` | authorizer_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `EnableCachingForHttp` | enable_caching_for_http | `boolean` | optional, computed, provider-chosen |  |  |
| `SigningDisabled` | signing_disabled | `boolean` | optional, computed, provider-chosen, replaces on change |  |  |
| `Status` |  | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | tags map |  |  |
| `TokenKeyName` | token_key_name | `string` | optional, computed, provider-chosen |  |  |
| `TokenSigningPublicKeys` | token_signing_public_keys | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
