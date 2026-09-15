# aws.cloudfront.publickey

**CloudFormation type:** `AWS::CloudFront::PublicKey`

A public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html).

Global type (no region attribute)

**Import ID:** `global/Id` (AWS::CloudFront::PublicKey)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreatedTime` | created_time | `string` | computed |  |  |
| `Id` |  | `string` | computed |  |  |
| `PublicKeyConfig` | public_key_config | `map` | required |  | Configuration information about a public key that you can use with [signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html), or with [field-level encryption](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/field-level-encryption.html). |

Supports update: yes

Discovery: supported
