# aws.datalakeassociation

**CloudFormation type:** `AWS::Connect::DataLakeAssociation`

Resource schema for AWS::Connect::DataLakeAssociation

Region attribute: `region`

**Import ID:** `<region>/InstanceId|DataSetId|TargetAccountId` (AWS::Connect::DataLakeAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DataSetId` | data_set_id | `string` | required, replaces on change |  | The identifier of the analytics data set. |
| `InstanceId` | instance_id | `string` | required, replaces on change | aws.connect.instance.Id | The identifier of the Amazon Connect instance |
| `ResourceShareArn` | resource_share_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the AWS Resource Access Manager share |
| `ResourceShareId` | resource_share_id | `string` | computed |  | The AWS Resource Access Manager share ID |
| `TargetAccountId` | target_account_id | `string` | optional, computed, provider-chosen, replaces on change |  | The identifier of the target account |

Supports update: no

Discovery: supported (parent resource required)
