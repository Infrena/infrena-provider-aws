# aws.keyvaluestore

**CloudFormation type:** `AWS::CloudFront::KeyValueStore`

The key value store. Use this to separate data from function code, allowing you to update data without having to publish a new version of a function. The key value store holds keys and their corresponding values.

Global type (no region attribute)

**Import ID:** `global/Name` (AWS::CloudFront::KeyValueStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  |  |
| `Comment` |  | `string` | optional, computed, provider-chosen |  | A comment to describe the Key Value Store. Omitting ``Comment`` from the template during updates will clear the existing comment (set to empty string). To preserve an existing comment, you must explicitly include it in the template. |
| `Id` |  | `string` | computed |  |  |
| `ImportSource` | import_source | `map` | optional, computed, provider-chosen, write-only |  | The import source for the key value store. |
| `Name` |  | `string` | required, replaces on change |  | The name of the key value store. |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | A complex type that contains zero or more ``Tag`` elements. |

Supports update: yes

Discovery: supported
