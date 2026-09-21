# aws.viewversion

**CloudFormation type:** `AWS::Connect::ViewVersion`

Resource Type definition for AWS::Connect::ViewVersion

Region attribute: `region`

**Import ID:** `<region>/ViewVersionArn` (AWS::Connect::ViewVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Version` |  | `integer` | computed |  | The version of the view. |
| `VersionDescription` | version_description | `string` | optional, computed, provider-chosen |  | The description for the view version. |
| `ViewArn` | view_arn | `string` | required, replaces on change | aws.connect.view.ViewArn | The Amazon Resource Name (ARN) of the view for which a version is being created. |
| `ViewContentSha256` | view_content_sha256 | `string` | optional, computed, provider-chosen, replaces on change |  | The view content hash to be checked. |
| `ViewVersionArn` | view_version_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the created view version. |

Supports update: yes

Discovery: supported (parent resource required)
