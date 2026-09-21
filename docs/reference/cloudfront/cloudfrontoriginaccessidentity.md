# aws.cloudfrontoriginaccessidentity

**CloudFormation type:** `AWS::CloudFront::CloudFrontOriginAccessIdentity`

The request to create a new origin access identity (OAI). An origin access identity is a special CloudFront user that you can associate with Amazon S3 origins, so that you can secure all or just some of your Amazon S3 content. For more information, see [Restricting Access to Amazon S3 Content by Using an Origin Access Identity](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide*.

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::CloudFrontOriginAccessIdentity)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CloudFrontOriginAccessIdentityConfig` | cloud_front_origin_access_identity_config | `map` | required |  | Origin access identity configuration. Send a ``GET`` request to the ``/CloudFront API version/CloudFront/identity ID/config`` resource. |
| `Id` |  | `string` | computed |  |  |
| `S3CanonicalUserId` | s3_canonical_user_id | `string` | computed |  |  |

Supports update: yes

Discovery: supported
