# aws.neptunegraph.graph

**CloudFormation type:** `AWS::NeptuneGraph::Graph`

The AWS::NeptuneGraph::Graph resource creates an Amazon NeptuneGraph Graph.

Region attribute: `region`

**Import ID:** `<region>/GraphId` (AWS::NeptuneGraph::Graph)

## Attributes

| Attribute | Also written as | Kind | Flags | References | Description |
|-----------|-----------------|------|-------|------------|-------------|
| `DeletionProtection` | deletion_protection | `boolean` | optional, computed, provider-chosen |  | Value that indicates whether the Graph has deletion protection enabled. The graph can't be deleted when deletion protection is enabled. |
| `Endpoint` |  | `string` | computed |  | The connection endpoint for the graph. For example: `g-12a3bcdef4.us-east-1.neptune-graph.amazonaws.com` |
| `GraphArn` | graph_arn | `string` | computed |  | Graph resource ARN |
| `GraphId` | graph_id | `string` | computed |  | The auto-generated id assigned by the service. |
| `GraphName` | graph_name | `string` | optional, computed, provider-chosen, replaces on change |  | Contains a user-supplied name for the Graph. |
| `ImportTask` | import_task | `map` | optional, computed, provider-chosen, replaces on change, write-only |  | The import task details to import data into the graph at creation time. |
| `KmsKeyIdentifier` | kms_key_identifier | `string` | optional, computed, provider-chosen, replaces on change |  | The ARN of the KMS key used to encrypt data in the Neptune Analytics graph. If not specified, the graph is encrypted with an AWS managed key. |
| `ProvisionedMemory` | provisioned_memory | `integer` | required |  | Memory for the Graph. |
| `PublicConnectivity` | public_connectivity | `boolean` | optional, computed, provider-chosen |  | Specifies whether the Graph can be reached over the internet. Access to all graphs requires IAM authentication. |
| `ReplicaCount` | replica_count | `integer` | optional, computed, provider-chosen, replaces on change |  | Specifies the number of replicas you want when finished. All replicas will be provisioned in different availability zones. |
| `Tags` |  | `map` | optional, computed, provider-chosen, tags map |  | The tags associated with this graph. |
| `VectorSearchConfiguration` | vector_search_configuration | `map` | optional, computed, provider-chosen, replaces on change |  | The vector search configuration. |

Supports update: yes

Discovery: supported
