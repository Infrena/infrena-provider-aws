# aws.s3objectlambda.accesspoint

**CloudFormation type:** `AWS::S3ObjectLambda::AccessPoint`

The AWS::S3ObjectLambda::AccessPoint resource is an Amazon S3ObjectLambda resource type that you can use to add computation to S3 actions

Region attribute: `region`

**Import ID:** `<region>/Name` (AWS::S3ObjectLambda::AccessPoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Alias` |  | `map` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CreationDate` | creation_date | `string` | computed |  | The date and time when the Object lambda Access Point was created. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name you want to assign to this Object lambda Access Point. |
| `ObjectLambdaConfiguration` | object_lambda_configuration | `map` | required |  | Configuration to be applied to this Object lambda Access Point. It specifies Supporting Access Point, Transformation Configurations. Customers can also set if they like to enable Cloudwatch metrics for accesses to this Object lambda Access Point. Default setting for Cloudwatch metrics is disable. |
| `PolicyStatus` | policy_status | `map` | computed |  |  |
| `PublicAccessBlockConfiguration` | public_access_block_configuration | `map` | computed |  | The Public Access Block Configuration is used to block policies that would allow public access to this Object lambda Access Point. All public access to Object lambda Access Points are blocked by default, and any policy that would give public access to them will be also blocked. This behavior cannot be changed for Object lambda Access Points. |

Supports update: yes

Discovery: supported
