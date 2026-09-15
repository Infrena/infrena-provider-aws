# aws.connect.view

**CloudFormation type:** `AWS::Connect::View`

Resource Type definition for AWS::Connect::View

Region attribute: `region`

**Import ID:** `<region>/ViewArn` (AWS::Connect::View)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Actions` |  | `list` | required |  | The actions of the view in an array. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the view. |
| `InstanceArn` | instance_arn | `string` | required | aws.connect.instance.Arn | The Amazon Resource Name (ARN) of the instance. |
| `Name` |  | `string` | required |  | The name of the view. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | One or more tags. |
| `Template` |  | `map` | required |  | The template of the view as JSON. |
| `ViewArn` | view_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the view. |
| `ViewContentSha256` | view_content_sha256 | `string` | computed |  | The view content hash. |
| `ViewId` | view_id | `string` | computed |  | The view id of the view. |

Supports update: yes

Discovery: supported (parent resource required)
