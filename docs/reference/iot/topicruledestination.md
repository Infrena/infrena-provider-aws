# aws.topicruledestination

**CloudFormation type:** `AWS::IoT::TopicRuleDestination`

Resource Type definition for AWS::IoT::TopicRuleDestination

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IoT::TopicRuleDestination)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name (ARN). |
| `HttpUrlProperties` | http_url_properties | `map` | optional, computed, provider-chosen, replaces on change |  | HTTP URL destination properties. |
| `InfluxDBProperties` | influx_db_properties | `map` | optional, computed, provider-chosen, replaces on change |  | InfluxDB destination properties. |
| `Status` |  | `string` | optional, computed, provider-chosen |  | The status of the TopicRuleDestination. |
| `StatusReason` | status_reason | `string` | computed |  | The reasoning for the current status of the TopicRuleDestination. |
| `VpcProperties` | vpc_properties | `map` | optional, computed, provider-chosen, replaces on change |  | VPC destination properties. |

Supports update: yes

Discovery: supported
