# aws.usageprofile

**CloudFormation type:** `AWS::Glue::UsageProfile`

This creates a Resource of UsageProfile type.

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::Glue::UsageProfile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Configuration` |  | `map` | optional, computed, provider-chosen |  | UsageProfile configuration for supported service ex: (Jobs, Sessions). |
| `CreatedOn` | created_on | `string` | computed |  | Creation time. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the UsageProfile. |
| `Name` |  | `string` | required, replaces on change |  | The name of the UsageProfile. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags to be applied to this UsageProfiles. |

Supports update: yes

Discovery: supported
