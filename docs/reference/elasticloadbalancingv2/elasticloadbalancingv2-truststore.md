# aws.elasticloadbalancingv2.truststore

**CloudFormation type:** `AWS::ElasticLoadBalancingV2::TrustStore`

Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStore

Region attribute: `region`

**Import ID:** `<region>/TrustStoreArn` (AWS::ElasticLoadBalancingV2::TrustStore)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CaCertificatesBundleS3Bucket` | ca_certificates_bundle_s3_bucket | `string` | optional, computed, provider-chosen, write-only |  | The name of the S3 bucket to fetch the CA certificate bundle from. |
| `CaCertificatesBundleS3Key` | ca_certificates_bundle_s3_key | `string` | optional, computed, provider-chosen, write-only |  | The name of the S3 object to fetch the CA certificate bundle from. |
| `CaCertificatesBundleS3ObjectVersion` | ca_certificates_bundle_s3_object_version | `string` | optional, computed, provider-chosen, write-only |  | The version of the S3 bucket that contains the CA certificate bundle. |
| `Name` |  | `string` | optional, computed, provider-chosen, replaces on change |  | The name of the trust store. |
| `NumberOfCaCertificates` | number_of_ca_certificates | `integer` | computed |  | The number of certificates associated with the trust store. |
| `Status` |  | `string` | computed |  | The status of the trust store, could be either of ACTIVE or CREATING. |
| `Tags` |  | `map` | tags map |  | The tags to assign to the trust store. |
| `TrustStoreArn` | trust_store_arn | `string` | computed |  | The Amazon Resource Name (ARN) of the trust store. |

Supports update: yes

Discovery: supported
