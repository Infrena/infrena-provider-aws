# aws.approvedorigin

**CloudFormation type:** `AWS::Connect::ApprovedOrigin`

Resource Type definition for AWS::Connect::ApprovedOrigin

Region attribute: `region`

**Import ID:** `<region>/InstanceId|Origin` (AWS::Connect::ApprovedOrigin)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `InstanceId` | instance_id | `string` | required, replaces on change | aws.connect.instance.Id | Amazon Connect instance identifier |
| `Origin` |  | `string` | required, replaces on change |  | Domain name to be added to the allowlist of instance |

Supports update: yes

Discovery: supported
