# aws.s3outposts.accesspoint

**CloudFormation type:** `AWS::S3Outposts::AccessPoint`

Resource Type Definition for AWS::S3Outposts::AccessPoint

Region attribute: `region`

**Import ID:** `<region>/Arn` (AWS::S3Outposts::AccessPoint)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Arn` |  | `string` | computed |  | The Amazon Resource Name (ARN) of the specified AccessPoint. |
| `Bucket` |  | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the bucket you want to associate this AccessPoint with. |
| `Name` |  | `string` | required, replaces on change |  | A name for the AccessPoint. |
| `Policy` |  | `map` | optional, computed, provider-chosen |  | The access point policy associated with this access point. |
| `VpcConfiguration` | vpc_configuration | `map` | required, replaces on change |  | Virtual Private Cloud (VPC) from which requests can be made to the AccessPoint. |

Supports update: yes

Discovery: supported
