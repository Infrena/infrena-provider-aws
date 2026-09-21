# aws.launchtemplate

**CloudFormation type:** `AWS::EC2::LaunchTemplate`

Specifies the properties for creating a launch template.

Region attribute: `region`

**Import ID:** `<region>/LaunchTemplateId` (AWS::EC2::LaunchTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DefaultVersionNumber` | default_version_number | `string` | computed |  |  |
| `LatestVersionNumber` | latest_version_number | `string` | computed |  |  |
| `LaunchTemplateData` | launch_template_data | `map` | required, write-only |  | The information to include in the launch template. |
| `LaunchTemplateId` | launch_template_id | `string` | computed |  |  |
| `LaunchTemplateName` | launch_template_name | `string` | optional, computed, provider-chosen, replaces on change |  | A name for the launch template. |
| `TagSpecifications` | tag_specifications | `list` | optional, computed, provider-chosen, write-only |  | The tags to apply to the launch template on creation. To tag the launch template, the resource type must be ``launch-template``. |
| `VersionDescription` | version_description | `string` | optional, computed, provider-chosen, write-only |  | A description for the first version of the launch template. |

Supports update: yes

Discovery: supported
