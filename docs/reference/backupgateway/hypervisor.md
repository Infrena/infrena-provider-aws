# aws.hypervisor

**CloudFormation type:** `AWS::BackupGateway::Hypervisor`

Definition of AWS::BackupGateway::Hypervisor Resource Type

Region attribute: `region`

**Import ID:** `<region>/HypervisorArn` (AWS::BackupGateway::Hypervisor)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Host` |  | `string` | optional, computed, provider-chosen |  |  |
| `HypervisorArn` | hypervisor_arn | `string` | computed |  |  |
| `KmsKeyArn` | kms_key_arn | `string` | optional, computed, provider-chosen, replaces on change, write-only |  |  |
| `LogGroupArn` | log_group_arn | `string` | optional, computed, provider-chosen, write-only | aws.loggroup.Arn |  |
| `Name` |  | `string` | optional, computed, provider-chosen, write-only |  |  |
| `Password` |  | `string` | optional, computed, provider-chosen, sensitive, write-only |  |  |
| `Tags` |  | `map` | optional, computed, provider-chosen, replaces on change, tags map |  |  |
| `Username` |  | `string` | optional, computed, provider-chosen, write-only |  |  |

Supports update: yes

Discovery: supported
