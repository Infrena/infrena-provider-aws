# aws.routeserver

**CloudFormation type:** `AWS::EC2::RouteServer`

VPC Route Server

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::EC2::RouteServer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AmazonSideAsn` | amazon_side_asn | `integer` | required, replaces on change |  | The Amazon-side ASN of the Route Server. |
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the Route Server. |
| `Id` |  | `string` | computed |  | The ID of the Route Server. |
| `PersistRoutes` | persist_routes | `string` | optional, computed, provider-chosen |  | Whether to enable persistent routes |
| `PersistRoutesDuration` | persist_routes_duration | `integer` | optional, computed, provider-chosen, write-only |  | The duration of persistent routes in minutes |
| `SnsNotificationsEnabled` | sns_notifications_enabled | `boolean` | optional, computed, provider-chosen |  | Whether to enable SNS notifications |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
