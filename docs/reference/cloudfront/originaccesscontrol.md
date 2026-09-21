# aws.originaccesscontrol

**CloudFormation type:** `AWS::CloudFront::OriginAccessControl`

Creates a new origin access control in CloudFront. After you create an origin access control, you can add it to an origin in a CloudFront distribution so that CloudFront sends authenticated (signed) requests to the origin.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::OriginAccessControl)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `Id` |  | `string` | computed |  |  |
| `OriginAccessControlConfig` | origin_access_control_config | `map` | required |  | Creates a new origin access control in CloudFront. After you create an origin access control, you can add it to an origin in a CloudFront distribution so that CloudFront sends authenticated (signed) requests to the origin. |

Supports update: yes

Discovery: supported
