# aws.servicetemplate

**CloudFormation type:** `AWS::Proton::ServiceTemplate`

Definition of AWS::Proton::ServiceTemplate Resource Type

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::Proton::ServiceTemplate)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | <p>The Amazon Resource Name (ARN) of the service template.</p> |
| `Description` |  | `string` | optional, computed, provider-chosen |  | <p>A description of the service template.</p> |
| `DisplayName` | display_name | `string` | optional, computed, provider-chosen |  | <p>The name of the service template as displayed in the developer interface.</p> |
| `EncryptionKey` | encryption_key | `string` | optional, computed, provider-chosen, replaces on change |  | <p>A customer provided encryption key that's used to encrypt data.</p> |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `PipelineProvisioning` | pipeline_provisioning | `string` | optional, computed, provider-chosen, replaces on change |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | <p>An optional list of metadata items that you can associate with the Proton service template. A tag is a key-value pair.</p> |

Supports update: yes

Discovery: supported
