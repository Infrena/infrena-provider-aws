# aws.logs.transformer

**CloudFormation type:** `AWS::Logs::Transformer`

Specifies a transformer on the log group to transform logs into consistent structured and information rich format.

Region attribute: `region`

**Import ID:** `<region>/LogGroupIdentifier` (AWS::Logs::Transformer)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `LogGroupIdentifier` | log_group_identifier | `string` | required, replaces on change |  | Existing log group that you want to associate with this transformer. |
| `TransformerConfig` | transformer_config | `list` | required |  | List of processors in a transformer |

Supports update: yes

Discovery: supported
