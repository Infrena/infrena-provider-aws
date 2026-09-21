# aws.healthcheck

**CloudFormation type:** `AWS::Route53::HealthCheck`

Resource schema for AWS::Route53::HealthCheck.

Global type (no region attribute)

**Import ID:** `global/HealthCheckId` (AWS::Route53::HealthCheck)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `HealthCheckConfig` | health_check_config | `map` | required |  | A complex type that contains information about the health check. |
| `HealthCheckId` | health_check_id | `string` | computed |  |  |
| `HealthCheckTags` | health_check_tags | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
