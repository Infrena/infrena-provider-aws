# aws.predefinedattribute

**CloudFormation type:** `AWS::Connect::PredefinedAttribute`

Resource Type definition for AWS::Connect::PredefinedAttribute

Region attribute: `region`

**Import ID:** `<region>/InstanceArn|Name` (AWS::Connect::PredefinedAttribute)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttributeConfiguration` | attribute_configuration | `map` | optional, computed, provider-chosen |  | Custom metadata associated to a Predefined attribute that controls how the attribute behaves when used by upstream services. |
| `InstanceArn` | instance_arn | `string` | required, replaces on change | aws.connect.instance.Arn | The identifier of the Amazon Connect instance. |
| `LastModifiedRegion` | last_modified_region | `string` | computed |  | Last modified region. |
| `LastModifiedTime` | last_modified_time | `float` | computed |  | Last modified time. |
| `Name` |  | `string` | required, replaces on change |  | The name of the predefined attribute. |
| `Purposes` |  | `list` | optional, computed, provider-chosen |  | The assigned purposes of the predefined attribute. |
| `Values` |  | `map` | optional, computed, provider-chosen |  | The values of a predefined attribute. |

Supports update: yes

Discovery: supported (parent resource required)
