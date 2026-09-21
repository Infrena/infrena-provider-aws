# aws.distributionconfiguration

**CloudFormation type:** `AWS::ImageBuilder::DistributionConfiguration`

Resource Type definition for AWS::ImageBuilder::DistributionConfiguration

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::ImageBuilder::DistributionConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the distribution configuration. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the distribution configuration. |
| `Distributions` |  | `list` | required |  | The distributions of the distribution configuration. |
| `Name` |  | `string` | required, replaces on change |  | The name of the distribution configuration. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | The tags associated with the component. |

Supports update: yes

Discovery: supported
