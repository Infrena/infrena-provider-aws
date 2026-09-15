# aws.preset

**CloudFormation type:** `AWS::MediaConvert::Preset`

Resource Type definition for AWS::MediaConvert::Preset

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::MediaConvert::Preset)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the output preset, such as arn:aws:mediaconvert:us-west-2:123456789012 |
| `Category` |  | `string` | optional, computed, provider-chosen |  | The new category for the preset, if you are changing it. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The new description for the preset, if you are changing it. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the preset that you are modifying. |
| `SettingsJson` | settings_json | `map` | required |  | Specify, in JSON format, the transcoding job settings for this output preset. This specification must conform to the AWS Elemental MediaConvert job validation. For information about forming this specification, see the Remarks section later in this topic. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  |  |

Supports update: yes

Discovery: supported
