# aws.securityhub.hub

**CloudFormation type:** `AWS::SecurityHub::Hub`

The AWS::SecurityHub::Hub resource represents the implementation of the AWS Security Hub service in your account. One hub resource is created for each Region in which you enable Security Hub.

Region attribute: `region`

**Import ID:** `<region>/ARN` (AWS::SecurityHub::Hub)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `ARN` |  | `string` | computed |  | An ARN is automatically created for the customer. |
| `AutoEnableControls` | auto_enable_controls | `boolean` | optional, computed, provider-chosen |  | Whether to automatically enable new controls when they are added to standards that are enabled |
| `ControlFindingGenerator` | control_finding_generator | `string` | optional, computed, provider-chosen |  | This field, used when enabling Security Hub, specifies whether the calling account has consolidated control findings turned on. If the value for this field is set to SECURITY_CONTROL, Security Hub generates a single finding for a control check even when the check applies to multiple enabled standards.  If the value for this field is set to STANDARD_CONTROL, Security Hub generates separate findings for a control check when the check applies to multiple enabled standards. |
| `EnableDefaultStandards` | enable_default_standards | `boolean` | optional, computed, provider-chosen, write-only |  | Whether to enable the security standards that Security Hub has designated as automatically enabled. |
| `SubscribedAt` | subscribed_at | `string` | computed |  | The date and time when Security Hub was enabled in the account. |
| `Tags` |  | `map` | optional, computed, provider-chosen |  | A key-value pair to associate with a resource. |

Supports update: yes

Discovery: supported
