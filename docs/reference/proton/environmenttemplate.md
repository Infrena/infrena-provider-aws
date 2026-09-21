# aws.environmenttemplate

**CloudFormation type:** `AWS::Proton::EnvironmentTemplate`

Definition of AWS::Proton::EnvironmentTemplate Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Proton::EnvironmentTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the environment template.</p> |
| `Description` |  | `string` | optional, computed, provider-chosen |  | <p>A description of the environment template.</p> |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | <p>The environment template name as displayed in the developer interface.</p> |
| `EncryptionKey` | encryption_key | `string` | optional, computed, provider-chosen, replaces on change |  | <p>A customer provided encryption key that Proton uses to encrypt data.</p> |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Provisioning` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | tags map |  | <p>An optional list of metadata items that you can associate with the Proton environment template. A tag is a key-value pair.</p> |

Supports update: yes

Discovery: supported
