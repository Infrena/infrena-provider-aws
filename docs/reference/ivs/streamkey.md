# aws.streamkey

**CloudFormation type:** `AWS::IVS::StreamKey`

Resource Type definition for AWS::IVS::StreamKey

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::IVS::StreamKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | Stream Key ARN is automatically generated on creation and assigned as the unique identifier. |
| `ChannelArn` | channel_arn | `string` | required, replaces on change | aws.ivs.channel.Arn | Channel ARN for the stream. |
| `Tags` |  | `map` | tags map |  | A list of key-value pairs that contain metadata for the asset model. |
| `Value` |  | `string` | computed |  | Stream-key value. |

Supports update: yes

Discovery: supported (parent resource required)
