# aws.apimapping

**CloudFormation type:** `AWS::ApiGatewayV2::ApiMapping`

The ``AWS::ApiGatewayV2::ApiMapping`` resource contains an API mapping. An API mapping relates a path of your custom domain name to a stage of your API. A custom domain name can have multiple API mappings, but the paths can't overlap. A custom domain can map only to APIs of the same protocol type. For more information, see [CreateApiMapping](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames-domainname-apimappings.html#CreateApiMapping) in the *Amazon API Gateway V2 API Reference*.

Region attribute: `region`

**Import ID:** `<region>/ApiMappingId|DomainName` (AWS::ApiGatewayV2::ApiMapping)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiId` | api_id | `string` | required | aws.apigatewayv2.api.ApiId | The API identifier. |
| `ApiMappingId` | api_mapping_id | `string` | computed |  |  |
| `ApiMappingKey` | api_mapping_key | `string` | optional, computed, provider-chosen |  | The API mapping key. |
| `DomainName` | domain_name | `string` | required, replaces on change |  | The domain name. |
| `Stage` |  | `string` | required |  | The API stage. |

Supports update: yes

Discovery: supported
