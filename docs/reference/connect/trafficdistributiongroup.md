# aws.trafficdistributiongroup

**CloudFormation type:** `AWS::Connect::TrafficDistributionGroup`

Resource Type definition for AWS::Connect::TrafficDistributionGroup

Region attribute: `region`

**Import ID:** `<region>/TrafficDistributionGroupArn` (AWS::Connect::TrafficDistributionGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | A description for the traffic distribution group. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The identifier of the Amazon Connect instance that has been replicated. |
| `IsDefault` | is_default | `boolean` | computed |  | If this is the default traffic distribution group. |
| `Name` |  | `string` | required, replaces on change |  | The name for the traffic distribution group. |
| `Status` |  | `string` | computed |  | The status of the traffic distribution group. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |
| `TrafficDistributionGroupArn` | traffic_distribution_group_arn | `string` | computed |  | The identifier of the traffic distribution group. |

Supports update: yes

Discovery: supported
