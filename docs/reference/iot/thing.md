# aws.thing

**CloudFormation type:** `AWS::IoT::Thing`

Resource Type definition for AWS::IoT::Thing

Region attribute: `region`

**Import ID:** `<region>/ThingName` (AWS::IoT::Thing)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `AttributePayload` | attribute_payload | `map` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `ThingName` | thing_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |

Supports update: yes

Discovery: supported
