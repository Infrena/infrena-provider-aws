# aws.thingtype

**CloudFormation type:** `AWS::IoT::ThingType`

Resource Type definition for AWS::IoT::ThingType

Region attribute: `region`

**Import ID:** `<region>/ThingTypeName` (AWS::IoT::ThingType)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `DeprecateThingType` | deprecate_thing_type | `boolean` | optional, computed, provider-chosen |  |  |
| `Id` |  | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `ThingTypeName` | thing_type_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ThingTypeProperties` | thing_type_properties | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
