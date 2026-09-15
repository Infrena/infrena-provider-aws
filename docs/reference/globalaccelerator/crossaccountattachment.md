# aws.crossaccountattachment

**CloudFormation type:** `AWS::GlobalAccelerator::CrossAccountAttachment`

Resource Type definition for AWS::GlobalAccelerator::CrossAccountAttachment

Region attribute: `region`

**Import ID:** `<region>/AttachmentArn` (AWS::GlobalAccelerator::CrossAccountAttachment)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AttachmentArn` | attachment_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the attachment. |
| `Name` |  | `string` | required |  | The Friendly identifier of the attachment. |
| `Principals` |  | `list` | optional, computed, provider-chosen |  | Principals to share the resources with. |
| `Resources` |  | `list` | optional, computed, provider-chosen |  | Resources shared using the attachment. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
