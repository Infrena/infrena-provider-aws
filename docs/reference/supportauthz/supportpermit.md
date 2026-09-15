# aws.supportpermit

**CloudFormation type:** `AWS::SupportAuthZ::SupportPermit`

Resource Type definition for AWS::SupportAuthZ::SupportPermit. Represents a support permit that grants AWS support time-bounded access to one or more resources for a set of actions.

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::SupportAuthZ::SupportPermit)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the support permit. |
| `CreatedAt` | created_at | `string` | computed |  | The time at which the support permit was created. |
| `Description` |  | `string` | optional, computed, provider-chosen, replaces on change |  | An optional description of the support permit. |
| `Name` |  | `string` | required, replaces on change |  | The name of the support permit. |
| `Permit` |  | `map` | required, replaces on change |  | The grant definition: which actions on which resources, optionally constrained by time conditions. |
| `PermitId` | permit_id | `string` | computed |  | The service-generated identifier of the support permit (the resource segment of the ARN). |
| `SigningKeyInfo` | signing_key_info | `map` | required, replaces on change |  | The signing key used by the permit. Exactly one key type must be provided. |
| `Status` |  | `string` | computed |  | The current status of the support permit. |
| `SupportCaseDisplayId` | support_case_display_id | `string` | optional, computed, provider-chosen, replaces on change |  | The support case display identifier associated with the permit. When provided, the permit is linked to the specified AWS Support case. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | A list of key-value pairs to apply to this resource. |

Supports update: yes

Discovery: supported
