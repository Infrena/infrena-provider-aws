# aws.enclavecertificateiamroleassociation

**CloudFormation type:** `AWS::EC2::EnclaveCertificateIamRoleAssociation`

Associates an AWS Identity and Access Management (IAM) role with an AWS Certificate Manager (ACM) certificate. This association is based on Amazon Resource Names and it enables the certificate to be used by the ACM for Nitro Enclaves application inside an enclave.

Region attribute: `region`

**Import ID:** `<region>/CertificateArn|RoleArn` (AWS::EC2::EnclaveCertificateIamRoleAssociation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CertificateArn` | certificate_arn | `string` | required, replaces on change |  | The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role. |
| `CertificateS3BucketName` | certificate_s3_bucket_name | `string` | computed |  | The name of the Amazon S3 bucket to which the certificate was uploaded. |
| `CertificateS3ObjectKey` | certificate_s3_object_key | `string` | computed |  | The Amazon S3 object key where the certificate, certificate chain, and encrypted private key bundle are stored. |
| `EncryptionKmsKeyId` | encryption_kms_key_id | `string` | computed |  | The ID of the AWS KMS CMK used to encrypt the private key of the certificate. |
| `RoleArn` | role_arn | `string` | required, replaces on change | aws.role.Arn | The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate. You can associate up to 16 IAM roles with an ACM certificate. |

Supports update: no

Discovery: supported
