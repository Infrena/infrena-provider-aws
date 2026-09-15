# aws.dsql.cluster

**CloudFormation type:** `AWS::DSQL::Cluster`

Resource Type definition for AWS::DSQL::Cluster

Region attribute: `region`

**Import ID:** `<region>/Identifier` (AWS::DSQL::Cluster)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `CreationTime` | creation_time | `string` | computed |  | The time of when the cluster was created in ISO-8601 format. |
| `DeletionProtectionEnabled` | deletion_protection_enabled | `boolean` | optional, computed, provider-chosen |  | Whether deletion protection is enabled in this cluster. |
| `EncryptionDetails` | encryption_details | `map` | computed |  | The encryption configuration details for the cluster. |
| `Endpoint` |  | `string` | computed |  | The DSQL cluster endpoint. |
| `Identifier` |  | `string` | computed |  | The ID of the created cluster. |
| `KmsEncryptionKey` | kms_encryption_key | `string` | optional, computed, provider-chosen, write-only |  | The KMS key that encrypts data on the cluster. |
| `MultiRegionProperties` | multi_region_properties | `map` | optional, computed, provider-chosen |  | The Multi-region properties associated to this cluster. |
| `PolicyDocument` | policy_document | `string` | optional, computed, provider-chosen |  | The IAM policy applied to the cluster resource. |
| `PolicyVersion` | policy_version | `string` | computed |  | The version number of the cluster's resource based policy |
| `ResourceArn` | resource_arn | `string` | computed |  | The Amazon Resource Name (ARN) for the cluster. |
| `Status` |  | `string` | computed |  | The status of the cluster. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  |  |
| `VpcEndpoint` | vpc_endpoint | `string` | computed |  | The DSQL cluster VPC endpoint. |
| `VpcEndpointServiceName` | vpc_endpoint_service_name | `string` | computed |  | The VPC endpoint service name. |

Supports update: yes

Discovery: supported
