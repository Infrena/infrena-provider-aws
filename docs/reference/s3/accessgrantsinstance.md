# aws.accessgrantsinstance

**CloudFormation type:** `AWS::S3::AccessGrantsInstance`

The AWS::S3::AccessGrantsInstance resource is an Amazon S3 resource type that hosts Access Grants and their associated locations

Region attribute: `region`

**Import ID:** `<region>/AccessGrantsInstanceArn` (AWS::S3::AccessGrantsInstance)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccessGrantsInstanceArn` | access_grants_instance_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the specified Access Grants instance. |
| `AccessGrantsInstanceId` | access_grants_instance_id | `string` | computed |  | A unique identifier for the specified access grants instance. |
| `IdentityCenterArn` | identity_center_arn | `string` | optional, computed, provider-chosen |  | The Amazon Resource Name (ARN) of the specified AWS Identity Center. |
| `Tags` |  | `map` | replaces on change, tags map |  |  |

Supports update: yes

Discovery: supported
