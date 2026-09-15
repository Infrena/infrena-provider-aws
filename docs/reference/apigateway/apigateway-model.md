# aws.apigateway.model

**CloudFormation type:** `AWS::ApiGateway::Model`

The ``AWS::ApiGateway::Model`` resource defines the structure of a request or response payload for an API method.

Region attribute: `region`

**Import ID:** `<region>/RestApiId|Name` (AWS::ApiGateway::Model)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ContentType` | content_type | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html). |
| `RestApiId` | rest_api_id | `string` | required, replaces on change | aws.restapi.RestApiId |  |
| `Schema` |  | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported (parent resource required)
