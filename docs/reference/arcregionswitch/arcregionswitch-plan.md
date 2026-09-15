# aws.arcregionswitch.plan

**CloudFormation type:** `AWS::ARCRegionSwitch::Plan`

Represents a plan that specifies Regions, IAM roles, and workflows of logic required to perform the desired change to your multi-Region application

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ARCRegionSwitch::Plan)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AssociatedAlarms` | associated_alarms | `map` | optional, computed, provider-chosen |  |  |
| `Description` |  | `string` | optional, computed, provider-chosen |  |  |
| `ExecutionRole` | execution_role | `string` | required |  |  |
| `HealthChecksForPlan` | health_checks_for_plan | `map` | computed |  |  |
| `Name` |  | `string` | required, replaces on change |  |  |
| `Owner` |  | `string` | computed |  |  |
| `PlanHealthChecks` | plan_health_checks | `list` | computed |  |  |
| `PrimaryRegion` | primary_region | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `RecoveryApproach` | recovery_approach | `string` | required, replaces on change |  |  |
| `RecoveryTimeObjectiveMinutes` | recovery_time_objective_minutes | `float` | optional, computed, provider-chosen |  |  |
| `Regions` |  | `list` | required, replaces on change |  |  |
| `ReportConfiguration` | report_configuration | `map` | optional, computed, provider-chosen |  |  |
| `Route53HealthChecks` | route53_health_checks | `map` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |
| `Triggers` |  | `list` | optional, computed, provider-chosen |  |  |
| `Version` |  | `string` | computed |  |  |
| `Workflows` |  | `list` | required |  |  |

Supports update: yes

Discovery: supported
