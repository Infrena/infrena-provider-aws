# aws.location

**CloudFormation type:** `AWS::GameLift::Location`

The AWS::GameLift::Location resource creates an Amazon GameLift (GameLift) custom location.

Region attribute: `region`

**Import ID:** `<region>/LocationName` (AWS::GameLift::Location)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LocationArn` | location_arn | `string` | computed |  |  |
| `LocationName` | location_name | `string` | required, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
