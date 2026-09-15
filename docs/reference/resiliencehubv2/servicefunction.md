# aws.servicefunction

**CloudFormation type:** `AWS::ResilienceHubV2::ServiceFunction`

Creates a service function within a Resilience Hub service.

Region attribute: `region`

**Import ID:** `<region>/ServiceArn|ServiceFunctionId` (AWS::ResilienceHubV2::ServiceFunction)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedAt` | created_at | `string` | computed |  | The timestamp when the service function was created. |
| `Criticality` |  | `string` | required |  | The criticality of the service function. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the service function. |
| `Name` |  | `string` | required |  | The name of the service function. |
| `ResourceCount` | resource_count | `integer` | computed |  | The number of resources associated with this function. |
| `ServiceArn` | service_arn | `string` | required, replaces on change | aws.resiliencehubv2.service.ServiceArn | The ARN of the parent service. |
| `ServiceFunctionId` | service_function_id | `string` | computed |  | The server-generated service function ID. |
| `Source` |  | `string` | computed |  | The source of the service function. |
| `UpdatedAt` | updated_at | `string` | computed |  | The timestamp when the service function was last updated. |

Supports update: yes

Discovery: supported (parent resource required)
