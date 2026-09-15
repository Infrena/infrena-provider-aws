# aws.apigateway.authorizer

**CloudFormation type:** `AWS::ApiGateway::Authorizer`

The ``AWS::ApiGateway::Authorizer`` resource creates an authorization layer that API Gateway activates for methods that have authorization enabled. API Gateway activates the authorizer when a client calls those methods.

Region attribute: `region`

**Import ID:** `<region>/RestApiId|AuthorizerId` (AWS::ApiGateway::Authorizer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AuthType` | auth_type | `string` | optional, computed, provider-chosen |  |  |
| `AuthorizerCredentials` | authorizer_credentials | `string` | optional, computed, provider-chosen |  |  |
| `AuthorizerId` | authorizer_id | `string` | computed |  |  |
| `AuthorizerResultTtlInSeconds` | authorizer_result_ttl_in_seconds | `integer` | optional, computed, provider-chosen |  |  |
| `AuthorizerUri` | authorizer_uri | `string` | optional, computed, provider-chosen |  |  |
| `IdentitySource` | identity_source | `string` | optional, computed, provider-chosen |  |  |
| `IdentityValidationExpression` | identity_validation_expression | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | required |  |  |
| `ProviderARNs` | provider_ar_ns | `list` | optional, computed, provider-chosen |  |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |
| `Type` | type_value | `string` | required |  |  |

Supports update: yes

Discovery: supported (parent resource required)
