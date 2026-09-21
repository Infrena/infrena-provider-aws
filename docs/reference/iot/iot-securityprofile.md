# aws.iot.securityprofile

**CloudFormation type:** `AWS::IoT::SecurityProfile`

A security profile defines a set of expected behaviors for devices in your account.

Region attribute: `region`

**Import ID:** `<region>/SecurityProfileName` (AWS::IoT::SecurityProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AdditionalMetricsToRetainV2` | additional_metrics_to_retain_v2 | `list` | optional, computed, provider-chosen |  | A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here. |
| `AlertTargets` | alert_targets | `map` | optional, computed, provider-chosen |  | Specifies the destinations to which alerts are sent. |
| `Behaviors` |  | `list` | optional, computed, provider-chosen |  | Specifies the behaviors that, when violated by a device (thing), cause an alert. |
| `MetricsExportConfig` | metrics_export_config | `map` | optional, computed, provider-chosen |  | A structure containing the mqtt topic for metrics export. |
| `SecurityProfileArn` | security_profile_arn | `string` | computed |  | The ARN (Amazon resource name) of the created security profile. |
| `SecurityProfileDescription` | security_profile_description | `string` | optional, computed, provider-chosen |  | A description of the security profile. |
| `SecurityProfileName` | security_profile_name | `string` | optional, computed, provider-chosen, replaces on change |  | A unique identifier for the security profile. |
| `Tags` |  | `map` | tags map |  | Metadata that can be used to manage the security profile. |
| `TargetArns` | target_arns | `list` | optional, computed, provider-chosen |  | A set of target ARNs that the security profile is attached to. |

Supports update: yes

Discovery: supported
