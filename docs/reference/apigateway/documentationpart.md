# aws.documentationpart

**CloudFormation type:** `AWS::ApiGateway::DocumentationPart`

The ``AWS::ApiGateway::DocumentationPart`` resource creates a documentation part for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/DocumentationPartId|RestApiId` (AWS::ApiGateway::DocumentationPart)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DocumentationPartId` | documentation_part_id | `string` | computed |  |  |
| `Location` |  | `map` | required, replaces on change |  | The ``Location`` property specifies the location of the Amazon API Gateway API entity that the documentation applies to. ``Location`` is a property of the [AWS::ApiGateway::DocumentationPart](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-documentationpart.html) resource. |
| `Properties` |  | `string` | required |  |  |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |

Supports update: yes

Discovery: supported (parent resource required)
