# aws.iot.policy

**CloudFormation type:** `AWS::IoT::Policy`

Resource Type definition for AWS::IoT::Policy

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::IoT::Policy)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `PolicyDocument` | policy_document | `string` | required |  |  |
| `PolicyName` | policy_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  |  |

Supports update: yes

Discovery: supported
