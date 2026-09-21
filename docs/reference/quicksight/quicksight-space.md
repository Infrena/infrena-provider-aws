# aws.quicksight.space

**CloudFormation type:** `AWS::QuickSight::Space`

Resource Type definition for AWS::QuickSight::Space

Region attribute: `region`

**Import ID:** `<region>/AwsAccountId|SpaceId` (AWS::QuickSight::Space)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the space. |
| `AwsAccountId` | aws_account_id | `string` | required, replaces on change |  | The ID of the Amazon Web Services account where the space is being created. |
| `CreatedAt` | created_at | `string` | computed |  | The date and time the space was created. |
| `CreatedBy` | created_by | `string` | computed |  | The user name of the principal who created the space. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | A description of the space. |
| `Name` |  | `string` | required |  | The display name of the space. |
| `Permissions` |  | `list` | optional, computed, provider-chosen |  | A list of permissions granted on the space. |
| `Resources` |  | `list` | optional, computed, provider-chosen |  | A list of QuickSight resources attached to the space. |
| `SpaceId` | space_id | `string` | required, replaces on change | aws.quicksight.space.SpaceId | The unique identifier for the space. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs to associate with the space resource. |
| `UpdatedAt` | updated_at | `string` | computed |  | The date and time the space was last updated. |

Supports update: yes

Discovery: supported (parent resource required)
