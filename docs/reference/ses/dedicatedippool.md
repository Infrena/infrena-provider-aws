# aws.dedicatedippool

**CloudFormation type:** `AWS::SES::DedicatedIpPool`

Resource Type definition for AWS::SES::DedicatedIpPool

Region attribute: `region`

**Import ID:** `<region>/PoolName` (AWS::SES::DedicatedIpPool)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `PoolName` | pool_name | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the dedicated IP pool. |
| `ScalingMode` | scaling_mode | `string` | optional, computed, provider-chosen |  | Specifies whether the dedicated IP pool is managed or not. The default value is STANDARD. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags (keys and values) associated with the dedicated IP pool. |

Supports update: yes

Discovery: supported
