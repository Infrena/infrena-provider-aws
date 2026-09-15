# aws.truststorerevocation

**CloudFormation type:** `AWS::ElasticLoadBalancingV2::TrustStoreRevocation`

Resource Type definition for AWS::ElasticLoadBalancingV2::TrustStoreRevocation

Region attribute: `region`

**Import ID:** `<region>/RevocationId|TrustStoreArn` (AWS::ElasticLoadBalancingV2::TrustStoreRevocation)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RevocationContents` | revocation_contents | `list` | optional, computed, provider-chosen, replaces on change, write-only |  | The attributes required to create a trust store revocation. |
| `RevocationId` | revocation_id | `integer` | computed |  | The ID associated with the revocation. |
| `TrustStoreArn` | trust_store_arn | `string` | optional, computed, provider-chosen, replaces on change | aws.elasticloadbalancingv2.truststore.TrustStoreArn | The Amazon Resource Name (ARN) of the trust store. |
| `TrustStoreRevocations` | trust_store_revocations | `list` | computed |  | The data associated with a trust store revocation |

Supports update: no

Discovery: supported (parent resource required)
