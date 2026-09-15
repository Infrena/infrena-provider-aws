# aws.portalproduct

**CloudFormation type:** `AWS::ApiGatewayV2::PortalProduct`

Represents a portal product in Amazon API Gateway V2, which is a logical grouping of APIs that can be published in a developer portal.

Region attribute: `region`

**Import ID:** `<region>/PortalProductArn` (AWS::ApiGatewayV2::PortalProduct)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the portal product. |
| `DisplayName` | display_name | `string` | required |  | The name of the portal product as it appears in a published portal. |
| `LastModified` | last_modified | `string` | computed |  | The timestamp when the portal product was last modified. |
| `PortalProductArn` | portal_product_arn | `string` | computed |  | The ARN of the portal product. |
| `PortalProductId` | portal_product_id | `string` | computed |  | The portal product identifier. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The collection of tags associated with the portal product. |

Supports update: yes

Discovery: supported
