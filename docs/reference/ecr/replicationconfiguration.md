# aws.replicationconfiguration

**CloudFormation type:** `AWS::ECR::ReplicationConfiguration`

The ``AWS::ECR::ReplicationConfiguration`` resource creates or updates the replication configuration for a private registry. The first time a replication configuration is applied to a private registry, a service-linked IAM role is created in your account for the replication process. For more information, see [Using Service-Linked Roles for Amazon ECR](https://docs.aws.amazon.com/AmazonECR/latest/userguide/using-service-linked-roles.html) in the *Amazon Elastic Container Registry User Guide*.

Region attribute: `region`

**Import ID:** `<region>/RegistryId` (AWS::ECR::ReplicationConfiguration)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `RegistryId` | registry_id | `string` | computed |  |  |
| `ReplicationConfiguration` | replication_configuration | `map` | required |  | The replication configuration for a registry. |

Supports update: yes

Discovery: supported
