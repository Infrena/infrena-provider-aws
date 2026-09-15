# aws.mpa.identitysource

**CloudFormation type:** `AWS::MPA::IdentitySource`

Resource Type definition for AWS::MPA::IdentitySource.

Region attribute: `region`

**Import ID:** `<region>/IdentitySourceArn` (AWS::MPA::IdentitySource)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  |  |
| `IdentitySourceArn` | identity_source_arn | `string` | computed |  |  |
| `IdentitySourceParameters` | identity_source_parameters | `map` | required, replaces on change |  |  |
| `IdentitySourceType` | identity_source_type | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `StatusCode` | status_code | `string` | computed |  |  |
| `StatusMessage` | status_message | `string` | computed |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |

Supports update: yes

Discovery: supported
