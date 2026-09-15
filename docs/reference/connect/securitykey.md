# aws.securitykey

**CloudFormation type:** `AWS::Connect::SecurityKey`

Resource Type definition for AWS::Connect::SecurityKey

Region attribute: `region`

**Import ID:** `<region>/InstanceId|AssociationId` (AWS::Connect::SecurityKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AssociationId` | association_id | `string` | computed |  | An associationID is automatically generated when a storage config is associated with an instance |
| `InstanceId` | instance_id | `string` | required, replaces on change | aws.connect.instance.Id | Amazon Connect instance identifier |
| `Key` |  | `string` | required, replaces on change |  | A valid security key in PEM format. |

Supports update: yes

Discovery: supported
