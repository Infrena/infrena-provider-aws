# aws.thinggroup

**CloudFormation type:** `AWS::IoT::ThingGroup`

Resource Type definition for AWS::IoT::ThingGroup

Region attribute: `region`

**Import ID:** `<region>/ThingGroupName` (AWS::IoT::ThingGroup)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `ParentGroupName` | parent_group_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `QueryString` | query_string | `string` | optional, computed, provider-chosen |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | An array of key-value pairs to apply to this resource. |
| `ThingGroupName` | thing_group_name | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `ThingGroupProperties` | thing_group_properties | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
