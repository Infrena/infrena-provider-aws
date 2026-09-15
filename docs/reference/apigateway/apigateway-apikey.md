# aws.apigateway.apikey

**CloudFormation type:** `AWS::ApiGateway::ApiKey`

The ``AWS::ApiGateway::ApiKey`` resource creates a unique key that you can distribute to clients who are executing API Gateway ``Method`` resources that require an API key. To specify which API key clients must use, map the API key with the ``RestApi`` and ``Stage`` resources that include the methods that require a key.

Region attribute: `region`

**Import ID:** `<region>/APIKeyId` (AWS::ApiGateway::ApiKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `APIKeyId` | api_key_id | `string` | computed |  |  |
| `CustomerId` | customer_id | `string` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Enabled` |  | `boolean` | optional, computed, provider-chosen |  |  |
| `GenerateDistinctId` | generate_distinct_id | `boolean` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the API key. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the API key name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html). |
| `StageKeys` | stage_keys | `list` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Value` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
