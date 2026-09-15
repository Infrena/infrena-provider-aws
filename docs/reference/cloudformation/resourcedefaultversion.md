# aws.resourcedefaultversion

**CloudFormation type:** `AWS::CloudFormation::ResourceDefaultVersion`

The default version of a resource that has been registered in the CloudFormation Registry.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::CloudFormation::ResourceDefaultVersion)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the type. This is used to uniquely identify a ResourceDefaultVersion |
| `TypeName` | type_name | `string` | optional, computed, provider-chosen |  | The name of the type being registered. |
| `TypeVersionArn` | type_version_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the type version. |
| `VersionId` | version_id | `string` | optional, computed, provider-chosen |  | The ID of an existing version of the resource to set as the default. |

Supports update: yes

Discovery: supported
