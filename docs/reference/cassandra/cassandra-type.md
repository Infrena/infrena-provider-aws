# aws.cassandra.type

**CloudFormation type:** `AWS::Cassandra::Type`

Resource schema for AWS::Cassandra::Type

Region attribute: `region`

**Import ID:** `<region>/KeyspaceName|TypeName` (AWS::Cassandra::Type)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DirectParentTypes` | direct_parent_types | `list` | computed |  | List of parent User-Defined Types that directly reference the User-Defined Type in their fields. |
| `DirectReferringTables` | direct_referring_tables | `list` | computed |  | List of Tables that directly reference the User-Defined Type in their columns. |
| `Fields` |  | `list` | required, replaces on change |  | Field definitions of the User-Defined Type |
| `KeyspaceArn` | keyspace_arn | `string` | computed |  | ARN of the Keyspace which contains the User-Defined Type. |
| `KeyspaceName` | keyspace_name | `string` | required, replaces on change |  | Name of the Keyspace which contains the User-Defined Type. |
| `LastModifiedTimestamp` | last_modified_timestamp | `float` | computed |  | Timestamp of the last time the User-Defined Type's meta data was modified. |
| `MaxNestingDepth` | max_nesting_depth | `integer` | computed |  | Maximum nesting depth of the User-Defined Type across the field types. |
| `TypeName` | type_name | `string` | required, replaces on change |  | Name of the User-Defined Type. |

Supports update: no

Discovery: supported
