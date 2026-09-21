# aws.feed

**CloudFormation type:** `AWS::ElementalInference::Feed`

Represents a feed that receives media for inference processing

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::ElementalInference::Feed)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessRoleArn` | access_role_arn | `string` | optional, computed, provider-chosen | aws.role.Arn |  |
| `Arn` |  | `string` | computed |  |  |
| `DataEndpoints` | data_endpoints | `list` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `Name` |  | `string` | required |  |  |
| `Outputs` |  | `list` | required |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
