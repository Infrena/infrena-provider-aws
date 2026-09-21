# aws.encoderconfiguration

**CloudFormation type:** `AWS::IVS::EncoderConfiguration`

Resource Type definition for AWS::IVS::EncoderConfiguration.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::EncoderConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Encoder configuration identifier. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | Encoder configuration name. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |
| `Video` |  | `map` | optional, computed, provider-chosen, replaces on change |  | Video configuration. Default: video resolution 1280x720, bitrate 2500 kbps, 30 fps |

Supports update: yes

Discovery: supported
