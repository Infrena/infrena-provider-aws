# aws.vpcorigin

**CloudFormation type:** `AWS::CloudFront::VpcOrigin`

An Amazon CloudFront VPC origin.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::VpcOrigin)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `AccountId` | account_id | `string` | computed |  |  |
| `Arn` |  | `string` | computed |  |  |
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `LastModifiedTime` | last_modified_time | `string` | computed |  |  |
| `Status` |  | `string` | computed |  |  |
| `Tags` |  | `map` | tags map |  | A complex type that contains zero or more ``Tag`` elements. |
| `VpcOriginEndpointConfig` | vpc_origin_endpoint_config | `map` | required |  | An Amazon CloudFront VPC origin endpoint configuration. |

Supports update: yes

Discovery: supported
