# aws.protection

**CloudFormation type:** `AWS::Shield::Protection`

Enables AWS Shield Advanced for a specific AWS resource. The resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone, AWS Global Accelerator standard accelerator, Elastic IP Address, Application Load Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and Network Load Balancers by association with protected Amazon EC2 Elastic IP addresses.

Region attribute: `region`

**Import ID:** `<region>/ProtectionArn` (AWS::Shield::Protection)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ApplicationLayerAutomaticResponseConfiguration` | application_layer_automatic_response_configuration | `map` | optional, computed, provider-chosen |  | The automatic application layer DDoS mitigation settings for a Protection. This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks. |
| `HealthCheckArns` | health_check_arns | `list` | optional, computed, provider-chosen |  | The Amazon Resource Names (ARNs) of the health check to associate with the protection. |
| `Name` |  | `string` | required, replaces on change |  | Friendly name for the Protection. |
| `ProtectionArn` | protection_arn | `string` | computed |  | The ARN (Amazon Resource Name) of the protection. |
| `ProtectionId` | protection_id | `string` | computed |  | The unique identifier (ID) of the protection. |
| `ResourceArn` | resource_arn | `string` | required, replaces on change |  | The ARN (Amazon Resource Name) of the resource to be protected. |
| `Tags` |  | `map` | tags map |  | One or more tag key-value pairs for the Protection object. |

Supports update: yes

Discovery: supported
