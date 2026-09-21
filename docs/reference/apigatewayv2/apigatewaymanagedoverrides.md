# aws.apigatewaymanagedoverrides

**CloudFormation type:** `AWS::ApiGatewayV2::ApiGatewayManagedOverrides`

Resource Type definition for AWS::ApiGatewayV2::ApiGatewayManagedOverrides

Region attribute: `region`

**Import ID:** `<region>/ApiId` (AWS::ApiGatewayV2::ApiGatewayManagedOverrides)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required, replaces on change | aws.apigatewayv2.api.ApiId | The ID of the API for which to override the configuration of API Gateway-managed resources. |
| `Id` |  | `string` | computed |  | A TPS Code is automatically generated on creation and assigned as the unique identifier. |
| `Integration` |  | `map` | optional, computed, provider-chosen |  | Overrides the integration configuration for an API Gateway-managed integration. |
| `Route` |  | `map` | optional, computed, provider-chosen |  | Overrides the route configuration for an API Gateway-managed route. |
| `Stage` |  | `map` | optional, computed, provider-chosen |  | Overrides the stage configuration for an API Gateway-managed stage. |

Supports update: yes

Discovery: not supported
