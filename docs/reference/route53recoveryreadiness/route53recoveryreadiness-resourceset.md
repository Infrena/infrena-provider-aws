# aws.route53recoveryreadiness.resourceset

**CloudFormation type:** `AWS::Route53RecoveryReadiness::ResourceSet`

Schema for the AWS Route53 Recovery Readiness ResourceSet Resource and API.

Region attribute: `region`

**Import ID:** `<region>/ResourceSetName` (AWS::Route53RecoveryReadiness::ResourceSet)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ResourceSetArn` | resource_set_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the resource set. |
| `ResourceSetName` | resource_set_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the resource set to create. |
| `ResourceSetType` | resource_set_type | `string` | required, replaces on change |  | The resource type of the resources in the resource set. Enter one of the following values for resource type: |
| `Resources` |  | `list` | required |  | A list of resource objects in the resource set. |
| `Tags` |  | `map` | tags map |  | A tag to associate with the parameters for a resource set. |

Supports update: yes

Discovery: supported
