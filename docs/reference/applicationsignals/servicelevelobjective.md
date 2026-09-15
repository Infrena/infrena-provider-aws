# aws.servicelevelobjective

**CloudFormation type:** `AWS::ApplicationSignals::ServiceLevelObjective`

Resource Type definition for AWS::ApplicationSignals::ServiceLevelObjective

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ApplicationSignals::ServiceLevelObjective)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The ARN of this SLO. |
| `BurnRateConfigurations` | burn_rate_configurations | `list` | optional, computed, provider-chosen |  | Each object in this array defines the length of the look-back window used to calculate one burn rate metric for this SLO. The burn rate measures how fast the service is consuming the error budget, relative to the attainment goal of the SLO. |
| `CreatedTime` | created_time | `integer` | computed |  | Epoch time in seconds of the time that this SLO was created |
| `Description` |  | `string` | optional, computed, provider-chosen |  | An optional description for this SLO. Default is 'No description' |
| `EvaluationType` | evaluation_type | `string` | computed |  | Displays whether this is a period-based SLO or a request-based SLO. |
| `ExclusionWindows` | exclusion_windows | `list` | optional, computed, provider-chosen |  | Each object in this array defines a time exclusion window for this SLO. The time exclusion window is used to exclude breaching data points from affecting attainment rate, error budget, and burn rate metrics. |
| `Goal` |  | `map` | optional, computed, provider-chosen |  | A structure that contains the attributes that determine the goal of the SLO. This includes the time period for evaluation and the attainment threshold. |
| `LastUpdatedTime` | last_updated_time | `integer` | computed |  | Epoch time in seconds of the time that this SLO was most recently updated |
| `Name` |  | `string` | required, replaces on change |  | The name of this SLO. |
| `RequestBasedSli` | request_based_sli | `map` | optional, computed, provider-chosen |  | This structure contains information about the performance metric that a request-based SLO monitors. |
| `Sli` |  | `map` | optional, computed, provider-chosen |  | This structure contains information about the performance metric that an SLO monitors. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The list of tag keys and values associated with the resource you specified |

Supports update: yes

Discovery: supported
