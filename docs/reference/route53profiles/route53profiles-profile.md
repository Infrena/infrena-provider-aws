# aws.route53profiles.profile

**CloudFormation type:** `AWS::Route53Profiles::Profile`

Resource Type definition for AWS::Route53Profiles::Profile

Region attribute: `region`

**Import ID:** `<region>/Id` (AWS::Route53Profiles::Profile)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the resolver profile. |
| `ClientToken` | client_token | `string` | computed |  | The id of the creator request |
| `Id` |  | `string` | computed |  | The ID of the profile. |
| `Name` |  | `string` | required, replaces on change |  | The name of the profile. |
| `ShareStatus` | share_status | `string` | computed |  | The sharing status of the profile. |
| `Tags` |  | `map` | tags map |  | An array of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
