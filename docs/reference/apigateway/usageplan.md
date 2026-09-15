# aws.usageplan

**CloudFormation type:** `AWS::ApiGateway::UsagePlan`

The ``AWS::ApiGateway::UsagePlan`` resource creates a usage plan for deployed APIs. A usage plan sets a target for the throttling and quota limits on individual client API keys. For more information, see [Creating and Using API Usage Plans in Amazon API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html) in the *API Gateway Developer Guide*.

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ApiGateway::UsagePlan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApiStages` | api_stages | `list` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `Quota` |  | `map` | optional, computed, provider-chosen |  | ``QuotaSettings`` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies a target for the maximum number of requests users can make to your REST APIs. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `Throttle` |  | `map` | optional, computed, provider-chosen |  | ``ThrottleSettings`` is a property of the [AWS::ApiGateway::UsagePlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html) resource that specifies the overall request rate (average requests per second) and burst capacity when users call your REST APIs. |
| `UsagePlanName` | usage_plan_name | `string` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
