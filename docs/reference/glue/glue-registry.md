# aws.glue.registry

**CloudFormation type:** `AWS::Glue::Registry`

This resource creates a Registry for authoring schemas as part of Glue Schema Registry.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Glue::Registry)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Amazon Resource Name for the created Registry. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the registry. If description is not provided, there will not be any default value for this. |
| `Name` |  | `string` | required, replaces on change |  | Name of the registry to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.  No whitespace. |
| `Tags` |  | `map` | tags map |  | List of tags to tag the Registry |

Supports update: yes

Discovery: supported
