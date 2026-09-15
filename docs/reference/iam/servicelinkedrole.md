# aws.servicelinkedrole

**CloudFormation type:** `AWS::IAM::ServiceLinkedRole`

Resource Type definition for AWS::IAM::ServiceLinkedRole

Global type (no region attribute)

**Import ID:** `global/RoleName` (AWS::IAM::ServiceLinkedRole)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AWSServiceName` | aws_service_name | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | The service principal for the AWS service to which this role is attached. |
| `CustomSuffix` | custom_suffix | `string` | optional, computed, provider-chosen, replaces on change, write-only |  | A string that you provide, which is combined with the service-provided prefix to form the complete role name. |
| `Description` |  | `string` | optional, computed, provider-chosen |  | The description of the role. |
| `RoleName` | role_name | `string` | computed |  | The name of the role. |

Supports update: yes

Discovery: not supported
